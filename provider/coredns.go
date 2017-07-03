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
	"container/list"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// skyDNSClient is an interface to work with SkyDNS service records in etcd
type skyDNSClient interface {
	GetServices(prefix string) ([]*Service, error)
	SaveService(value *Service) error
	DeleteService(key string) error
}

type coreDNSProvider struct {
	dryRun       bool
	domainFilter DomainFilter
	client       skyDNSClient
}

// Service represents SkyDNS/CoreDNS etcd record
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
	api etcd.KeysAPI
}

var _ skyDNSClient = etcdClient{}

// GetService return all Service records stored in etcd stored anywhere under the given key (recursively)
func (c etcdClient) GetServices(prefix string) ([]*Service, error) {
	var result []*Service
	opts := &etcd.GetOptions{Recursive: true}
	data, err := c.api.Get(context.Background(), prefix, opts)
	if err != nil {
		if etcd.IsKeyNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	queue := list.New()
	queue.PushFront(data.Node)
	for queueNode := queue.Front(); queueNode != nil; queueNode = queueNode.Next() {
		node := queueNode.Value.(*etcd.Node)
		if node.Dir {
			for _, childNode := range node.Nodes {
				queue.PushBack(childNode)
			}
			continue
		}
		service := &Service{}
		err = json.Unmarshal([]byte(node.Value), service)
		if err != nil {
			log.Error("Cannot parse JSON value ", node.Value)
			continue
		}
		service.Key = node.Key
		result = append(result, service)
	}
	return result, nil
}

// SaveService persists service data into etcd
func (c etcdClient) SaveService(service *Service) error {
	value, err := json.Marshal(&service)
	if err != nil {
		return err
	}
	_, err = c.api.Set(context.Background(), service.Key, string(value), nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteService deletes service record from etcd
func (c etcdClient) DeleteService(key string) error {
	_, err := c.api.Delete(context.Background(), key, nil)
	return err

}

//newETCDClient is an etcd client constructor
func newETCDClient(etcdURIs string) (skyDNSClient, error) {
	cfg := etcd.Config{Endpoints: strings.Split(etcdURIs, ",")}
	c, err := etcd.New(cfg)
	if err != nil {
		return nil, err
	}
	return etcdClient{etcd.NewKeysAPI(c)}, nil
}

// NewCoreDNSProvider is a CoreDNS provider constructor
func NewCoreDNSProvider(domainFilter DomainFilter, etcdURIs string, dryRun bool) (Provider, error) {
	client, err := newETCDClient(etcdURIs)
	if err != nil {
		return nil, err
	}
	return coreDNSProvider{
		client:       client,
		dryRun:       dryRun,
		domainFilter: domainFilter,
	}, nil
}

// Records returns all DNS records found in SkyDNS/CoreDNS etcd backend. Depending on the record fields
// it may be mapped to one or two records of type A, CNAME, TXT, A+TXT, CNAME+TXT
func (p coreDNSProvider) Records() ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	services, err := p.client.GetServices("/skydns")
	if err != nil {
		return nil, err
	}
	for _, service := range services {
		domains := strings.Split(strings.TrimPrefix(service.Key, "/skydns/"), "/")
		reverse(domains)
		dnsName := strings.Join(domains[service.TargetStrip:], ".")
		if !p.domainFilter.Match(dnsName) {
			continue
		}
		prefix := strings.Join(domains[:service.TargetStrip], ".")
		if service.Host != "" {
			ep := endpoint.NewEndpoint(
				dnsName,
				service.Host,
				guessRecordType(service.Host),
			)
			ep.Labels["originalText"] = service.Text
			ep.Labels["prefix"] = prefix
			result = append(result, ep)
		}
		if service.Text != "" {
			ep := endpoint.NewEndpoint(
				dnsName,
				service.Text,
				endpoint.RecordTypeTXT,
			)
			ep.Labels["prefix"] = prefix
			result = append(result, ep)
		}
	}
	return result, nil
}

// ApplyChanges stores changes back to etcd converting them to SkyDNS format and aggregating A/CNAME and TXT records
func (p coreDNSProvider) ApplyChanges(changes *plan.Changes) error {
	grouped := map[string][]*endpoint.Endpoint{}
	for _, ep := range changes.Create {
		grouped[ep.DNSName] = append(grouped[ep.DNSName], ep)
	}
	for _, ep := range changes.UpdateNew {
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
			prefix := ep.Labels["prefix"]
			if prefix == "" {
				prefix = fmt.Sprintf("%08x", rand.Int31())
			}
			service := Service{
				Host:        ep.Target,
				Text:        ep.Labels["originalText"],
				Key:         etcdKeyFor(prefix + "." + dnsName),
				TargetStrip: strings.Count(prefix, ".") + 1,
			}
			services = append(services, service)
		}
		index := 0
		for _, ep := range group {
			if ep.RecordType != "TXT" {
				continue
			}
			if index >= len(services) {
				prefix := ep.Labels["prefix"]
				if prefix == "" {
					prefix = fmt.Sprintf("%08x", rand.Int31())
				}
				services = append(services, Service{
					Key:         etcdKeyFor(prefix + "." + dnsName),
					TargetStrip: strings.Count(prefix, ".") + 1,
				})
			}
			services[index].Text = ep.Target
			index++
		}

		for i := index; index > 0 && i < len(services); i++ {
			services[i].Text = ""
		}

		for _, service := range services {
			log.Infof("Add/set key %s to Host=%s, Text=%s", service.Key, service.Host, service.Text)
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
		if ep.Labels["prefix"] != "" {
			dnsName = ep.Labels["prefix"] + "." + dnsName
		}
		key := etcdKeyFor(dnsName)
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

func guessRecordType(target string) string {
	if net.ParseIP(target) != nil {
		return endpoint.RecordTypeA
	}
	return endpoint.RecordTypeCNAME
}

func etcdKeyFor(dnsName string) string {
	domains := strings.Split(dnsName, ".")
	reverse(domains)
	return "/skydns/" + strings.Join(domains, "/")
}

func reverse(slice []string) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}
