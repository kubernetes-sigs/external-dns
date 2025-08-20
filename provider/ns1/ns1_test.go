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
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type MockNS1DomainClient struct {
	mock.Mock
}

func (m *MockNS1DomainClient) CreateRecord(r *dns.Record) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1DomainClient) DeleteRecord(zone string, domain string, t string) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1DomainClient) UpdateRecord(r *dns.Record) (*http.Response, error) {
	return &http.Response{}, nil
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

func (m *MockNS1GetZoneFail) CreateRecord(_ *dns.Record) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1GetZoneFail) DeleteRecord(_ string, _ string, _ string) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1GetZoneFail) UpdateRecord(r *dns.Record) (*http.Response, error) {
	return &http.Response{}, nil
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

func (m *MockNS1ListZonesFail) CreateRecord(_ *dns.Record) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1ListZonesFail) DeleteRecord(_ string, _ string, _ string) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1ListZonesFail) UpdateRecord(r *dns.Record) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m *MockNS1ListZonesFail) GetZone(zone string) (*dns.Zone, *http.Response, error) {
	return &dns.Zone{}, &http.Response{}, nil
}

func (m *MockNS1ListZonesFail) ListZones() ([]*dns.Zone, *http.Response, error) {
	return nil, nil, fmt.Errorf("no zones available")
}

func TestNS1Records(t *testing.T) {
	provider := &NS1Provider{
		client:        &MockNS1DomainClient{},
		domainFilter:  endpoint.NewDomainFilter([]string{"foo.com."}),
		zoneIDFilter:  provider.NewZoneIDFilter([]string{""}),
		minTTLSeconds: 3600,
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	require.NoError(t, err)
	assert.Len(t, records, 1)

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
	}

	record := provider.ns1BuildRecord("foo.com", change)
	assert.Equal(t, "foo.com", record.Zone)
	assert.Equal(t, "new.foo.com", record.Domain)
	assert.Equal(t, 300, record.TTL)

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

	changes := provider.ns1ChangesByZone(zones, changeSets)
	assert.Len(t, changes["bar.com"], 1)
	assert.Len(t, changes["foo.com"], 3)
}

// helper: build a change
func change(action, name, rtype string) *ns1Change {
	return &ns1Change{
		Action: action,
		Endpoint: &endpoint.Endpoint{
			DNSName:    name,
			Targets:    endpoint.Targets{"target"},
			RecordType: rtype,
		},
	}
}

func TestNS1ChangesByZone_WithZoneOverrideForOneZone(t *testing.T) {
	provider := &NS1Provider{
		client:              &MockNS1DomainClient{},
		zoneHandleOverrides: normalizeOverrides(map[string]string{"foo.com": "foo-handle"}),
	}

	zones, _ := provider.zonesFiltered()
	changeSets := []*ns1Change{
		change(ns1Create, "new.foo.com", "A"),
		change(ns1Create, "unrelated.bar.com", "A"),
		change(ns1Delete, "test.foo.com", "A"),
		change(ns1Update, "test.foo.com", "A"),
	}

	changes := provider.ns1ChangesByZone(zones, changeSets)

	// bar.com unchanged; foo.com should be bucketed under the handle
	assert.Len(t, changes["bar.com"], 1, "bar.com should still use its FQDN key")
	_, hasFooFQDN := changes["foo.com"]
	assert.False(t, hasFooFQDN, "foo.com key should not exist when overridden")

	assert.Len(t, changes["foo-handle"], 3, "foo records should bucket under handle when overridden")
}

func TestNS1ChangesByZone_WithOverridesForBothZones(t *testing.T) {
	provider := &NS1Provider{
		client: &MockNS1DomainClient{},
		zoneHandleOverrides: normalizeOverrides(map[string]string{
			"foo.com": "corp-prod-zone",
			"bar.com": "bar-view-handle",
		}),
	}

	zones, _ := provider.zonesFiltered()
	changeSets := []*ns1Change{
		change(ns1Create, "new.foo.com", "A"),
		change(ns1Create, "unrelated.bar.com", "A"),
		change(ns1Delete, "test.foo.com", "A"),
		change(ns1Update, "test.foo.com", "A"),
	}

	changes := provider.ns1ChangesByZone(zones, changeSets)

	// Both zones should use their mapped handles; no FQDN keys present
	_, hasFooFQDN := changes["foo.com"]
	_, hasBarFQDN := changes["bar.com"]
	assert.False(t, hasFooFQDN, "foo.com key should not exist when overridden")
	assert.False(t, hasBarFQDN, "bar.com key should not exist when overridden")

	assert.Len(t, changes["corp-prod-zone"], 3)
	assert.Len(t, changes["bar-view-handle"], 1)
}

func TestNS1ChangesByZone_OverrideNormalizationAndSuffix(t *testing.T) {
	// Uppercase + trailing dot should normalize; override should still take effect.
	provider := &NS1Provider{
		client:              &MockNS1DomainClient{},
		zoneHandleOverrides: normalizeOverrides(map[string]string{"Foo.COM.": "FOO-HANDLE"}),
	}

	zones, _ := provider.zonesFiltered()
	changeSets := []*ns1Change{
		change(ns1Create, "new.foo.com", "A"),
		change(ns1Delete, "test.foo.com", "A"),
		change(ns1Update, "test.foo.com", "A"),
	}

	changes := provider.ns1ChangesByZone(zones, changeSets)

	_, hasFooFQDN := changes["foo.com"]
	assert.False(t, hasFooFQDN, "normalized override should suppress FQDN key")
	if _, ok := changes["foo-handle"]; ok {
		assert.Len(t, changes["foo-handle"], 3)
	} else {
		assert.Len(t, changes["FOO-HANDLE"], 3)
	}
}

func TestNS1ChangesByZone_IgnoresUnmatchedRecords(t *testing.T) {
	provider := &NS1Provider{
		client:              &MockNS1DomainClient{},
		zoneHandleOverrides: normalizeOverrides(map[string]string{"foo.com": "foo-handle"}),
	}

	zones, _ := provider.zonesFiltered()
	changeSets := []*ns1Change{
		change(ns1Create, "unknown.baz.com", "A"), // does not match foo.com or bar.com
	}

	changes := provider.ns1ChangesByZone(zones, changeSets)

	// Should still have exactly the zone keys for the provided zones, but no entries inside.
	if gs, ok := changes["foo-handle"]; ok {
		assert.Empty(t, gs)
	}
	if gs, ok := changes["bar.com"]; ok {
		assert.Empty(t, gs)
	}
}
