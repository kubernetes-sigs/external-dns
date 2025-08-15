/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package coredns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	etcdcv3 "go.etcd.io/etcd/client/v3"

	"sigs.k8s.io/external-dns/pkg/tlsutils"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	priority    = 10 // default priority when nothing is set
	etcdTimeout = 5 * time.Second

	randomPrefixLabel = "prefix"
)

// coreDNSClient is an interface to work with CoreDNS service records in etcd
type coreDNSClient interface {
	GetServices(prefix string) ([]*Service, error)
	SaveService(value *Service) error
	DeleteService(key string) error
}

type coreDNSProvider struct {
	provider.BaseProvider
	dryRun        bool
	coreDNSPrefix string
	domainFilter  *endpoint.DomainFilter
	client        coreDNSClient
	txtOwnerID    string
}

// Service represents CoreDNS etcd record
type Service struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Text     string `json:"text,omitempty"`
	Mail     bool   `json:"mail,omitempty"` // Be an MX record. Priority becomes Preference.
	TTL      uint32 `json:"ttl,omitempty"`

	// When a SRV record with a "Host: IP-address" is added, we synthesize
	// a srv.Target domain name.  Normally we convert the full Key where
	// the record lives to a DNS name and use this as the srv.Target.  When
	// TargetStrip > 0 we strip the left most TargetStrip labels from the
	// DNS name.
	TargetStrip int `json:"targetstrip,omitempty"`

	// Group is used to group (or *not* to group) different services
	// together. Services with an identical Group are returned in the same
	// answer.
	Group string `json:"group,omitempty"`

	// Etcd key where we found this service and ignored from json un-/marshaling
	Key string `json:"-"`
}

type etcdClient struct {
	client *etcdcv3.Client
	ctx    context.Context
}

var _ coreDNSClient = etcdClient{}

// GetServices GetService return all Service records stored in etcd stored anywhere under the given key (recursively)
func (c etcdClient) GetServices(prefix string) ([]*Service, error) {
	ctx, cancel := context.WithTimeout(c.ctx, etcdTimeout)
	defer cancel()

	path := prefix
	r, err := c.client.Get(ctx, path, etcdcv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var svcs []*Service
	bx := make(map[Service]bool)
	for _, n := range r.Kvs {
		svc := new(Service)
		if err := json.Unmarshal(n.Value, svc); err != nil {
			return nil, fmt.Errorf("%s: %w", n.Key, err)
		}
		b := Service{Host: svc.Host, Port: svc.Port, Priority: svc.Priority, Weight: svc.Weight, Text: svc.Text, Key: string(n.Key)}
		if _, ok := bx[b]; ok {
			// skip the service if already added to service list.
			// the same service might be found in multiple etcd nodes.
			continue
		}
		bx[b] = true

		svc.Key = string(n.Key)
		if svc.Priority == 0 {
			svc.Priority = priority
		}
		svcs = append(svcs, svc)
	}
	return svcs, nil
}

// SaveService persists service data into etcd
func (c etcdClient) SaveService(service *Service) error {
	ctx, cancel := context.WithTimeout(c.ctx, etcdTimeout)
	defer cancel()

	value, err := json.Marshal(&service)
	if err != nil {
		return err
	}
	_, err = c.client.Put(ctx, service.Key, string(value))
	if err != nil {
		return err
	}
	return nil
}

// DeleteService deletes service record from etcd
func (c etcdClient) DeleteService(key string) error {
	ctx, cancel := context.WithTimeout(c.ctx, etcdTimeout)
	defer cancel()

	_, err := c.client.Delete(ctx, key, etcdcv3.WithPrefix())
	return err
}

// builds etcd client config depending on connection scheme and TLS parameters
func getETCDConfig() (*etcdcv3.Config, error) {
	etcdURLsStr := os.Getenv("ETCD_URLS")
	if etcdURLsStr == "" {
		etcdURLsStr = "http://localhost:2379"
	}
	etcdURLs := strings.Split(etcdURLsStr, ",")
	firstURL := strings.ToLower(etcdURLs[0])
	etcdUsername := os.Getenv("ETCD_USERNAME")
	etcdPassword := os.Getenv("ETCD_PASSWORD")
	if strings.HasPrefix(firstURL, "http://") {
		return &etcdcv3.Config{Endpoints: etcdURLs, Username: etcdUsername, Password: etcdPassword}, nil
	} else if strings.HasPrefix(firstURL, "https://") {
		tlsConfig, err := tlsutils.CreateTLSConfig("ETCD")
		if err != nil {
			return nil, err
		}
		log.Debug("using TLS for etcd")
		return &etcdcv3.Config{
			Endpoints: etcdURLs,
			TLS:       tlsConfig,
			Username:  etcdUsername,
			Password:  etcdPassword,
		}, nil
	} else {
		return nil, errors.New("etcd URLs must start with either http:// or https://")
	}
}

// the newETCDClient is an etcd client constructor
func newETCDClient() (coreDNSClient, error) {
	cfg, err := getETCDConfig()
	if err != nil {
		return nil, err
	}
	c, err := etcdcv3.New(*cfg)
	if err != nil {
		return nil, err
	}
	return etcdClient{c, context.Background()}, nil
}

// NewCoreDNSProvider is a CoreDNS provider constructor
func NewCoreDNSProvider(domainFilter *endpoint.DomainFilter, prefix, txtOwnerID string, dryRun bool) (provider.Provider, error) {
	client, err := newETCDClient()
	if err != nil {
		return nil, err
	}

	return coreDNSProvider{
		client:        client,
		dryRun:        dryRun,
		coreDNSPrefix: prefix,
		domainFilter:  domainFilter,
		txtOwnerID:    txtOwnerID,
	}, nil
}

// findEp takes an Endpoint slice and looks for an element in it. If found it will
// return Endpoint, otherwise it will return nil and a bool of false.
func findEp(slice []*endpoint.Endpoint, dnsName string) (*endpoint.Endpoint, bool) {
	for _, item := range slice {
		if item.DNSName == dnsName {
			return item, true
		}
	}
	return nil, false
}

// findLabelInTargets takes an ep.Targets string slice and looks for an element in it. If found it will
// return its string value, otherwise it will return empty string and a bool of false.
func findLabelInTargets(targets []string, label string) (string, bool) {
	for _, target := range targets {
		if target == label {
			return target, true
		}
	}
	return "", false
}

// Records returns all DNS records found in CoreDNS etcd backend. Depending on the record fields
// it may be mapped to one or two records of type A, CNAME, TXT, A+TXT, CNAME+TXT
func (p coreDNSProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	services, err := p.client.GetServices(p.coreDNSPrefix)
	if err != nil {
		return nil, err
	}
	// Group TXT services by dnsName to build multi-target endpoints
	txtServicesByDNS := make(map[string][]*Service)

	for _, service := range services {
		domains := strings.Split(strings.TrimPrefix(service.Key, p.coreDNSPrefix), "/")
		reverse(domains)
		dnsName := strings.Join(domains[service.TargetStrip:], ".")
		if !p.domainFilter.Match(dnsName) {
			continue
		}
		log.Debugf("Getting service (%v) with service host (%s)", service, service.Host)
		prefix := strings.Join(domains[:service.TargetStrip], ".")
		if service.Host != "" {
			ep, found := findEp(result, dnsName)
			if found {
				ep.Targets = append(ep.Targets, service.Host)
				log.Debugf("Extending ep (%s) with new service host (%s)", ep, service.Host)
			} else {
				ep = endpoint.NewEndpointWithTTL(
					dnsName,
					guessRecordType(service.Host),
					endpoint.TTL(service.TTL),
					service.Host,
				)
				log.Debugf("Creating new ep (%s) with new service host (%s)", ep, service.Host)
			}
			ep.Labels["originalText"] = service.Text
			ep.Labels[randomPrefixLabel] = prefix
			ep.Labels[service.Host] = prefix
			result = append(result, ep)
		}
		if service.Text != "" {
			txtServicesByDNS[dnsName] = append(txtServicesByDNS[dnsName], service)
		}
	}

	// Create multi-target TXT endpoints
	for dnsName, txtServices := range txtServicesByDNS {
		if len(txtServices) == 0 {
			continue
		}

		var targets []string
		labels := make(map[string]string)
		var ttl uint32

		for _, service := range txtServices {
			targets = append(targets, service.Text)
			domains := strings.Split(strings.TrimPrefix(service.Key, p.coreDNSPrefix), "/")
			reverse(domains)
			prefix := strings.Join(domains[:service.TargetStrip], ".")
			labels[service.Text] = prefix
			if ttl == 0 {
				ttl = service.TTL
			}
		}

		ep := endpoint.NewEndpointWithTTL(
			dnsName,
			endpoint.RecordTypeTXT,
			endpoint.TTL(ttl),
			targets...,
		)
		ep.Labels = labels
		// Set a default prefix label if none exists
		if ep.Labels[randomPrefixLabel] == "" {
			ep.Labels[randomPrefixLabel] = "default"
		}
		ep.Labels[endpoint.OwnerLabelKey] = p.txtOwnerID
		result = append(result, ep)
	}
	return result, nil
}

func (p coreDNSProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	grouped := p.groupEndpoints(changes)

	for dnsName, group := range grouped {
		if !p.domainFilter.Match(dnsName) {
			log.Debugf("Skipping record %q due to domain filter", dnsName)
			continue
		}
		if err := p.applyGroup(dnsName, group); err != nil {
			return err
		}
	}

	return p.deleteEndpoints(changes.Delete)
}

func (p coreDNSProvider) groupEndpoints(changes *plan.Changes) map[string][]*endpoint.Endpoint {
	grouped := make(map[string][]*endpoint.Endpoint)
	for _, ep := range changes.Create {
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	for i, ep := range changes.UpdateNew {
		ep.Labels = changes.UpdateOld[i].Labels
		log.Debugf("Updating labels (%s) with old labels(%s)", ep.Labels, changes.UpdateOld[i].Labels)
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	return grouped
}

func (p coreDNSProvider) applyGroup(dnsName string, group []*endpoint.Endpoint) error {
	var services []*Service

	for _, ep := range group {
		if ep.RecordType != endpoint.RecordTypeTXT {
			srvs, err := p.createServicesForEndpoint(dnsName, ep)
			if err != nil {
				return err
			}
			services = append(services, srvs...)
		}
	}

	services = p.updateTXTRecords(dnsName, group, services)

	for _, service := range services {
		log.Infof("Add/set key %s to Host=%s, Text=%s, TTL=%d", service.Key, service.Host, service.Text, service.TTL)
		if p.dryRun {
			continue
		}
		if err := p.client.SaveService(service); err != nil {
			return err
		}
	}

	return nil
}

func (p coreDNSProvider) createServicesForEndpoint(dnsName string, ep *endpoint.Endpoint) ([]*Service, error) {
	var services []*Service

	for _, target := range ep.Targets {
		prefix := ep.Labels[target]
		if prefix == "" {
			prefix = fmt.Sprintf("%08x", rand.Int31())
			log.Infof("Generating new prefix: (%s)", prefix)
		}
		service := Service{
			Host:        target,
			Text:        ep.Labels["originalText"],
			Key:         p.etcdKeyFor(prefix + "." + dnsName),
			TargetStrip: strings.Count(prefix, ".") + 1,
			TTL:         uint32(ep.RecordTTL),
		}
		services = append(services, &service)
		ep.Labels[target] = prefix
	}

	// Clean outdated labels
	for label, labelPrefix := range ep.Labels {
		if shouldSkipLabel(label) {
			continue
		}
		if _, ok := findLabelInTargets(ep.Targets, label); !ok {
			key := p.etcdKeyFor(labelPrefix + "." + dnsName)
			log.Infof("Delete key %s", key)
			if p.dryRun {
				continue
			}
			if err := p.client.DeleteService(key); err != nil {
				return nil, err
			}
		}
	}
	return services, nil
}

func shouldSkipLabel(label string) bool {
	skip := []string{"originalText", "prefix", "resource"}
	_, ok := findLabelInTargets(skip, label)
	return ok
}

// updateTXTRecords updates the TXT records in the provided services slice based on the given group of endpoints.
func (p coreDNSProvider) updateTXTRecords(dnsName string, group []*endpoint.Endpoint, services []*Service) []*Service {
	// Collect desired TXT targets in order (preserving user-defined order)
	var orderedTargets []string
	var targetTTL uint32
	var labels map[string]string
	
	for _, ep := range group {
		if ep.RecordType != endpoint.RecordTypeTXT {
			continue
		}
		if ep.Labels == nil {
			ep.Labels = map[string]string{}
		}
		if labels == nil {
			labels = ep.Labels
		}
		
		// Append targets in the order they appear in the endpoint
		for _, t := range ep.Targets {
			orderedTargets = append(orderedTargets, t)
			targetTTL = uint32(ep.RecordTTL)
		}
	}

	if len(orderedTargets) == 0 {
		// no TXT endpoints present
		return services
	}

	// Clear existing TXT services that are no longer needed
	for i, svc := range services {
		if svc.Text != "" {
			found := false
			for _, target := range orderedTargets {
				if target == svc.Text {
					found = true
					break
				}
			}
			if !found {
				services[i].Text = ""
			}
		}
	}

	// Check if we need to reorder existing targets based on current labels
	needsReorder := p.checkIfReorderNeeded(orderedTargets, labels)
	
	if needsReorder {
		// Clean up all existing TXT prefixes and regenerate with correct order
		p.cleanupTXTLabels(labels, orderedTargets, dnsName)
	}

	// Create/update services for each target in order
	for i, target := range orderedTargets {
		prefix := labels[target]
		if prefix == "" || needsReorder {
			// Generate ordered prefix: index + random hex for uniqueness
			prefix = fmt.Sprintf("%d-%06x", i, rand.Int31()&0xFFFFFF)
		}
		
		svc := &Service{
			Key:         p.etcdKeyFor(prefix + "." + dnsName),
			TargetStrip: strings.Count(prefix, ".") + 1,
			TTL:         targetTTL,
			Text:        target,
		}
		services = append(services, svc)
		labels[target] = prefix
	}

	// Cleanup stale TXT keys for targets that are no longer desired
	for label, labelPrefix := range labels {
		if shouldSkipLabel(label) {
			continue
		}
		
		found := false
		for _, target := range orderedTargets {
			if target == label {
				found = true
				break
			}
		}
		
		if !found {
			key := p.etcdKeyFor(labelPrefix + "." + dnsName)
			log.Infof("Delete key %s", key)
			if !p.dryRun && p.client != nil {
				if err := p.client.DeleteService(key); err != nil {
					log.Warnf("Failed to delete stale TXT key %s: %v", key, err)
				}
			}
		}
	}

	return services
}

// checkIfReorderNeeded determines if the current label prefixes match the expected order
func (p coreDNSProvider) checkIfReorderNeeded(orderedTargets []string, labels map[string]string) bool {
	// Count how many existing targets we have vs. total targets
	existingTargets := 0
	for _, target := range orderedTargets {
		if labels[target] != "" {
			existingTargets++
		}
	}
	
	// If we have fewer existing targets than total targets, new targets were added
	// This requires reordering to maintain sequential indices
	if existingTargets != len(orderedTargets) {
		log.Debugf("Reorder needed: have %d existing targets, but %d total targets", existingTargets, len(orderedTargets))
		return true
	}
	
	// Check if all existing targets match their expected indices
	for i, target := range orderedTargets {
		prefix := labels[target]
		if prefix == "" {
			// New target, already handled above
			continue
		}
		
		// Check if prefix starts with the expected index
		expectedPrefix := fmt.Sprintf("%d-", i)
		if !strings.HasPrefix(prefix, expectedPrefix) {
			log.Debugf("Reorder needed: target %q at index %d has prefix %q, expected prefix starting with %q", 
				target, i, prefix, expectedPrefix)
			return true
		}
	}
	return false
}

// cleanupTXTLabels removes existing TXT prefixes that need reordering
func (p coreDNSProvider) cleanupTXTLabels(labels map[string]string, orderedTargets []string, dnsName string) {
	for _, target := range orderedTargets {
		if prefix := labels[target]; prefix != "" {
			key := p.etcdKeyFor(prefix + "." + dnsName)
			log.Infof("Delete key for reordering %s", key)
			if !p.dryRun && p.client != nil {
				if err := p.client.DeleteService(key); err != nil {
					log.Warnf("Failed to delete key for reordering %s: %v", key, err)
				}
			}
			// Clear the prefix so it gets regenerated with correct order
			delete(labels, target)
		}
	}
}

// deleteTXTRecordsForDNSName finds and deletes TXT services that match the specified targets
func (p coreDNSProvider) deleteTXTRecordsForDNSName(dnsName string, targets []string) error {
	// Convert DNS name to etcd path format
	domains := strings.Split(dnsName, ".")
	reverse(domains)
	searchPrefix := p.coreDNSPrefix + strings.Join(domains, "/")

	// Get all services under this DNS name
	services, err := p.client.GetServices(searchPrefix)
	if err != nil {
		return err
	}

	// Create a set of targets to delete for efficient lookup
	targetSet := make(map[string]bool)
	for _, target := range targets {
		targetSet[target] = true
	}

	// Find and delete matching TXT services
	for _, service := range services {
		// Only process TXT services (ones with text content and no host)
		if service.Text != "" && service.Host == "" {
			// Check if this TXT target should be deleted
			if targetSet[service.Text] {
				log.Infof("Delete TXT key %s", service.Key)
				if p.dryRun {
					continue
				}
				if err := p.client.DeleteService(service.Key); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (p coreDNSProvider) deleteEndpoints(endpoints []*endpoint.Endpoint) error {
	for _, ep := range endpoints {
		if ep.RecordType == endpoint.RecordTypeTXT {
			// For TXT records, we need to find and delete all matching services from etcd
			// since TXT records can have multiple targets stored as separate etcd keys
			if err := p.deleteTXTRecordsForDNSName(ep.DNSName, ep.Targets); err != nil {
				return err
			}
		} else {
			// For non-TXT records, use standard deletion logic
			dnsName := ep.DNSName
			if ep.Labels[randomPrefixLabel] != "" {
				dnsName = ep.Labels[randomPrefixLabel] + "." + dnsName
			}
			key := p.etcdKeyFor(dnsName)
			log.Infof("Delete key %s", key)
			if p.dryRun {
				continue
			}
			if err := p.client.DeleteService(key); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p coreDNSProvider) etcdKeyFor(dnsName string) string {
	domains := strings.Split(dnsName, ".")
	reverse(domains)
	return p.coreDNSPrefix + strings.Join(domains, "/")
}

func guessRecordType(target string) string {
	if net.ParseIP(target) != nil {
		return endpoint.RecordTypeA
	}
	return endpoint.RecordTypeCNAME
}

func reverse(slice []string) {
	for i := range len(slice) / 2 {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}
