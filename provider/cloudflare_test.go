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
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockCloudFlareClient struct{}

func (m *mockCloudFlareClient) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareClient) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	if zoneID == "1234567890" {
		return []cloudflare.DNSRecord{
				{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to.", Type: endpoint.RecordTypeA, TTL: 120},
				{ID: "1231231233", Name: "foo.bar.com", TTL: 1}},
			nil
	}
	return nil, nil
}

func (m *mockCloudFlareClient) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareClient) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareClient) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareClient) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareClient) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{ID: "1234567890", Name: "ext-dns-test.zalando.to."}, {ID: "1234567891", Name: "foo.com."}}, nil
}

func (m *mockCloudFlareClient) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{
		Result: []cloudflare.Zone{
			{ID: "1234567890", Name: "ext-dns-test.zalando.to."},
			{ID: "1234567891", Name: "foo.com."}},
		ResultInfo: cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
	}, nil
}

type mockCloudFlareDNSRecordsFail struct{}

func (m *mockCloudFlareDNSRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareDNSRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, fmt.Errorf("can not get records from zone")
}
func (m *mockCloudFlareDNSRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareDNSRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareDNSRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareDNSRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudFlareDNSRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDNSRecordsFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{
		Result: []cloudflare.Zone{
			{ID: "1234567890", Name: "ext-dns-test.zalando.to."},
			{ID: "1234567891", Name: "foo.com."}},
		ResultInfo: cloudflare.ResultInfo{
			TotalPages: 1,
		},
	}, nil
}

type mockCloudFlareListZonesFail struct{}

func (m *mockCloudFlareListZonesFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareListZonesFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareListZonesFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareListZonesFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareListZonesFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudFlareListZonesFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareListZonesFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

func (m *mockCloudFlareListZonesFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{}, fmt.Errorf("no zones available")
}

func TestNewCloudFlareChanges(t *testing.T) {
	expect := []struct {
		Name string
		TTL  int
	}{
		{
			"CustomRecordTTL",
			120,
		},
		{
			"DefaultRecordTTL",
			1,
		},
	}
	endpoints := []*endpoint.Endpoint{
		{DNSName: "new", Targets: endpoint.Targets{"target"}, RecordTTL: 120},
		{DNSName: "new2", Targets: endpoint.Targets{"target2"}},
	}
	changes := newCloudFlareChanges(cloudFlareCreate, endpoints, true)
	for i, change := range changes {
		assert.Equal(
			t,
			change.ResourceRecordSet[0].TTL,
			expect[i].TTL,
			expect[i].Name)
	}
}

func TestNewCloudFlareChangeNoProxied(t *testing.T) {
	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}}, false)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudFlareProxiedAnnotationTrue(t *testing.T) {
	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}, ProviderSpecific: endpoint.ProviderSpecific{
		endpoint.ProviderSpecificProperty{
			Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
			Value: "true",
		},
	}}, false)
	assert.True(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudFlareProxiedAnnotationFalse(t *testing.T) {
	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}, ProviderSpecific: endpoint.ProviderSpecific{
		endpoint.ProviderSpecificProperty{
			Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
			Value: "false",
		},
	}}, true)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudFlareProxiedAnnotationIllegalValue(t *testing.T) {
	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}, ProviderSpecific: endpoint.ProviderSpecific{
		endpoint.ProviderSpecificProperty{
			Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
			Value: "asdaslkjndaslkdjals",
		},
	}}, false)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudFlareChangeProxiable(t *testing.T) {
	var cloudFlareTypes = []struct {
		recordType string
		proxiable  bool
	}{
		{"A", true},
		{"CNAME", true},
		{"LOC", false},
		{"MX", false},
		{"NS", false},
		{"SPF", false},
		{"TXT", false},
		{"SRV", false},
	}

	for _, cloudFlareType := range cloudFlareTypes {
		change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: cloudFlareType.recordType, Targets: endpoint.Targets{"target"}}, true)

		if cloudFlareType.proxiable {
			assert.True(t, change.ResourceRecordSet[0].Proxied)
		} else {
			assert.False(t, change.ResourceRecordSet[0].Proxied)
		}
	}

	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "*.foo", RecordType: "A", Targets: endpoint.Targets{"target"}}, true)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestCloudFlareZones(t *testing.T) {
	provider := &CloudFlareProvider{
		Client:       &mockCloudFlareClient{},
		domainFilter: endpoint.NewDomainFilter([]string{"zalando.to."}),
		zoneIDFilter: NewZoneIDFilter([]string{""}),
	}

	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	validateCloudFlareZones(t, zones, []cloudflare.Zone{
		{Name: "ext-dns-test.zalando.to."},
	})
}

func TestRecords(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	assert.Equal(t, 1, len(records))
	provider.Client = &mockCloudFlareDNSRecordsFail{}
	_, err = provider.Records(ctx)
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockCloudFlareListZonesFail{}
	_, err = provider.Records(ctx)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestNewCloudFlareProvider(t *testing.T) {
	_ = os.Setenv("CF_API_TOKEN", "abc123def")
	_, err := NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}),
		NewZoneIDFilter([]string{""}),
		25,
		false,
		true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("CF_API_TOKEN")
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	_, err = NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}),
		NewZoneIDFilter([]string{""}),
		1,
		false,
		true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("CF_API_KEY")
	_ = os.Unsetenv("CF_API_EMAIL")
	_, err = NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}),
		NewZoneIDFilter([]string{""}),
		50,
		false,
		true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	changes.Create = []*endpoint.Endpoint{{DNSName: "new.ext-dns-test.zalando.to.", Targets: endpoint.Targets{"target"}}, {DNSName: "new.ext-dns-test.unrelated.to.", Targets: endpoint.Targets{"target"}}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Targets: endpoint.Targets{"target"}}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Targets: endpoint.Targets{"target-old"}}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Targets: endpoint.Targets{"target-new"}}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	// empty changes
	changes.Create = []*endpoint.Endpoint{}
	changes.Delete = []*endpoint.Endpoint{}
	changes.UpdateOld = []*endpoint.Endpoint{}
	changes.UpdateNew = []*endpoint.Endpoint{}

	err = provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestCloudFlareGetRecordID(t *testing.T) {
	p := &CloudFlareProvider{}
	records := []cloudflare.DNSRecord{
		{
			Name: "foo.com",
			Type: endpoint.RecordTypeCNAME,
			ID:   "1",
		},
		{
			Name: "bar.de",
			Type: endpoint.RecordTypeA,
			ID:   "2",
		},
	}

	assert.Len(t, p.getRecordIDs(records, cloudflare.DNSRecord{
		Name: "foo.com",
		Type: endpoint.RecordTypeA,
	}), 0)
	assert.Len(t, p.getRecordIDs(records, cloudflare.DNSRecord{
		Name: "bar.de",
		Type: endpoint.RecordTypeA,
	}), 1)
	assert.Equal(t, "2", p.getRecordIDs(records, cloudflare.DNSRecord{
		Name: "bar.de",
		Type: endpoint.RecordTypeA,
	})[0])
}

func validateCloudFlareZones(t *testing.T, zones []cloudflare.Zone, expected []cloudflare.Zone) {
	require.Len(t, zones, len(expected))

	for i, zone := range zones {
		assert.Equal(t, expected[i].Name, zone.Name)
	}
}

func TestGroupByNameAndType(t *testing.T) {
	testCases := []struct {
		Name              string
		Records           []cloudflare.DNSRecord
		ExpectedEndpoints []*endpoint.Endpoint
	}{
		{
			Name:              "empty",
			Records:           []cloudflare.DNSRecord{},
			ExpectedEndpoints: []*endpoint.Endpoint{},
		},
		{
			Name: "single record - single target",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "single record - multiple targets",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "multiple record - multiple targets",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
				{
					DNSName:    "bar.de",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "multiple record - mixed single/multiple targets",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
				{
					DNSName:    "bar.de",
					Targets:    endpoint.Targets{"10.10.10.1"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "unsupported record type",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    "NOT SUPPORTED",
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, groupByNameAndType(tc.Records), tc.ExpectedEndpoints)
	}
}
