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
	"slices"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcdcv3 "go.etcd.io/etcd/client/v3"

	"sigs.k8s.io/external-dns/pkg/tlsutils"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	priority    = 10 // default priority when nothing is set
	etcdTimeout = 5 * time.Second

	randomPrefixLabel     = "prefix"
	providerSpecificGroup = "coredns/group"
)

var (
	// avoids allocating a new slice on every call
	skipLabels = []string{"originalText", "prefix", "resource"}
)

// coreDNSClient is an interface to work with CoreDNS service records in etcd
type coreDNSClient interface {
	GetServices(ctx context.Context, prefix string) ([]*Service, error)
	SaveService(ctx context.Context, value *Service) error
	DeleteService(ctx context.Context, key string) error
}

type coreDNSProvider struct {
	provider.BaseProvider
	dryRun        bool
	strictlyOwned bool
	coreDNSPrefix string
	domainFilter  *endpoint.DomainFilter
	client        coreDNSClient
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

	// Owner is used to prevent service to be added by different external-dns (only used by external-dns)
	Owner string `json:"owner,omitempty"`
}

type etcdClient struct {
	client        *etcdcv3.Client
	owner         string
	strictlyOwned bool
}

var _ coreDNSClient = etcdClient{}

// GetServices GetService return all Service records stored in etcd stored anywhere under the given key (recursively)
func (c etcdClient) GetServices(ctx context.Context, prefix string) ([]*Service, error) {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()

	path := prefix
	r, err := c.client.Get(ctx, path, etcdcv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var svcs []*Service
	bx := make(map[Service]bool)
	for _, n := range r.Kvs {
		svc, err := c.unmarshalService(n)
		if err != nil {
			return nil, err
		}
		if c.strictlyOwned && svc.Owner != c.owner {
			continue
		}
		b := Service{
			Host:     svc.Host,
			Port:     svc.Port,
			Priority: svc.Priority,
			Weight:   svc.Weight,
			Text:     svc.Text,
			Key:      string(n.Key),
		}
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
func (c etcdClient) SaveService(ctx context.Context, service *Service) error {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()

	// check only for empty OwnedBy
	if c.strictlyOwned && service.Owner != c.owner {
		r, err := c.client.Get(ctx, service.Key)
		if err != nil {
			return fmt.Errorf("etcd get %q: %w", service.Key, err)
		}
		// Key missing -> treat as owned (safe to create)
		if r != nil && len(r.Kvs) != 0 {
			svc, err := c.unmarshalService(r.Kvs[0])
			if err != nil {
				return fmt.Errorf("failed to unmarshal value for key %q: %w", service.Key, err)
			}
			if svc.Owner != c.owner {
				return fmt.Errorf("key %q is not owned by this provider", service.Key)
			}
		}
		service.Owner = c.owner
	}

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
func (c etcdClient) DeleteService(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()

	if c.strictlyOwned {
		rs, err := c.client.Get(ctx, key, etcdcv3.WithPrefix())
		if err != nil {
			return err
		}
		for _, r := range rs.Kvs {
			svc, err := c.unmarshalService(r)
			if err != nil {
				return err
			}
			if svc.Owner != c.owner {
				continue
			}

			_, err = c.client.Delete(ctx, string(r.Key))
			if err != nil {
				return err
			}
		}
		return err
	} else {
		_, err := c.client.Delete(ctx, key, etcdcv3.WithPrefix())
		return err
	}
}

func (c etcdClient) unmarshalService(n *mvccpb.KeyValue) (*Service, error) {
	svc := new(Service)
	if err := json.Unmarshal(n.Value, svc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %q: %w", n.Key, err)
	}
	return svc, nil
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
	switch {
	case strings.HasPrefix(firstURL, "http://"):
		return &etcdcv3.Config{Endpoints: etcdURLs, Username: etcdUsername, Password: etcdPassword}, nil
	case strings.HasPrefix(firstURL, "https://"):
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
	default:
		return nil, errors.New("etcd URLs must start with either http:// or https://")
	}
}

// the newETCDClient is an etcd client constructor
func newETCDClient(owner string, strictlyOwned bool) (coreDNSClient, error) {
	cfg, err := getETCDConfig()
	if err != nil {
		return nil, err
	}
	c, err := etcdcv3.New(*cfg)
	if err != nil {
		return nil, err
	}
	return etcdClient{c, owner, strictlyOwned}, nil
}

// NewCoreDNSProvider is a CoreDNS provider constructor
func NewCoreDNSProvider(domainFilter *endpoint.DomainFilter, prefix, owner string, strictlyOwned, dryRun bool) (provider.Provider, error) {
	client, err := newETCDClient(owner, strictlyOwned)
	if err != nil {
		return nil, err
	}

	return coreDNSProvider{
		client:        client,
		dryRun:        dryRun,
		strictlyOwned: strictlyOwned,
		coreDNSPrefix: prefix,
		domainFilter:  domainFilter,
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

// Records returns all DNS records found in CoreDNS etcd backend. Depending on the record fields
// it may be mapped to one or two records of type A, CNAME, TXT, A+TXT, CNAME+TXT
func (p coreDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	services, err := p.client.GetServices(ctx, p.coreDNSPrefix)
	if err != nil {
		return nil, err
	}
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
				if service.Group != "" {
					ep.WithProviderSpecific(providerSpecificGroup, service.Group)
				}
				log.Debugf("Creating new ep (%s) with new service host (%s)", ep, service.Host)
			}
			if p.strictlyOwned {
				ep.Labels[endpoint.OwnerLabelKey] = service.Owner
			}
			ep.Labels["originalText"] = service.Text
			ep.Labels[randomPrefixLabel] = prefix
			ep.Labels[service.Host] = prefix
			result = append(result, ep)
		}
		if service.Text != "" {
			ep := endpoint.NewEndpoint(
				dnsName,
				endpoint.RecordTypeTXT,
				service.Text,
			)
			if p.strictlyOwned {
				ep.Labels[endpoint.OwnerLabelKey] = service.Owner
			}
			ep.Labels[randomPrefixLabel] = prefix
			result = append(result, ep)
		}
	}
	return result, nil
}

func (p coreDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	grouped := p.groupEndpoints(changes)

	for dnsName, group := range grouped {
		if !p.domainFilter.Match(dnsName) {
			log.Debugf("Skipping record %q due to domain filter", dnsName)
			continue
		}
		if err := p.applyGroup(ctx, dnsName, group); err != nil {
			return err
		}
	}

	return p.deleteEndpoints(ctx, changes.Delete)
}

func (p coreDNSProvider) groupEndpoints(changes *plan.Changes) map[string][]*endpoint.Endpoint {
	grouped := make(map[string][]*endpoint.Endpoint)
	for _, ep := range changes.Create {
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	for i, ep := range changes.UpdateNew {
		log.Debugf("Updating labels (%s) with old labels (%s)", ep.Labels, changes.UpdateOld[i].Labels)
		ep.Labels = changes.UpdateOld[i].Labels
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	return grouped
}

func (p coreDNSProvider) applyGroup(ctx context.Context, dnsName string, group []*endpoint.Endpoint) error {
	var services []*Service

	for _, ep := range group {
		if ep.RecordType != endpoint.RecordTypeTXT {
			srvs, err := p.createServicesForEndpoint(ctx, dnsName, ep)
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
		if err := p.client.SaveService(ctx, service); err != nil {
			return err
		}
	}

	return nil
}

func (p coreDNSProvider) createServicesForEndpoint(ctx context.Context, dnsName string, ep *endpoint.Endpoint) ([]*Service, error) {
	var services []*Service

	for _, target := range ep.Targets {
		prefix := ep.Labels[target]
		if prefix == "" {
			prefix = fmt.Sprintf("%08x", rand.Int31())
			log.Infof("Generating new prefix: (%s)", prefix)
		}
		group := ""
		if prop, ok := ep.GetProviderSpecificProperty(providerSpecificGroup); ok {
			group = prop
		}
		service := Service{
			Host:        target,
			Text:        ep.Labels["originalText"],
			Key:         p.etcdKeyFor(prefix + "." + dnsName),
			TargetStrip: strings.Count(prefix, ".") + 1,
			TTL:         uint32(ep.RecordTTL),
			Group:       group,
		}
		services = append(services, &service)
		ep.Labels[target] = prefix
	}

	// Clean outdated labels
	for label, labelPrefix := range ep.Labels {
		if slices.Contains(skipLabels, label) {
			continue
		}
		if !slices.Contains(ep.Targets, label) {
			key := p.etcdKeyFor(labelPrefix + "." + dnsName)
			log.Infof("Delete key %s", key)
			if p.dryRun {
				continue
			}
			if err := p.client.DeleteService(ctx, key); err != nil {
				return nil, err
			}
		}
	}
	return services, nil
}

// updateTXTRecords updates the TXT records in the provided services slice based on the given group of endpoints.
func (p coreDNSProvider) updateTXTRecords(dnsName string, group []*endpoint.Endpoint, services []*Service) []*Service {
	index := 0
	for _, ep := range group {
		if ep.RecordType != endpoint.RecordTypeTXT {
			continue
		}
		if index >= len(services) {
			prefix := ep.Labels[randomPrefixLabel]
			if prefix == "" {
				prefix = fmt.Sprintf("%08x", rand.Int31())
			}
			services = append(services, &Service{
				Key:         p.etcdKeyFor(prefix + "." + dnsName),
				TargetStrip: strings.Count(prefix, ".") + 1,
				TTL:         uint32(ep.RecordTTL),
			})
		}
		services[index].Text = ep.Targets[0]
		index++
	}

	for i := index; index > 0 && i < len(services); i++ {
		services[i].Text = ""
	}
	return services
}

func (p coreDNSProvider) deleteEndpoints(ctx context.Context, endpoints []*endpoint.Endpoint) error {
	for _, ep := range endpoints {
		dnsName := ep.DNSName
		if ep.Labels[randomPrefixLabel] != "" {
			dnsName = ep.Labels[randomPrefixLabel] + "." + dnsName
		}
		key := p.etcdKeyFor(dnsName)
		log.Infof("Delete key %s", key)
		if p.dryRun {
			continue
		}
		if err := p.client.DeleteService(ctx, key); err != nil {
			return err
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
