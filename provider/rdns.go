/*
Copyright 2019 The Kubernetes Authors.

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
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	rdnsMaxHosts      = 10
	rdnsOriginalLabel = "originalText"
	rdnsPrefix        = "/rdnsv3"
	rdnsTimeout       = 5 * time.Second
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RDNSClient is an interface to work with Rancher DNS(RDNS) records in etcdv3 backend.
type RDNSClient interface {
	Get(key string) ([]RDNSRecord, error)
	List(rootDomain string) ([]RDNSRecord, error)
	Set(value RDNSRecord) error
	Delete(key string) error
}

// RDNSConfig contains configuration to create a new Rancher DNS(RDNS) provider.
type RDNSConfig struct {
	DryRun       bool
	DomainFilter endpoint.DomainFilter
	RootDomain   string
}

// RDNSProvider is an implementation of Provider for Rancher DNS(RDNS).
type RDNSProvider struct {
	client       RDNSClient
	dryRun       bool
	domainFilter endpoint.DomainFilter
	rootDomain   string
}

// RDNSRecord represents Rancher DNS(RDNS) etcdv3 record.
type RDNSRecord struct {
	AggregationHosts []string `json:"aggregation_hosts,omitempty"`
	Host             string   `json:"host,omitempty"`
	Text             string   `json:"text,omitempty"`
	TTL              uint32   `json:"ttl,omitempty"`
	Key              string   `json:"-"`
}

// RDNSRecordType represents Rancher DNS(RDNS) etcdv3 record type.
type RDNSRecordType struct {
	Type   string `json:"type,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type etcdv3Client struct {
	client *clientv3.Client
	ctx    context.Context
}

var _ RDNSClient = etcdv3Client{}

// NewRDNSProvider initializes a new Rancher DNS(RDNS) based Provider.
func NewRDNSProvider(config RDNSConfig) (*RDNSProvider, error) {
	client, err := newEtcdv3Client()
	if err != nil {
		return nil, err
	}
	domain := os.Getenv("RDNS_ROOT_DOMAIN")
	if domain == "" {
		return nil, errors.New("needed root domain environment")
	}
	return &RDNSProvider{
		client:       client,
		dryRun:       config.DryRun,
		domainFilter: config.DomainFilter,
		rootDomain:   domain,
	}, nil
}

// Records returns all DNS records found in Rancher DNS(RDNS) etcdv3 backend. Depending on the record fields
// it may be mapped to one or two records of type A, TXT, A+TXT.
func (p RDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint

	rs, err := p.client.List(p.rootDomain)
	if err != nil {
		return nil, err
	}

	for _, r := range rs {
		domains := strings.Split(strings.TrimPrefix(r.Key, rdnsPrefix+"/"), "/")
		keyToDNSNameSplits(domains)
		dnsName := strings.Join(domains, ".")
		if !p.domainFilter.Match(dnsName) {
			continue
		}

		// only return rdnsMaxHosts at most
		if len(r.AggregationHosts) > 0 {
			if len(r.AggregationHosts) > rdnsMaxHosts {
				r.AggregationHosts = r.AggregationHosts[:rdnsMaxHosts]
			}
			ep := endpoint.NewEndpointWithTTL(
				dnsName,
				endpoint.RecordTypeA,
				endpoint.TTL(r.TTL),
				r.AggregationHosts...,
			)
			ep.Labels[rdnsOriginalLabel] = r.Text
			result = append(result, ep)
		}
		if r.Text != "" {
			ep := endpoint.NewEndpoint(
				dnsName,
				endpoint.RecordTypeTXT,
				r.Text,
			)
			result = append(result, ep)
		}
	}

	return result, nil
}

// ApplyChanges stores changes back to etcdv3 converting them to Rancher DNS(RDNS) format and aggregating A and TXT records.
func (p RDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	grouped := map[string][]*endpoint.Endpoint{}

	for _, ep := range changes.Create {
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}

	for _, ep := range changes.UpdateNew {
		if ep.RecordType == endpoint.RecordTypeA {
			// append useless domain records to the changes.Delete
			if err := p.filterAndRemoveUseless(ep, changes); err != nil {
				return err
			}
		}
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}

	for dnsName, group := range grouped {
		if !p.domainFilter.Match(dnsName) {
			log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", dnsName)
			continue
		}

		var rs []RDNSRecord

		for _, ep := range group {
			if ep.RecordType == endpoint.RecordTypeTXT {
				continue
			}
			for _, target := range ep.Targets {
				rs = append(rs, RDNSRecord{
					Host: target,
					Text: ep.Labels[rdnsOriginalLabel],
					Key:  keyFor(ep.DNSName) + "/" + formatKey(target),
					TTL:  uint32(ep.RecordTTL),
				})
			}
		}

		// Add the TXT attribute to the existing A record
		for _, ep := range group {
			if ep.RecordType != endpoint.RecordTypeTXT {
				continue
			}
			for i, r := range rs {
				if strings.Contains(r.Key, keyFor(ep.DNSName)) {
					r.Text = ep.Targets[0]
					rs[i] = r
				}
			}
		}

		for _, r := range rs {
			log.Infof("Add/set key %s to Host=%s, Text=%s, TTL=%d", r.Key, r.Host, r.Text, r.TTL)
			if !p.dryRun {
				err := p.client.Set(r)
				if err != nil {
					return err
				}
			}
		}
	}

	for _, ep := range changes.Delete {
		key := keyFor(ep.DNSName)
		log.Infof("Delete key %s", key)
		if !p.dryRun {
			err := p.client.Delete(key)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// filterAndRemoveUseless filter and remove useless records.
func (p *RDNSProvider) filterAndRemoveUseless(ep *endpoint.Endpoint, changes *plan.Changes) error {
	rs, err := p.client.Get(keyFor(ep.DNSName))
	if err != nil {
		return err
	}
	for _, r := range rs {
		exist := false
		for _, target := range ep.Targets {
			if strings.Contains(r.Key, formatKey(target)) {
				exist = true
				continue
			}
		}
		if !exist {
			ds := strings.Split(strings.TrimPrefix(r.Key, rdnsPrefix+"/"), "/")
			keyToDNSNameSplits(ds)
			changes.Delete = append(changes.Delete, &endpoint.Endpoint{
				DNSName: strings.Join(ds, "."),
			})
		}
	}
	return nil
}

// newEtcdv3Client is an etcdv3 client constructor.
func newEtcdv3Client() (RDNSClient, error) {
	cfg := &clientv3.Config{}

	endpoints := os.Getenv("ETCD_URLS")
	ca := os.Getenv("ETCD_CA_FILE")
	cert := os.Getenv("ETCD_CERT_FILE")
	key := os.Getenv("ETCD_KEY_FILE")
	name := os.Getenv("ETCD_TLS_SERVER_NAME")
	insecure := os.Getenv("ETCD_TLS_INSECURE")

	if endpoints == "" {
		endpoints = "http://localhost:2379"
	}

	urls := strings.Split(endpoints, ",")
	scheme := strings.ToLower(urls[0])[0:strings.Index(strings.ToLower(urls[0]), "://")]

	switch scheme {
	case "http":
		cfg.Endpoints = urls
	case "https":
		var certificates []tls.Certificate

		insecure = strings.ToLower(insecure)
		isInsecure := insecure == "true" || insecure == "yes" || insecure == "1"

		if ca != "" && key == "" || cert == "" && key != "" {
			return nil, errors.New("either both cert and key or none must be provided")
		}

		if cert != "" {
			cert, err := tls.LoadX509KeyPair(cert, key)
			if err != nil {
				return nil, fmt.Errorf("could not load TLS cert: %s", err)
			}
			certificates = append(certificates, cert)
		}

		config := &tls.Config{
			Certificates:       certificates,
			InsecureSkipVerify: isInsecure,
			ServerName:         name,
		}

		if ca != "" {
			roots := x509.NewCertPool()
			pem, err := ioutil.ReadFile(ca)
			if err != nil {
				return nil, fmt.Errorf("error reading %s: %s", ca, err)
			}
			ok := roots.AppendCertsFromPEM(pem)
			if !ok {
				return nil, fmt.Errorf("could not read root certs: %s", err)
			}
			config.RootCAs = roots
		}

		cfg.Endpoints = urls
		cfg.TLS = config
	default:
		return nil, errors.New("etcdv3 URLs must start with either http:// or https://")
	}

	c, err := clientv3.New(*cfg)
	if err != nil {
		return nil, err
	}

	return etcdv3Client{c, context.Background()}, nil
}

// Get return A records stored in etcdv3 stored anywhere under the given key (recursively).
func (c etcdv3Client) Get(key string) ([]RDNSRecord, error) {
	ctx, cancel := context.WithTimeout(c.ctx, rdnsTimeout)
	defer cancel()

	result, err := c.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	rs := make([]RDNSRecord, 0)
	for _, v := range result.Kvs {
		r := new(RDNSRecord)
		if err := json.Unmarshal(v.Value, r); err != nil {
			return nil, fmt.Errorf("%s: %s", v.Key, err.Error())
		}
		r.Key = string(v.Key)
		rs = append(rs, *r)
	}

	return rs, nil
}

// List return all records stored in etcdv3 stored anywhere under the given rootDomain (recursively).
func (c etcdv3Client) List(rootDomain string) ([]RDNSRecord, error) {
	ctx, cancel := context.WithTimeout(c.ctx, rdnsTimeout)
	defer cancel()

	path := keyFor(rootDomain)

	result, err := c.client.Get(ctx, path, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	return c.aggregationRecords(result)
}

// Set persists records data into etcdv3.
func (c etcdv3Client) Set(r RDNSRecord) error {
	ctx, cancel := context.WithTimeout(c.ctx, etcdTimeout)
	defer cancel()

	v, err := json.Marshal(&r)
	if err != nil {
		return err
	}

	if r.Text == "" && r.Host == "" {
		return nil
	}

	_, err = c.client.Put(ctx, r.Key, string(v))
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes record from etcdv3.
func (c etcdv3Client) Delete(key string) error {
	ctx, cancel := context.WithTimeout(c.ctx, etcdTimeout)
	defer cancel()

	_, err := c.client.Delete(ctx, key, clientv3.WithPrefix())
	return err
}

// aggregationRecords will aggregation multi A records under the given path.
// e.g. A: 1_1_1_1.xxx.lb.rancher.cloud & 2_2_2_2.sample.lb.rancher.cloud => sample.lb.rancher.cloud {"aggregation_hosts": ["1.1.1.1", "2.2.2.2"]}
// e.g. TXT: sample.lb.rancher.cloud => sample.lb.rancher.cloud => {"text": "xxx"}
func (c etcdv3Client) aggregationRecords(result *clientv3.GetResponse) ([]RDNSRecord, error) {
	var rs []RDNSRecord
	bx := make(map[RDNSRecordType]RDNSRecord)

	for _, n := range result.Kvs {
		r := new(RDNSRecord)
		if err := json.Unmarshal(n.Value, r); err != nil {
			return nil, fmt.Errorf("%s: %s", n.Key, err.Error())
		}

		r.Key = string(n.Key)

		if r.Host == "" && r.Text == "" {
			continue
		}

		if r.Host != "" {
			c := RDNSRecord{
				AggregationHosts: r.AggregationHosts,
				Host:             r.Host,
				Text:             r.Text,
				TTL:              r.TTL,
				Key:              r.Key,
			}
			n, isContinue := appendRecords(c, endpoint.RecordTypeA, bx, rs)
			if isContinue {
				continue
			}
			rs = n
		}

		if r.Text != "" && r.Host == "" {
			c := RDNSRecord{
				AggregationHosts: []string{},
				Host:             r.Host,
				Text:             r.Text,
				TTL:              r.TTL,
				Key:              r.Key,
			}
			n, isContinue := appendRecords(c, endpoint.RecordTypeTXT, bx, rs)
			if isContinue {
				continue
			}
			rs = n
		}
	}

	return rs, nil
}

// appendRecords append record to an array
func appendRecords(r RDNSRecord, dnsType string, bx map[RDNSRecordType]RDNSRecord, rs []RDNSRecord) ([]RDNSRecord, bool) {
	dnsName := keyToParentDNSName(r.Key)
	bt := RDNSRecordType{Domain: dnsName, Type: dnsType}
	if v, ok := bx[bt]; ok {
		// skip the TXT records if already added to record list.
		// append A record if dnsName already added to record list but not found the value.
		// the same record might be found in multiple etcdv3 nodes.
		if bt.Type == endpoint.RecordTypeA {
			exist := false
			for _, h := range v.AggregationHosts {
				if h == r.Host {
					exist = true
					break
				}
			}
			if !exist {
				for i, t := range rs {
					if !strings.HasPrefix(r.Key, t.Key) {
						continue
					}
					t.Host = ""
					t.AggregationHosts = append(t.AggregationHosts, r.Host)
					bx[bt] = t
					rs[i] = t
				}
			}
		}
		return rs, true
	}

	if bt.Type == endpoint.RecordTypeA {
		r.AggregationHosts = append(r.AggregationHosts, r.Host)
	}

	r.Key = rdnsPrefix + dnsNameToKey(dnsName)
	r.Host = ""
	bx[bt] = r
	rs = append(rs, r)
	return rs, false
}

// keyFor used to get a path as etcdv3 preferred.
// e.g. sample.lb.rancher.cloud => /rdnsv3/cloud/rancher/lb/sample
func keyFor(fqdn string) string {
	return rdnsPrefix + dnsNameToKey(fqdn)
}

// keyToParentDNSName used to get dnsName.
// e.g. /rdnsv3/cloud/rancher/lb/sample/xxx => xxx.sample.lb.rancher.cloud
// e.g. /rdnsv3/cloud/rancher/lb/sample/xxx/1_1_1_1 => xxx.sample.lb.rancher.cloud
func keyToParentDNSName(key string) string {
	ds := strings.Split(strings.TrimPrefix(key, rdnsPrefix+"/"), "/")
	keyToDNSNameSplits(ds)

	dns := strings.Join(ds, ".")
	prefix := strings.Split(dns, ".")[0]

	p := `^\d{1,3}_\d{1,3}_\d{1,3}_\d{1,3}$`
	m, _ := regexp.MatchString(p, prefix)
	if prefix != "" && strings.Contains(prefix, "_") && m {
		// 1_1_1_1.xxx.sample.lb.rancher.cloud => xxx.sample.lb.rancher.cloud
		return strings.Join(strings.Split(dns, ".")[1:], ".")
	}

	return dns
}

// dnsNameToKey used to convert domain to a path as etcdv3 preferred.
// e.g. sample.lb.rancher.cloud => /cloud/rancher/lb/sample
func dnsNameToKey(domain string) string {
	ss := strings.Split(domain, ".")
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return "/" + strings.Join(ss, "/")
}

// keyToDNSNameSplits used to reverse etcdv3 path to domain splits.
// e.g. /cloud/rancher/lb/sample => [sample lb rancher cloud]
func keyToDNSNameSplits(ss []string) {
	for i := 0; i < len(ss)/2; i++ {
		j := len(ss) - i - 1
		ss[i], ss[j] = ss[j], ss[i]
	}
}

// formatKey used to format a key as etcdv3 preferred
// e.g. 1.1.1.1 => 1_1_1_1
// e.g. sample.lb.rancher.cloud => sample_lb_rancher_cloud
func formatKey(key string) string {
	return strings.Replace(key, ".", "_", -1)
}
