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

package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	etcdcv3 "github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	dryRun        bool
	coreDNSPrefix string
	domainFilter  endpoint.DomainFilter
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

	// Etcd key where we found this service and ignored from json un-/marshalling
	Key string `json:"-"`
}

type etcdClient struct {
	client *etcdcv3.Client
	ctx    context.Context
}

var _ coreDNSClient = etcdClient{}

// GetService return all Service records stored in etcd stored anywhere under the given key (recursively)
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
			return nil, fmt.Errorf("%s: %s", n.Key, err.Error())
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

// loads TLS artifacts and builds tls.Clonfig object
func newTLSConfig(certPath, keyPath, caPath, serverName string, insecure bool) (*tls.Config, error) {
	if certPath != "" && keyPath == "" || certPath == "" && keyPath != "" {
		return nil, errors.New("either both cert and key or none must be provided")
	}
	var certificates []tls.Certificate
	if certPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, fmt.Errorf("could not load TLS cert: %s", err)
		}
		certificates = append(certificates, cert)
	}
	roots, err := loadRoots(caPath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates:       certificates,
		RootCAs:            roots,
		InsecureSkipVerify: insecure,
		ServerName:         serverName,
	}, nil
}

// loads CA cert
func loadRoots(caPath string) (*x509.CertPool, error) {
	if caPath == "" {
		return nil, nil
	}

	roots := x509.NewCertPool()
	pem, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %s", caPath, err)
	}
	ok := roots.AppendCertsFromPEM(pem)
	if !ok {
		return nil, fmt.Errorf("could not read root certs: %s", err)
	}
	return roots, nil
}

// builds etcd client config depending on connection scheme and TLS parameters
func getETCDConfig() (*etcdcv3.Config, error) {
	etcdURLsStr := os.Getenv("ETCD_URLS")
	if etcdURLsStr == "" {
		etcdURLsStr = "http://localhost:2379"
	}
	etcdURLs := strings.Split(etcdURLsStr, ",")
	firstURL := strings.ToLower(etcdURLs[0])
	if strings.HasPrefix(firstURL, "http://") {
		return &etcdcv3.Config{Endpoints: etcdURLs}, nil
	} else if strings.HasPrefix(firstURL, "https://") {
		caFile := os.Getenv("ETCD_CA_FILE")
		certFile := os.Getenv("ETCD_CERT_FILE")
		keyFile := os.Getenv("ETCD_KEY_FILE")
		serverName := os.Getenv("ETCD_TLS_SERVER_NAME")
		isInsecureStr := strings.ToLower(os.Getenv("ETCD_TLS_INSECURE"))
		isInsecure := isInsecureStr == "true" || isInsecureStr == "yes" || isInsecureStr == "1"
		tlsConfig, err := newTLSConfig(certFile, keyFile, caFile, serverName, isInsecure)
		if err != nil {
			return nil, err
		}
		return &etcdcv3.Config{
			Endpoints: etcdURLs,
			TLS:       tlsConfig,
		}, nil
	} else {
		return nil, errors.New("etcd URLs must start with either http:// or https://")
	}
}

//newETCDClient is an etcd client constructor
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
func NewCoreDNSProvider(domainFilter endpoint.DomainFilter, prefix string, dryRun bool) (Provider, error) {
	client, err := newETCDClient()
	if err != nil {
		return nil, err
	}

	return coreDNSProvider{
		client:        client,
		dryRun:        dryRun,
		coreDNSPrefix: prefix,
		domainFilter:  domainFilter,
	}, nil
}

// Records returns all DNS records found in CoreDNS etcd backend. Depending on the record fields
// it may be mapped to one or two records of type A, CNAME, TXT, A+TXT, CNAME+TXT
func (p coreDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	services, err := p.client.GetServices(p.coreDNSPrefix)
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
		prefix := strings.Join(domains[:service.TargetStrip], ".")
		if service.Host != "" {
			ep := endpoint.NewEndpointWithTTL(
				dnsName,
				guessRecordType(service.Host),
				endpoint.TTL(service.TTL),
				service.Host,
			)
			ep.Labels["originalText"] = service.Text
			ep.Labels[randomPrefixLabel] = prefix
			result = append(result, ep)
		}
		if service.Text != "" {
			ep := endpoint.NewEndpoint(
				dnsName,
				endpoint.RecordTypeTXT,
				service.Text,
			)
			ep.Labels[randomPrefixLabel] = prefix
			result = append(result, ep)
		}
	}
	return result, nil
}

// ApplyChanges stores changes back to etcd converting them to CoreDNS format and aggregating A/CNAME and TXT records
func (p coreDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	grouped := map[string][]*endpoint.Endpoint{}
	for _, ep := range changes.Create {
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	for i, ep := range changes.UpdateNew {
		ep.Labels[randomPrefixLabel] = changes.UpdateOld[i].Labels[randomPrefixLabel]
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	for dnsName, group := range grouped {
		if !p.domainFilter.Match(dnsName) {
			log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", dnsName)
			continue
		}
		var services []Service
		for _, ep := range group {
			if ep.RecordType == endpoint.RecordTypeTXT {
				continue
			}

			for _, target := range ep.Targets {
				prefix := ep.Labels[randomPrefixLabel]
				if prefix == "" {
					prefix = fmt.Sprintf("%08x", rand.Int31())
				}

				service := Service{
					Host:        target,
					Text:        ep.Labels["originalText"],
					Key:         p.etcdKeyFor(prefix + "." + dnsName),
					TargetStrip: strings.Count(prefix, ".") + 1,
					TTL:         uint32(ep.RecordTTL),
				}
				services = append(services, service)
			}
		}
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
				services = append(services, Service{
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

		for _, service := range services {
			log.Infof("Add/set key %s to Host=%s, Text=%s, TTL=%d", service.Key, service.Host, service.Text, service.TTL)
			if !p.dryRun {
				err := p.client.SaveService(&service)
				if err != nil {
					return err
				}
			}
		}
	}

	for _, ep := range changes.Delete {
		dnsName := ep.DNSName
		if ep.Labels[randomPrefixLabel] != "" {
			dnsName = ep.Labels[randomPrefixLabel] + "." + dnsName
		}
		key := p.etcdKeyFor(dnsName)
		log.Infof("Delete key %s", key)
		if !p.dryRun {
			err := p.client.DeleteService(key)
			if err != nil {
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
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}
