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
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type MockNS1DomainClient struct {
	mock.Mock
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
	provider := &NS1Provider{
		client:       &MockNS1DomainClient{},
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com."}),
		zoneIDFilter: NewZoneIDFilter([]string{""}),
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, len(records))

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
		ZoneIDFilter: NewZoneIDFilter([]string{""}),
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
		zoneIDFilter: NewZoneIDFilter([]string{""}),
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
	record := ns1BuildRecord("foo.com", change)
	assert.Equal(t, "foo.com", record.Zone)
	assert.Equal(t, "new.foo.com", record.Domain)
	assert.Equal(t, ns1DefaultTTL, record.TTL)

	changeWithTTL := &ns1Change{
		Action: ns1Create,
		Endpoint: &endpoint.Endpoint{
			DNSName:    "new-b",
			Targets:    endpoint.Targets{"target"},
			RecordType: "A",
			RecordTTL:  100,
		},
	}
	record = ns1BuildRecord("foo.com", changeWithTTL)
	assert.Equal(t, "foo.com", record.Zone)
	assert.Equal(t, "new-b.foo.com", record.Domain)
	assert.Equal(t, 100, record.TTL)
}

func TestNS1ApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	provider := &NS1Provider{
		client: &MockNS1DomainClient{},
	}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "new.foo.com", Targets: endpoint.Targets{"target"}},
		{DNSName: "new.subdomain.bar.com", Targets: endpoint.Targets{"target"}},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test.foo.com", Targets: endpoint.Targets{"target"}}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test.foo.com", Targets: endpoint.Targets{"target-new"}}}
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
