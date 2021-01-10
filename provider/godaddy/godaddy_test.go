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

package godaddy

import (
	"context"
	"encoding/json"
	"sort"
	"testing"

	"github.com/ovh/go-ovh/ovh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockGoDaddyClient struct {
	mock.Mock
}

var (
	zoneNameExampleOrg string = "example.org"
	zoneNameExampleNet string = "example.net"
)

func (c *mockGoDaddyClient) Post(endpoint string, input interface{}, output interface{}) error {
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Patch(endpoint string, input interface{}, output interface{}) error {
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Put(endpoint string, input interface{}, output interface{}) error {
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Get(endpoint string, output interface{}) error {
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Delete(endpoint string, output interface{}) error {
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func TestGoDaddyZones(t *testing.T) {
	assert := assert.New(t)
	client := new(mockGoDaddyClient)
	provider := &GDProvider{
		client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"com"}),
	}

	// Basic zones
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: "example.com",
		},
		{
			Domain: "example.net",
		},
	}, nil).Once()

	domains, err := provider.zones()

	assert.NoError(err)
	assert.Contains(domains, "example.com")
	assert.NotContains(domains, "example.net")

	client.AssertExpectations(t)

	// Error on getting zones
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return(nil, ErrAPIDown).Once()
	domains, err = provider.zones()
	assert.Error(err)
	assert.Nil(domains)
	client.AssertExpectations(t)
}

func TestGoDaddyZoneRecords(t *testing.T) {
	assert := assert.New(t)
	client := new(mockGoDaddyClient)
	provider := &GDProvider{
		client: client,
	}

	// Basic zones records
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: zoneNameExampleNet,
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records").Return([]gdRecord{
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "NS",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
	}, nil).Once()

	zones, records, err := provider.zonesRecords(context.TODO())

	assert.NoError(err)

	assert.ElementsMatch(zones, []string{
		zoneNameExampleNet,
	})

	assert.ElementsMatch(records, []gdRecord{
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "NS",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
	})

	client.AssertExpectations(t)

	// Error on getting zones list
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return(nil, ErrAPIDown).Once()
	zones, records, err = provider.zonesRecords(context.TODO())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone records
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: zoneNameExampleNet,
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records").Return(nil, ErrAPIDown).Once()

	zones, records, err = provider.zonesRecords(context.TODO())

	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)

	// Error on getting zone record detail
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: zoneNameExampleNet,
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records").Return(nil, ErrAPIDown).Once()

	zones, records, err = provider.zonesRecords(context.TODO())
	assert.Error(err)
	assert.Nil(zones)
	assert.Nil(records)
	client.AssertExpectations(t)
}

func TestGoDaddyRecords(t *testing.T) {
	assert := assert.New(t)
	client := new(mockGoDaddyClient)
	provider := &GDProvider{
		client: client,
	}

	// Basic zones records
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: zoneNameExampleOrg,
		},
		{
			Domain: zoneNameExampleNet,
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.org/records").Return([]gdRecord{
		{
			zone: &zoneNameExampleOrg,
			gdRecordField: gdRecordField{
				Name: "@",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
		{
			zone: &zoneNameExampleOrg,
			gdRecordField: gdRecordField{
				Name: "www",
				Type: "CNAME",
				TTL:  10,
				Data: "example.org",
			},
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records").Return([]gdRecord{
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.43",
			},
		},
	}, nil).Once()

	endpoints, err := provider.Records(context.TODO())
	assert.NoError(err)

	// Little fix for multi targets endpoint
	for _, endpoint := range endpoints {
		sort.Strings(endpoint.Targets)
	}

	assert.ElementsMatch(endpoints, []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			RecordType: "A",
			RecordTTL:  10,
			Labels:     endpoint.NewLabels(),
			Targets: []string{
				"203.0.113.42",
			},
		},
		{
			DNSName:    "www.example.org",
			RecordType: "CNAME",
			RecordTTL:  10,
			Labels:     endpoint.NewLabels(),
			Targets: []string{
				"example.org",
			},
		},
		{
			DNSName:    "godaddy.example.net",
			RecordType: "A",
			RecordTTL:  10,
			Labels:     endpoint.NewLabels(),
			Targets: []string{
				"203.0.113.42",
				"203.0.113.43",
			},
		},
	})

	client.AssertExpectations(t)

	// Error getting zone
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return(nil, ovh.ErrAPIDown).Once()
	endpoints, err = provider.Records(context.TODO())
	assert.Error(err)
	assert.Nil(endpoints)
	client.AssertExpectations(t)
}

func TestGoDaddyNewChange(t *testing.T) {
	assert := assert.New(t)
	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    ".example.net",
			RecordType: "A",
			RecordTTL:  10, Targets: []string{
				"203.0.113.42",
			},
		},
		{
			DNSName:    "godaddy.example.net",
			RecordType: "A",
			Targets: []string{
				"203.0.113.43",
			},
		},
		{
			DNSName:    "godaddy2.example.net",
			RecordType: "CNAME",
			Targets: []string{
				"godaddy.example.net",
			},
		},
		{
			DNSName: "test.example.org",
		},
	}

	// Create change
	changes := newGoDaddyChange(gdCreate, endpoints, []string{"example.net"}, []gdRecord{})

	assert.ElementsMatch(changes, []gdChange{
		{
			Action: gdCreate,
			gdRecord: gdRecord{
				zone: &zoneNameExampleNet,
				gdRecordField: gdRecordField{
					Name: "",
					Type: "A",
					TTL:  10,
					Data: "203.0.113.42",
				},
			},
		},
		{
			Action: gdCreate,
			gdRecord: gdRecord{
				zone: &zoneNameExampleNet,
				gdRecordField: gdRecordField{
					Name: "godaddy",
					Type: "A",
					TTL:  gdDefaultTTL,
					Data: "203.0.113.43",
				},
			},
		},
		{
			Action: gdCreate,
			gdRecord: gdRecord{
				zone: &zoneNameExampleNet,
				gdRecordField: gdRecordField{
					Name: "godaddy2",
					Type: "CNAME",
					TTL:  gdDefaultTTL,
					Data: "godaddy.example.net.",
				},
			},
		}})

	// Delete change
	endpoints = []*endpoint.Endpoint{
		{
			DNSName:    "godaddy.example.net",
			RecordType: "A",
			Targets: []string{
				"203.0.113.42",
			},
		},
	}

	records := []gdRecord{
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Type: "A",
				Name: "godaddy",
				Data: "203.0.113.42",
			},
		},
	}

	changes = newGoDaddyChange(gdDelete, endpoints, []string{
		zoneNameExampleNet,
	}, records)

	assert.ElementsMatch(changes, []gdChange{
		{
			Action: gdDelete,
			gdRecord: gdRecord{
				zone: &zoneNameExampleNet,
				gdRecordField: gdRecordField{
					Name: "godaddy",
					Type: "A",
					TTL:  gdDefaultTTL,
					Data: "203.0.113.42",
				},
			},
		},
	})
}

func TestGoDaddyApplyChanges(t *testing.T) {
	assert := assert.New(t)
	client := new(mockGoDaddyClient)

	provider := &GDProvider{
		client: client,
	}

	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    ".example.net",
				RecordType: "A",
				RecordTTL:  10,
				Targets: []string{
					"203.0.113.42",
				},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "godaddy.example.net",
				RecordType: "A",
				Targets: []string{
					"203.0.113.43",
				},
			},
		},
	}

	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: zoneNameExampleNet,
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records").Return([]gdRecord{
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.43",
			},
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records/A/goddady").Return([]gdRecord{
		{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.43",
			},
		},
	}, nil).Once()

	client.On("Patch", "/v1/domains/example.net/records", []gdRecordField{
		{
			Name: "@",
			Type: "A",
			TTL:  10,
			Data: "203.0.113.42",
		},
	}).Return(nil, nil).Once()

	client.On("Delete", "/v1/domains/example.net/records/A/@").Return(nil, nil).Once()

	// Basic changes
	assert.NoError(provider.ApplyChanges(context.TODO(), &changes))

	client.AssertExpectations(t)

	// Getting zones failed
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return(nil, ErrAPIDown).Once()
	assert.Error(provider.ApplyChanges(context.TODO(), &changes))
	client.AssertExpectations(t)

	// Apply change failed
	client.On("Get", "/v1/domains?statuses=ACTIVE").Return([]gdZone{
		{
			Domain: zoneNameExampleNet,
		},
	}, nil).Once()

	client.On("Get", "/v1/domains/example.net/records").Return([]gdRecord{}, nil).Once()

	client.On("Patch", "/v1/domains/example.net/records/A/godaddy", []gdRecordField{
		{
			Name: "",
			Type: "A",
			TTL:  10,
			Data: "203.0.113.42",
		},
	}).Return(nil, ErrAPIDown).Once()

	assert.Error(provider.ApplyChanges(context.TODO(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    ".example.net",
				RecordType: "A",
				RecordTTL:  10,
				Targets: []string{
					"203.0.113.42",
				},
			},
		},
	}))

	client.AssertExpectations(t)
}

func TestGoDaddyChange(t *testing.T) {
	assert := assert.New(t)
	client := new(mockGoDaddyClient)
	provider := &GDProvider{
		client: client,
	}

	// Record creation
	client.On("Patch", "/v1/domains/example.net/records", []gdRecordField{
		{
			Name: "godaddy",
			Type: "A",
			TTL:  10,
			Data: "203.0.113.42",
		},
	}).Return(nil, nil).Once()

	assert.NoError(provider.change(gdChange{
		Action: gdCreate,
		gdRecord: gdRecord{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
				TTL:  10,
				Data: "203.0.113.42",
			},
		},
	}))

	client.AssertExpectations(t)

	// Record deletion
	client.On("Delete", "/v1/domains/example.net/records/A/godaddy").Return(nil, nil).Once()

	assert.NoError(provider.change(gdChange{
		Action: gdDelete,
		gdRecord: gdRecord{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
			},
		},
	}))

	client.AssertExpectations(t)

	// Record deletion error
	assert.Error(provider.change(gdChange{
		Action: gdDelete,
		gdRecord: gdRecord{
			zone: &zoneNameExampleNet,
			gdRecordField: gdRecordField{
				Name: "godaddy",
				Type: "A",
			},
		},
	}))

	client.AssertExpectations(t)
}

func TestOGoDaddyCountTargets(t *testing.T) {
	cases := []struct {
		endpoints [][]*endpoint.Endpoint
		count     int
	}{
		{[][]*endpoint.Endpoint{{{DNSName: "godaddy.example.net", Targets: endpoint.Targets{"target"}}}}, 1},
		{[][]*endpoint.Endpoint{{{DNSName: "godaddy.example.net", Targets: endpoint.Targets{"target"}}, {DNSName: "godaddy.example.net", Targets: endpoint.Targets{"target"}}}}, 2},
		{[][]*endpoint.Endpoint{{{DNSName: "godaddy.example.net", Targets: endpoint.Targets{"target", "target", "target"}}}}, 3},
		{[][]*endpoint.Endpoint{{{DNSName: "godaddy.example.net", Targets: endpoint.Targets{"target", "target"}}}, {{DNSName: "godaddy.example.net", Targets: endpoint.Targets{"target", "target"}}}}, 4},
	}
	for _, test := range cases {
		count := countTargets(test.endpoints...)
		if count != test.count {
			t.Errorf("Wrong targets counts (Should be %d, get %d)", test.count, count)
		}
	}
}
