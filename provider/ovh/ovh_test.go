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
	"sort"
	"testing"

	"github.com/ovh/go-ovh/ovh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/ratelimit"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockOvhClient struct {
	mock.Mock
}

func (c *mockOvhClient) Post(endpoint string, input interface{}, output interface{}) error {
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockOvhClient) Get(endpoint string, output interface{}) error {
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockOvhClient) Delete(endpoint string, output interface{}) error {
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func TestOvhZones(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{
		client:         client,
		apiRateLimiter: ratelimit.New(10),
		domainFilter:   endpoint.NewDomainFilter([]string{"com"}),
	}

	// Basic zones
	client.On("Get", "/domain/zone").Return([]string{"example.com", "example.net"}, nil).Once()
	domains, err := provider.zones()
	assert.NoError(err)
	assert.Contains(domains, "example.com")
	assert.NotContains(domains, "example.net")
	client.AssertExpectations(t)

	// Error on getting zones
	client.On("Get", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	domains, err = provider.zones()
	assert.Error(err)
	assert.Nil(domains)
	client.AssertExpectations(t)
}

func TestOvhZoneRecords(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10)}

	// Basic zones records
	client.On("Get", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record").Return([]uint64{24, 42}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "NS", TTL: 10, Target: "203.0.113.42"}}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record/42").Return(ovhRecord{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: 10, Target: "203.0.113.42"}}, nil).Once()
	zones, records, err := provider.zonesRecords(context.TODO())
	assert.NoError(err)
	assert.ElementsMatch(zones, []string{"example.org"})
	assert.ElementsMatch(records, []ovhRecord{{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: 10, Target: "203.0.113.42"}}, {ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "NS", TTL: 10, Target: "203.0.113.42"}}})
	client.AssertExpectations(t)

	// Error on getting zones list
	client.On("Get", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(context.TODO())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone records
	client.On("Get", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(context.TODO())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone record detail
	client.On("Get", "/domain/zone").Return([]string{"example.org"}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record").Return([]uint64{42}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record/42").Return(nil, ovh.ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(context.TODO())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)
}

func TestOvhRecords(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10)}

	// Basic zones records
	client.On("Get", "/domain/zone").Return([]string{"example.org", "example.net"}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record").Return([]uint64{24, 42}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record/24").Return(ovhRecord{ID: 24, Zone: "example.org", ovhRecordFields: ovhRecordFields{SubDomain: "", FieldType: "A", TTL: 10, Target: "203.0.113.42"}}, nil).Once()
	client.On("Get", "/domain/zone/example.org/record/42").Return(ovhRecord{ID: 42, Zone: "example.org", ovhRecordFields: ovhRecordFields{SubDomain: "www", FieldType: "CNAME", TTL: 10, Target: "example.org."}}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record").Return([]uint64{24, 42}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record/24").Return(ovhRecord{ID: 24, Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: 10, Target: "203.0.113.42"}}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: 10, Target: "203.0.113.43"}}, nil).Once()
	endpoints, err := provider.Records(context.TODO())
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
	client.On("Get", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	endpoints, err = provider.Records(context.TODO())
	assert.Error(err)
	assert.Nil(endpoints)
	client.AssertExpectations(t)
}

func TestOvhRefresh(t *testing.T) {
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10)}

	// Basic zone refresh
	client.On("Post", "/domain/zone/example.net/refresh", nil).Return(nil, nil).Once()
	provider.refresh("example.net")
	client.AssertExpectations(t)
}

func TestOvhNewChange(t *testing.T) {
	assert := assert.New(t)
	endpoints := []*endpoint.Endpoint{
		{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		{DNSName: "ovh2.example.net", RecordType: "CNAME", Targets: []string{"ovh.example.net"}},
		{DNSName: "test.example.org"},
	}

	// Create change
	changes := newOvhChange(ovhCreate, endpoints, []string{"example.net"}, []ovhRecord{})
	assert.ElementsMatch(changes, []ovhChange{
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "", FieldType: "A", TTL: 10, Target: "203.0.113.42"}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: ovhDefaultTTL, Target: "203.0.113.43"}}},
		{Action: ovhCreate, ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh2", FieldType: "CNAME", TTL: ovhDefaultTTL, Target: "ovh.example.net."}}},
	})

	// Delete change
	endpoints = []*endpoint.Endpoint{
		{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.42"}},
	}
	records := []ovhRecord{
		{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{FieldType: "A", SubDomain: "ovh", Target: "203.0.113.42"}},
	}
	changes = newOvhChange(ovhDelete, endpoints, []string{"example.net"}, records)
	assert.ElementsMatch(changes, []ovhChange{
		{Action: ovhDelete, ovhRecord: ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: ovhDefaultTTL, Target: "203.0.113.42"}}},
	})
}

func TestOvhApplyChanges(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10)}
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
		Delete: []*endpoint.Endpoint{
			{DNSName: "ovh.example.net", RecordType: "A", Targets: []string{"203.0.113.43"}},
		},
	}

	client.On("Get", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record").Return([]uint64{42}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record/42").Return(ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh", FieldType: "A", TTL: 10, Target: "203.0.113.43"}}, nil).Once()
	client.On("Post", "/domain/zone/example.net/record", ovhRecordFields{SubDomain: "", FieldType: "A", TTL: 10, Target: "203.0.113.42"}).Return(nil, nil).Once()
	client.On("Delete", "/domain/zone/example.net/record/42").Return(nil, nil).Once()
	client.On("Post", "/domain/zone/example.net/refresh", nil).Return(nil, nil).Once()

	// Basic changes
	assert.NoError(provider.ApplyChanges(context.TODO(), &changes))
	client.AssertExpectations(t)

	// Getting zones failed
	client.On("Get", "/domain/zone").Return(nil, ovh.ErrAPIDown).Once()
	assert.Error(provider.ApplyChanges(context.TODO(), &changes))
	client.AssertExpectations(t)

	// Apply change failed
	client.On("Get", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record").Return([]uint64{}, nil).Once()
	client.On("Post", "/domain/zone/example.net/record", ovhRecordFields{SubDomain: "", FieldType: "A", TTL: 10, Target: "203.0.113.42"}).Return(nil, ovh.ErrAPIDown).Once()
	assert.Error(provider.ApplyChanges(context.TODO(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
	}))
	client.AssertExpectations(t)

	// Refresh failed
	client.On("Get", "/domain/zone").Return([]string{"example.net"}, nil).Once()
	client.On("Get", "/domain/zone/example.net/record").Return([]uint64{}, nil).Once()
	client.On("Post", "/domain/zone/example.net/record", ovhRecordFields{SubDomain: "", FieldType: "A", TTL: 10, Target: "203.0.113.42"}).Return(nil, nil).Once()
	client.On("Post", "/domain/zone/example.net/refresh", nil).Return(nil, ovh.ErrAPIDown).Once()
	assert.Error(provider.ApplyChanges(context.TODO(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: ".example.net", RecordType: "A", RecordTTL: 10, Targets: []string{"203.0.113.42"}},
		},
	}))
	client.AssertExpectations(t)
}

func TestOvhChange(t *testing.T) {
	assert := assert.New(t)
	client := new(mockOvhClient)
	provider := &OVHProvider{client: client, apiRateLimiter: ratelimit.New(10)}

	// Record creation
	client.On("Post", "/domain/zone/example.net/record", ovhRecordFields{SubDomain: "ovh"}).Return(nil, nil).Once()
	assert.NoError(provider.change(ovhChange{
		Action:    ovhCreate,
		ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh"}},
	}))
	client.AssertExpectations(t)

	// Record deletion
	client.On("Delete", "/domain/zone/example.net/record/42").Return(nil, nil).Once()
	assert.NoError(provider.change(ovhChange{
		Action:    ovhDelete,
		ovhRecord: ovhRecord{ID: 42, Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh"}},
	}))
	client.AssertExpectations(t)

	// Record deletion error
	assert.Error(provider.change(ovhChange{
		Action:    ovhDelete,
		ovhRecord: ovhRecord{Zone: "example.net", ovhRecordFields: ovhRecordFields{SubDomain: "ovh"}},
	}))
	client.AssertExpectations(t)
}

func TestOvhCountTargets(t *testing.T) {
	cases := []struct {
		endpoints [][]*endpoint.Endpoint
		count     int
	}{
		{[][]*endpoint.Endpoint{{{DNSName: "ovh.example.net", Targets: endpoint.Targets{"target"}}}}, 1},
		{[][]*endpoint.Endpoint{{{DNSName: "ovh.example.net", Targets: endpoint.Targets{"target"}}, {DNSName: "ovh.example.net", Targets: endpoint.Targets{"target"}}}}, 2},
		{[][]*endpoint.Endpoint{{{DNSName: "ovh.example.net", Targets: endpoint.Targets{"target", "target", "target"}}}}, 3},
		{[][]*endpoint.Endpoint{{{DNSName: "ovh.example.net", Targets: endpoint.Targets{"target", "target"}}}, {{DNSName: "ovh.example.net", Targets: endpoint.Targets{"target", "target"}}}}, 4},
	}
	for _, test := range cases {
		count := countTargets(test.endpoints...)
		if count != test.count {
			t.Errorf("Wrong targets counts (Should be %d, get %d)", test.count, count)
		}
	}
}
