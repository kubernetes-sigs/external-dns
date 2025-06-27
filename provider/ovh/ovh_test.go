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

package ovh

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"testing"
	"time"

	"github.com/maxatome/go-testdeep/td"
	"github.com/miekg/dns"
	"github.com/ovh/go-ovh/ovh"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/ratelimit"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockOvhClient struct {
	mock.Mock
}

func (c *mockOvhClient) PostWithContext(ctx context.Context, endpoint string, input interface{}, output interface{}) error {
	stub := c.Called(endpoint, input)
	data, err := json.Marshal(stub.Get(0))
	if err != nil {
		return err
	}
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockOvhClient) PutWithContext(ctx context.Context, endpoint string, input interface{}, output interface{}) error {
	stub := c.Called(endpoint, input)
	data, err := json.Marshal(stub.Get(0))
	if err != nil {
		return err
	}
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockOvhClient) GetWithContext(ctx context.Context, endpoint string, output interface{}) error {
	stub := c.Called(endpoint)
	data, err := json.Marshal(stub.Get(0))
	if err != nil {
		return err
	}
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockOvhClient) DeleteWithContext(ctx context.Context, endpoint string, output interface{}) error {
	stub := c.Called(endpoint)
	data, err := json.Marshal(stub.Get(0))
	if err != nil {
		return err
	}
	json.Unmarshal(data, output)
	return stub.Error(1)
}

type mockDnsClient struct {
	mock.Mock
}

func (c *mockDnsClient) ExchangeContext(ctx context.Context, m *dns.Msg, addr string) (*dns.Msg, time.Duration, error) {
	args := c.Called(ctx, m, addr)

	msg := args.Get(0).(*dns.Msg)
	err := args.Error(1)

	return msg, time.Duration(0), err
}

func TestOvhZones(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{
		client:         client,
		apiRateLimiter: ratelimit.New(10),
		domainFilter:   endpoint.NewDomainFilter([]string{"com"}),
		cacheInstance:  cache.New(cache.NoExpiration, cache.NoExpiration),
		dnsClient:      new(mockDnsClient),
	}

	// Basic zones
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.com", "example.net"}, nil).Once()
	domains, err := provider.zones(t.Context())
	assert.NoError(err)
	assert.Contains(domains, "example.com")
	assert.NotContains(domains, "example.net")
	client.AssertExpectations(t)

	// Error on getting zones
	client.On("GetWithContext", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	domains, err = provider.zones(t.Context())
	assert.Error(err)
	assert.Nil(domains)
	client.AssertExpectations(t)
}

func TestOvhZoneRecords(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration), dnsClient: nil, UseCache: true}

	// Basic zones records
	t.Log("Basic zones records")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(ovhSoa{Server: "ns.example.org.", Serial: 2022090901}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return([]uint64{24, 42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/42").Return(ovhRecord{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	zones, records, err := provider.zonesRecords(t.Context())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, {ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}})
	client.AssertExpectations(t)

	// Error on getting zones list
	t.Log("Error on getting zones list")
	client.On("GetWithContext", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(t.Context())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone SOA
	t.Log("Error on getting zone SOA")
	provider.cacheInstance = cache.New(cache.NoExpiration, cache.NoExpiration)
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(t.Context())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone records
	t.Log("Error on getting zone records")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(ovhSoa{Server: "ns.example.org.", Serial: 2022090902}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(t.Context())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone record detail
	t.Log("Error on getting zone record detail")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(ovhSoa{Server: "ns.example.org.", Serial: 2022090902}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return([]uint64{42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/42").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(t.Context())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)
}

func TestOvhZoneRecordsCache(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	dnsClient := new(mockDnsClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration), dnsClient: dnsClient, UseCache: true}

	// First call, cache miss
	t.Log("First call, cache miss")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(ovhSoa{Server: "ns.example.org.", Serial: 2022090901}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return([]uint64{24, 42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/42").Return(ovhRecord{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()

	zones, records, err := provider.zonesRecords(t.Context())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, {ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}})
	client.AssertExpectations(t)
	dnsClient.AssertExpectations(t)

	// reset mock
	client = new(mockOvhClient)
	dnsClient = new(mockDnsClient)
	provider.client, provider.dnsClient = client, dnsClient

	// second call, cache hit
	t.Log("second call, cache hit")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	dnsClient.On("ExchangeContext", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("*dns.Msg"), "ns.example.org:53").
		Return(&dns.Msg{Answer: []dns.RR{&dns.SOA{Serial: 2022090901}}}, nil)
	zones, records, err = provider.zonesRecords(t.Context())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, {ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}})
	client.AssertExpectations(t)
	dnsClient.AssertExpectations(t)

	// reset mock
	client = new(mockOvhClient)
	dnsClient = new(mockDnsClient)
	provider.client, provider.dnsClient = client, dnsClient

	// third call, cache out of date
	t.Log("third call, cache out of date")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	dnsClient.On("ExchangeContext", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("*dns.Msg"), "ns.example.org:53").
		Return(&dns.Msg{Answer: []dns.RR{&dns.SOA{Serial: 2022090902}}}, nil)
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(ovhSoa{Server: "ns.example.org.", Serial: 2022090902}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return([]uint64{24}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()

	zones, records, err = provider.zonesRecords(t.Context())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}})
	client.AssertExpectations(t)
	dnsClient.AssertExpectations(t)

	// reset mock
	client = new(mockOvhClient)
	dnsClient = new(mockDnsClient)
	provider.client, provider.dnsClient = client, dnsClient

	// fourth call, cache hit
	t.Log("fourth call, cache hit")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	dnsClient.On("ExchangeContext", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("*dns.Msg"), "ns.example.org:53").
		Return(&dns.Msg{Answer: []dns.RR{&dns.SOA{Serial: 2022090902}}}, nil)

	zones, records, err = provider.zonesRecords(t.Context())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}})
	client.AssertExpectations(t)
	dnsClient.AssertExpectations(t)

	// reset mock
	client = new(mockOvhClient)
	dnsClient = new(mockDnsClient)
	provider.client, provider.dnsClient = client, dnsClient

	// fifth call, dns issue
	t.Log("fourth call, cache hit")
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	dnsClient.On("ExchangeContext", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("*dns.Msg"), "ns.example.org:53").
		Return(&dns.Msg{Answer: []dns.RR{}}, errors.New("dns issue"))
	client.On("GetWithContext", "/domain/zone/example.org/soa").Return(ovhSoa{Server: "ns.example.org.", Serial: 2022090903}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return([]uint64{24, 42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/42").Return(ovhRecord{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()

	zones, records, err = provider.zonesRecords(t.Context())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, {ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "NS", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}})
	client.AssertExpectations(t)
	dnsClient.AssertExpectations(t)
}

func TestOvhRecords(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}

	// Basic zones records
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.org", "example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record").Return([]uint64{24, 42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.org/record/42").Return(ovhRecord{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "CNAME", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "www", TTL: 10, Target: "example.org."}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{24, 42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/24").Return(ovhRecord{ID: 24, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.43"}}}, nil).Once()
	endpoints, err := provider.Records(t.Context())
	assert.NoError(err)
	// Little fix for multi targets endpoint
	for _, endpoint := range endpoints {
		sort.Strings(endpoint.Targets)
	}
	assert.ElementsMatch(endpoints, []*endpoint.Endpoint{
		{DNSName: "example.org", RecordType: "A", RecordTTL: 10, Labels: endpoint.NewLabels(), Targets: []string{"203.0.113.42"}},
		{DNSName: "www.example.org", RecordType: "CNAME", RecordTTL: 10, Labels: endpoint.NewLabels(), Targets: []string{"example.org"}},
		{DNSName: "ovh.example.net", RecordType: "A", RecordTTL: 10, Labels: endpoint.NewLabels(), Targets: []string{"203.0.113.42", "203.0.113.43"}},
	})
	client.AssertExpectations(t)

	// Error getting zone
	client.On("GetWithContext", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	endpoints, err = provider.Records(t.Context())
	assert.Error(err)
	assert.Nil(endpoints)
	client.AssertExpectations(t)
}

func TestOvhComputeChanges(t *testing.T) {
	existingRecords := []ovhRecord{
		{
			ID:   1,
			Zone: "example.net",
			ovhRecordFields: ovhRecordFields{
				FieldType: "A",
				ovhRecordFieldUpdate: ovhRecordFieldUpdate{
					SubDomain: "",
					Target:    "203.0.113.42",
				},
			},
		},
	}

	changes := plan.Changes{
		UpdateOld: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", Targets: []string{"203.0.113.42"}},
		},
		UpdateNew: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", Targets: []string{"203.0.113.43", "203.0.113.42"}},
		},
	}

	provider := &OVHProvider{client: nil, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}
	ovhChanges, err := provider.computeSingleZoneChanges(t.Context(), "example.net", existingRecords, &changes)
	td.CmpNoError(t, err)
	td.Cmp(t, ovhChanges, []ovhChange{
		{
			Action: ovhCreate,
			ovhRecord: ovhRecord{
				Zone: "example.net",
				ovhRecordFields: ovhRecordFields{
					FieldType: "A",
					ovhRecordFieldUpdate: ovhRecordFieldUpdate{
						SubDomain: "",
						Target:    "203.0.113.43",
					},
				},
			},
		},
	})

}

func TestOvhRefresh(t *testing.T) {
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}

	// Basic zone refresh
	client.On("PostWithContext", "/domain/zone/example.net/refresh", nil).Return(nil, nil).Once()
	provider.refresh(t.Context(), "example.net")
	client.AssertExpectations(t)
}

func TestOvhNewChange(t *testing.T) {
	provider := &OVHProvider{client: nil, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}

	endpoints := []*endpoint.Endpoint{
		{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		{DNSName: "ovh2.example.net", RecordType: "CNAME", Targets: []string{"ovh.example.net"}},
		{DNSName: "test.example.org"},
	}

	// Create change
	changes, _ := provider.newOvhChangeCreateDelete(ovhCreate, endpoints, "example.net", []ovhRecord{})
	td.Cmp(t, changes, []ovhChange{
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: defaultTTL, Target: "203.0.113.43"}}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "CNAME", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh2", TTL: defaultTTL, Target: "ovh.example.net."}}}},
	})

	// Delete change
	endpoints = []*endpoint.Endpoint{
		{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.42", "203.0.113.42", "203.0.113.43"}},
	}
	records := []ovhRecord{
		{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", Target: "203.0.113.43"}}},
		{ID: 43, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", Target: "203.0.113.42"}}},
		{ID: 44, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", Target: "203.0.113.42"}}},
	}
	changes, _ = provider.newOvhChangeCreateDelete(ovhDelete, endpoints, "example.net", records)
	td.Cmp(t, changes, []ovhChange{
		{Action: ovhDelete, ovhRecord: ovhRecord{ID: 43, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: defaultTTL, Target: "203.0.113.42"}}}},
		{Action: ovhDelete, ovhRecord: ovhRecord{ID: 44, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: defaultTTL, Target: "203.0.113.42"}}}},
		{Action: ovhDelete, ovhRecord: ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: defaultTTL, Target: "203.0.113.43"}}}},
	})

	// Create change with CNAME relative
	endpoints = []*endpoint.Endpoint{
		{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		{DNSName: "ovh2.example.net", RecordType: "CNAME", Targets: []string{"ovh"}},
		{DNSName: "test.example.org"},
	}

	provider = &OVHProvider{client: nil, EnableCNAMERelativeTarget: true, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}
	changes, _ = provider.newOvhChangeCreateDelete(ovhCreate, endpoints, "example.net", []ovhRecord{})
	td.Cmp(t, changes, []ovhChange{
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: defaultTTL, Target: "203.0.113.43"}}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "CNAME", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh2", TTL: defaultTTL, Target: "ovh"}}}},
	})

	// Test with CNAME when target has already final dot
	endpoints = []*endpoint.Endpoint{
		{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		{DNSName: "ovh2.example.net", RecordType: "CNAME", Targets: []string{"ovh.example.com."}},
		{DNSName: "test.example.org"},
	}

	provider = &OVHProvider{client: nil, EnableCNAMERelativeTarget: false, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}
	changes, _ = provider.newOvhChangeCreateDelete(ovhCreate, endpoints, "example.net", []ovhRecord{})
	td.Cmp(t, changes, []ovhChange{
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: defaultTTL, Target: "203.0.113.43"}}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "CNAME", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh2", TTL: defaultTTL, Target: "ovh.example.com."}}}},
	})
}

func TestOvhApplyChanges(t *testing.T) {
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
		Delete: []*endpoint.Endpoint{
			{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		},
	}

	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.43"}}}, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/record", ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}).Return(nil, nil).Once()
	client.On("DeleteWithContext", "/domain/zone/example.net/record/42").Return(nil, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/refresh", nil).Return(nil, nil).Once()

	_, err := provider.Records(t.Context())
	td.CmpNoError(t, err)
	// Basic changes
	td.CmpNoError(t, provider.ApplyChanges(t.Context(), &changes))
	client.AssertExpectations(t)

	// Apply change failed
	client = new(mockOvhClient)
	provider.client = client
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{}, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/record", ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}).Return(nil, ovh.ErrAPIDown).Once()

	_, err = provider.Records(t.Context())
	td.CmpNoError(t, err)
	td.CmpError(t, provider.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
	}))
	client.AssertExpectations(t)

	// Refresh failed
	client = new(mockOvhClient)
	provider.client = client
	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{}, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/record", ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}).Return(nil, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/refresh", nil).Return(nil, ovh.ErrAPIDown).Once()

	_, err = provider.Records(t.Context())
	td.CmpNoError(t, err)
	td.CmpError(t, provider.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
	}))
	client.AssertExpectations(t)

	// Test Dry-Run
	client = new(mockOvhClient)
	provider = &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration), DryRun: true}
	changes = plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
		Delete: []*endpoint.Endpoint{
			{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		},
	}

	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.43"}}}, nil).Once()

	_, err = provider.Records(t.Context())
	td.CmpNoError(t, err)
	td.CmpNoError(t, provider.ApplyChanges(t.Context(), &changes))
	client.AssertExpectations(t)

	// Test Update
	client = new(mockOvhClient)
	provider = &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration), DryRun: false}
	changes = plan.Changes{
		UpdateOld: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
		UpdateNew: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.43"}},
		},
	}

	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("PutWithContext", "/domain/zone/example.net/record/42", ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.43"}).Return(nil, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/refresh", nil).Return(nil, nil).Once()

	_, err = provider.Records(t.Context())
	td.CmpNoError(t, err)
	td.CmpNoError(t, provider.ApplyChanges(t.Context(), &changes))
	client.AssertExpectations(t)

	// Test Update DryRun
	client = new(mockOvhClient)
	provider = &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration), DryRun: true}
	changes = plan.Changes{
		UpdateOld: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
		UpdateNew: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.43"}},
		},
	}

	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{42}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()

	_, err = provider.Records(t.Context())
	td.CmpNoError(t, err)
	td.CmpNoError(t, provider.ApplyChanges(t.Context(), &changes))
	client.AssertExpectations(t)

	// Test Update 2 records => 1 record
	client = new(mockOvhClient)
	provider = &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration), DryRun: false}
	changes = plan.Changes{
		UpdateOld: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42", "203.0.113.43"}},
		},
		UpdateNew: []*endpoint.Endpoint{
			{DNSName: "example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.43"}},
		},
	}

	client.On("GetWithContext", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record").Return([]uint64{42, 43}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.42"}}}, nil).Once()
	client.On("GetWithContext", "/domain/zone/example.net/record/43").Return(ovhRecord{ID: 43, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "", TTL: 10, Target: "203.0.113.43"}}}, nil).Once()
	client.On("DeleteWithContext", "/domain/zone/example.net/record/42").Return(nil, nil).Once()
	client.On("PostWithContext", "/domain/zone/example.net/refresh", nil).Return(nil, nil).Once()

	_, err = provider.Records(t.Context())
	td.CmpNoError(t, err)
	td.CmpNoError(t, provider.ApplyChanges(t.Context(), &changes))
	client.AssertExpectations(t)
}

func TestOvhChange(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10), cacheInstance: cache.New(cache.NoExpiration, cache.NoExpiration)}

	// Record creation
	client.On("PostWithContext", "/domain/zone/example.net/record", ovhRecordFields{ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh"}}).Return(nil, nil).Once()
	assert.NoError(provider.change(t.Context(), ovhChange{
		Action:    ovhCreate,
		ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh"}}},
	}))
	client.AssertExpectations(t)

	// Record deletion
	client.On("DeleteWithContext", "/domain/zone/example.net/record/42").Return(nil, nil).Once()
	assert.NoError(provider.change(t.Context(), ovhChange{
		Action:    ovhDelete,
		ovhRecord: ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh"}}},
	}))
	client.AssertExpectations(t)

	// Record deletion error
	assert.Error(provider.change(t.Context(), ovhChange{
		Action:    ovhDelete,
		ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh"}}},
	}))
	client.AssertExpectations(t)
}

func TestOvhRecordString(t *testing.T) {
	record := ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{FieldType: "A", ovhRecordFieldUpdate: ovhRecordFieldUpdate{SubDomain: "ovh", TTL: 10, Target: "203.0.113.42"}}}

	td.Cmp(t, record.String(), "record#24: A | ovh => 203.0.113.42 (10)")
}

func TestNewOvhProvider(t *testing.T) {
	domainFilter := &endpoint.DomainFilter{}
	_, err := NewOVHProvider(t.Context(), domainFilter, "ovh-eu", 20, false, true)
	td.CmpError(t, err)

	t.Setenv("OVH_APPLICATION_KEY", "aaaaaa")
	t.Setenv("OVH_APPLICATION_SECRET", "bbbbbb")
	t.Setenv("OVH_CONSUMER_KEY", "cccccc")

	_, err = NewOVHProvider(t.Context(), domainFilter, "ovh-eu", 20, false, true)
	td.CmpNoError(t, err)
}
