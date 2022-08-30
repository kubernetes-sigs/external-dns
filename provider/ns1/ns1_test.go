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

package ns1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/data"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type MockNS1DomainClient struct {
	mock.Mock
}

func (m *MockNS1DomainClient) GetRecord(zone string, domain string, t string) (*dns.Record, *http.Response, error) {
	args := m.Called(zone, domain, t)
	r1 := args.Get(0).(*dns.Record)
	r2 := args.Get(1).(*http.Response)
	return r1, r2, args.Error(2)
}

func (m *MockNS1DomainClient) CreateRecord(r *dns.Record) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1DomainClient) DeleteRecord(zone string, domain string, t string) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1DomainClient) UpdateRecord(r *dns.Record) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1DomainClient) GetZone(zone string) (*dns.Zone, *http.Response, error) {
	r := &dns.ZoneRecord{
		Domain:   "test.foo.com",
		ShortAns: []string{"2.2.2.2"},
		TTL:      3600,
		Type:     "A",
		ID:       "123456789abcdefghijklmno",
	}
	z := &dns.Zone{
		Zone:    "foo.com",
		Records: []*dns.ZoneRecord{r},
		TTL:     3600,
		ID:      "12345678910111213141516a",
	}

	if zone == "foo.com" {
		return z, nil, nil
	}
	return nil, nil, nil
}

func (m *MockNS1DomainClient) ListZones() ([]*dns.Zone, *http.Response, error) {
	zones := []*dns.Zone{
		{Zone: "foo.com", ID: "12345678910111213141516a"},
		{Zone: "bar.com", ID: "12345678910111213141516b"},
	}
	return zones, nil, nil
}

type MockNS1GetZoneFail struct{}

func (m *MockNS1GetZoneFail) GetRecord(zone string, domain string, t string) (*dns.Record, *http.Response, error) {
	return nil, nil, nil
}

func (m *MockNS1GetZoneFail) CreateRecord(r *dns.Record) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1GetZoneFail) DeleteRecord(zone string, domain string, t string) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1GetZoneFail) UpdateRecord(r *dns.Record) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1GetZoneFail) GetZone(zone string) (*dns.Zone, *http.Response, error) {
	return nil, nil, api.ErrZoneMissing
}

func (m *MockNS1GetZoneFail) ListZones() ([]*dns.Zone, *http.Response, error) {
	zones := []*dns.Zone{
		{Zone: "foo.com", ID: "12345678910111213141516a"},
		{Zone: "bar.com", ID: "12345678910111213141516b"},
	}
	return zones, nil, nil
}

type MockNS1ListZonesFail struct{}

func (m *MockNS1ListZonesFail) GetRecord(zone string, domain string, t string) (*dns.Record, *http.Response, error) {
	return nil, nil, nil
}

func (m *MockNS1ListZonesFail) CreateRecord(r *dns.Record) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1ListZonesFail) DeleteRecord(zone string, domain string, t string) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1ListZonesFail) UpdateRecord(r *dns.Record) (*http.Response, error) {
	return nil, nil
}

func (m *MockNS1ListZonesFail) GetZone(zone string) (*dns.Zone, *http.Response, error) {
	return &dns.Zone{}, nil, nil
}

func (m *MockNS1ListZonesFail) ListZones() ([]*dns.Zone, *http.Response, error) {
	return nil, nil, fmt.Errorf("no zones available")
}

func TestNS1Records(t *testing.T) {
	mock := &MockNS1DomainClient{}
	provider := &NS1Provider{
		client:        mock,
		domainFilter:  endpoint.NewDomainFilter([]string{"foo.com."}),
		zoneIDFilter:  provider.NewZoneIDFilter([]string{""}),
		minTTLSeconds: 3600,
		OwnerID:       "testOwner",
	}
	ctx := context.Background()

	mock.On("GetRecord", "foo.com", "test.foo.com", "A").Return(&dns.Record{
		Answers: []*dns.Answer{{
			Meta: &data.Meta{
				Note: "ownerId:testOwner",
			},
		}},
	}, &http.Response{}, nil)

	records, err := provider.Records(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, len(records))

	// make sure it only returns owned anwswers
	mock2 := &MockNS1DomainClient{}
	provider.client = mock2
	mock2.On("GetRecord", "foo.com", "test.foo.com", "A").Return(&dns.Record{
		Answers: []*dns.Answer{{
			Meta: &data.Meta{
				Note: "ownerId:testOwner2",
			},
		}},
	}, &http.Response{}, nil)
	records, err = provider.Records(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, len(records))

	provider.client = &MockNS1GetZoneFail{}
	_, err = provider.Records(ctx)
	require.Error(t, err)

	provider.client = &MockNS1ListZonesFail{}
	_, err = provider.Records(ctx)
	require.Error(t, err)
}

func TestNewNS1Provider(t *testing.T) {
	_ = os.Setenv("NS1_APIKEY", "xxxxxxxxxxxxxxxxx")
	testNS1Config := NS1Config{
		DomainFilter: endpoint.NewDomainFilter([]string{"foo.com."}),
		ZoneIDFilter: provider.NewZoneIDFilter([]string{""}),
		DryRun:       false,
	}
	_, err := NewNS1Provider(testNS1Config)
	require.NoError(t, err)

	_ = os.Unsetenv("NS1_APIKEY")
	_, err = NewNS1Provider(testNS1Config)
	require.Error(t, err)
}

func TestNS1Zones(t *testing.T) {
	provider := &NS1Provider{
		client:       &MockNS1DomainClient{},
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com."}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	zones, err := provider.zonesFiltered()
	require.NoError(t, err)

	validateNS1Zones(t, zones, []*dns.Zone{
		{Zone: "foo.com"},
	})
}

func validateNS1Zones(t *testing.T, zones []*dns.Zone, expected []*dns.Zone) {
	require.Len(t, zones, len(expected))

	for i, zone := range zones {
		assert.Equal(t, expected[i].Zone, zone.Zone)
	}
}

func TestNS1BuildRecord(t *testing.T) {
	change := &ns1Change{
		Action: ns1Create,
		Endpoint: &endpoint.Endpoint{
			DNSName:    "new",
			Targets:    endpoint.Targets{"target"},
			RecordType: "A",
		},
	}

	provider := &NS1Provider{
		client:        &MockNS1DomainClient{},
		domainFilter:  endpoint.NewDomainFilter([]string{"foo.com."}),
		zoneIDFilter:  provider.NewZoneIDFilter([]string{""}),
		minTTLSeconds: 300,
		OwnerID:       "testOwner",
	}

	record := provider.ns1BuildRecord("foo.com", change)
	assert.Equal(t, "foo.com", record.Zone)
	assert.Equal(t, "new.foo.com", record.Domain)
	assert.Equal(t, 300, record.TTL)
	assert.Equal(t, "ownerId:testOwner", record.Answers[0].Meta.Note)

	changeWithTTL := &ns1Change{
		Action: ns1Create,
		Endpoint: &endpoint.Endpoint{
			DNSName:    "new-b",
			Targets:    endpoint.Targets{"target"},
			RecordType: "A",
			RecordTTL:  3600,
		},
	}
	record = provider.ns1BuildRecord("foo.com", changeWithTTL)
	assert.Equal(t, "foo.com", record.Zone)
	assert.Equal(t, "new-b.foo.com", record.Domain)
	assert.Equal(t, 3600, record.TTL)
	assert.Equal(t, "ownerId:testOwner", record.Answers[0].Meta.Note)

	changeWithWeight := &ns1Change{
		Action: ns1Create,
		Endpoint: &endpoint.Endpoint{
			DNSName:    "new-c",
			Targets:    endpoint.Targets{"target"},
			RecordType: "A",
			ProviderSpecific: endpoint.ProviderSpecific{{
				Name:  "weight",
				Value: "80",
			}},
		},
	}

	record = provider.ns1BuildRecord("foo.com", changeWithWeight)
	assert.Equal(t, "ownerId:testOwner", record.Answers[0].Meta.Note)
	assert.Equal(t, "80", record.Answers[0].Meta.Weight)
}

func TestNS1ApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	mock := &MockNS1DomainClient{}
	provider := &NS1Provider{
		client: mock,
	}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "new.foo.com", Targets: endpoint.Targets{"target"}, RecordType: "A"},
		{DNSName: "new.subdomain.bar.com", Targets: endpoint.Targets{"target"}, RecordType: "A"},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test.foo.com", Targets: endpoint.Targets{"target"}, RecordType: "A"}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test.foo.com", Targets: endpoint.Targets{"target-new"}, RecordType: "A"}}

	mock.On("GetRecord", "foo.com", "new.foo.com", "A").Return(&dns.Record{}, &http.Response{}, nil)
	mock.On("GetRecord", "bar.com", "new.subdomain.bar.com", "A").Return(&dns.Record{}, &http.Response{}, nil)
	mock.On("GetRecord", "foo.com", "test.foo.com", "A").Return(&dns.Record{}, &http.Response{}, nil)

	err := provider.ApplyChanges(context.Background(), changes)
	require.NoError(t, err)

	// empty changes
	changes.Create = []*endpoint.Endpoint{}
	changes.Delete = []*endpoint.Endpoint{}
	changes.UpdateNew = []*endpoint.Endpoint{}
	err = provider.ApplyChanges(context.Background(), changes)
	require.NoError(t, err)
}

func TestNewNS1Changes(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "testa.foo.com",
			Targets:    endpoint.Targets{"target-old"},
			RecordType: "A",
		},
		{
			DNSName:    "testba.bar.com",
			Targets:    endpoint.Targets{"target-new"},
			RecordType: "A",
		},
	}
	expected := []*ns1Change{
		{
			Action:   "ns1Create",
			Endpoint: endpoints[0],
		},
		{
			Action:   "ns1Create",
			Endpoint: endpoints[1],
		},
	}
	changes := newNS1Changes("ns1Create", endpoints)
	require.Len(t, changes, len(expected))
	assert.Equal(t, expected, changes)
}

func TestNewNS1ChangesByZone(t *testing.T) {
	provider := &NS1Provider{
		client: &MockNS1DomainClient{},
	}
	zones, _ := provider.zonesFiltered()
	changeSets := []*ns1Change{
		{
			Action: "ns1Create",
			Endpoint: &endpoint.Endpoint{
				DNSName:    "new.foo.com",
				Targets:    endpoint.Targets{"target"},
				RecordType: "A",
			},
		},
		{
			Action: "ns1Create",
			Endpoint: &endpoint.Endpoint{
				DNSName:    "unrelated.bar.com",
				Targets:    endpoint.Targets{"target"},
				RecordType: "A",
			},
		},
		{
			Action: "ns1Delete",
			Endpoint: &endpoint.Endpoint{
				DNSName:    "test.foo.com",
				Targets:    endpoint.Targets{"target"},
				RecordType: "A",
			},
		},
		{
			Action: "ns1Update",
			Endpoint: &endpoint.Endpoint{
				DNSName:    "test.foo.com",
				Targets:    endpoint.Targets{"target-new"},
				RecordType: "A",
			},
		},
	}

	changes := ns1ChangesByZone(zones, changeSets)
	assert.Len(t, changes["bar.com"], 1)
	assert.Len(t, changes["foo.com"], 3)
}

func TestOwnerNote(t *testing.T) {
	ownerId := "cluster1"
	note := ownerNote(ownerId)
	assert.Equal(t, "ownerId:cluster1", note)
}

func TestCheckOwnerNote(t *testing.T) {
	metaNote := "ownerId:cluster1"

	check1 := checkOwnerNote("cluster1", metaNote)
	check2 := checkOwnerNote("cluster2", metaNote)

	assert.True(t, check1)
	assert.False(t, check2)
}

func TestReconcileRecordChanges(t *testing.T) {
	table := []struct {
		message        string
		record         *dns.Record
		action         string
		ns1Record      *dns.Record
		ns1Error       error
		expectedRecord *dns.Record
		expectedAction string
	}{{
		message: "ns1Create with ErrRecordMissing from ns1 should trigger ns1Create",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Error: api.ErrRecordMissing,
		action:   ns1Create,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		expectedAction: ns1Create,
	}, {
		message: "ns1Create with existing records from ns1 should trigger ns1Update by adding new targets",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a2",
				Meta: &data.Meta{
					Note: "ownerId:cluster2",
				},
			}},
		},
		action: ns1Create,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a1",
					Meta: &data.Meta{
						Note: "ownerId:cluster1",
					},
				},
				{
					ID: "a2",
					Meta: &data.Meta{
						Note: "ownerId:cluster2",
					},
				},
			},
		},
		expectedAction: ns1Update,
	}, {
		message: "ns1Create with empty record from ns1 should trigger ns1Update by adding new targets",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Record: &dns.Record{
			Zone:    "zone1",
			Domain:  "domain1",
			Type:    "A",
			Answers: []*dns.Answer{},
		},
		action: ns1Create,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		expectedAction: ns1Update,
	}, {
		message: "ns1Update with non-owned targets in ns1 should trigger ns1Update by adding new targets",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a2",
				Meta: &data.Meta{
					Note: "ownerId:cluster2",
				},
			}},
		},
		action: ns1Create,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a1",
					Meta: &data.Meta{
						Note: "ownerId:cluster1",
					},
				},
				{
					ID: "a2",
					Meta: &data.Meta{
						Note: "ownerId:cluster2",
					},
				},
			},
		},
		expectedAction: ns1Update,
	}, {
		message: "ns1Update with additional non-owned targets in ns1 should trigger ns1Update by adjusting only owned targets",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a2",
					Meta: &data.Meta{
						Note: "ownerId:cluster1",
					},
				}, {
					ID: "a3",
					Meta: &data.Meta{
						Note: "ownerId:cluster2",
					},
				}},
		},
		action: ns1Update,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a1",
					Meta: &data.Meta{
						Note: "ownerId:cluster1",
					},
				},
				{
					ID: "a3",
					Meta: &data.Meta{
						Note: "ownerId:cluster2",
					},
				},
			},
		},
		expectedAction: ns1Update,
	}, {
		message: "ns1Delete with only owned targets in ns1 should trigger ns1Delete",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		action: ns1Delete,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a1",
					Meta: &data.Meta{
						Note: "ownerId:cluster1",
					},
				},
			},
		},
		expectedAction: ns1Delete,
	}, {
		message: "ns1Delete with non-owned targets in ns1 should trigger ns1Update by removing owned targets",
		record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{{
				ID: "a1",
				Meta: &data.Meta{
					Note: "ownerId:cluster1",
				},
			}},
		},
		ns1Record: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a1",
					Meta: &data.Meta{
						Note: "ownerId:cluster1",
					},
				},
				{
					ID: "a2",
					Meta: &data.Meta{
						Note: "ownerId:cluster2",
					},
				},
			},
		},
		action: ns1Delete,
		expectedRecord: &dns.Record{
			Zone:   "zone1",
			Domain: "domain1",
			Type:   "A",
			Answers: []*dns.Answer{
				{
					ID: "a2",
					Meta: &data.Meta{
						Note: "ownerId:cluster2",
					},
				},
			},
		},
		expectedAction: ns1Delete,
	}}

	for _, tt := range table {
		t.Run(tt.message, func(t *testing.T) {
			mock := &MockNS1DomainClient{}
			provider := &NS1Provider{
				client:  mock,
				OwnerID: "cluster1",
			}
			mock.On("GetRecord", tt.record.Zone, tt.record.Domain, tt.record.Type).Return(
				tt.ns1Record,
				&http.Response{},
				tt.ns1Error,
			)

			record, action := provider.reconcileRecordChanges(tt.record, tt.action)
			assert.NotNil(t, record)
			assert.Equal(t, action, tt.expectedAction)
			assert.Equal(t, record.Answers, tt.expectedRecord.Answers)
		})
	}
}
