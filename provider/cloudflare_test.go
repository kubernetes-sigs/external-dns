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
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockCloudflareClient struct{}

func (m *mockCloudflareClient) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareClient) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	if zoneID == "1234567890" {
		return []cloudflare.DNSRecord{
				{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to.", Type: endpoint.RecordTypeA, TTL: 120},
				{ID: "1231231233", Name: "foo.bar.com", TTL: 1}},
			nil
	}
	return nil, nil
}

func (m *mockCloudflareClient) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareClient) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareClient) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudflareClient) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudflareClient) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{ID: "1234567890", Name: "ext-dns-test.zalando.to."}, {ID: "1234567891", Name: "foo.com."}}, nil
}

func (m *mockCloudflareClient) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
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

type mockCloudflareUserDetailsFail struct{}

func (m *mockCloudflareUserDetailsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareUserDetailsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudflareUserDetailsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareUserDetailsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareUserDetailsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, fmt.Errorf("could not get ID from zone name")
}

func (m *mockCloudflareUserDetailsFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudflareUserDetailsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareUserDetailsFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{}, nil
}

type mockCloudflareCreateZoneFail struct{}

func (m *mockCloudflareCreateZoneFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareCreateZoneFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudflareCreateZoneFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareCreateZoneFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareCreateZoneFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudflareCreateZoneFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudflareCreateZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudflareDNSRecordsFail struct{}

func (m *mockCloudflareDNSRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareDNSRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, fmt.Errorf("can not get records from zone")
}
func (m *mockCloudflareDNSRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareDNSRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareDNSRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudflareDNSRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudflareDNSRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareDNSRecordsFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{
		Result: []cloudflare.Zone{
			{ID: "1234567890", Name: "ext-dns-test.zalando.to."},
			{ID: "1234567891", Name: "foo.com."}},
		ResultInfo: cloudflare.ResultInfo{
			TotalPages: 1,
		},
	}, nil
}

type mockCloudflareZoneIDByNameFail struct{}

func (m *mockCloudflareZoneIDByNameFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareZoneIDByNameFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudflareZoneIDByNameFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareZoneIDByNameFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareZoneIDByNameFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudflareZoneIDByNameFail) ZoneIDByName(zoneName string) (string, error) {
	return "", fmt.Errorf("no ID for zone found")
}

func (m *mockCloudflareZoneIDByNameFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudflareDeleteZoneFail struct{}

func (m *mockCloudflareDeleteZoneFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareDeleteZoneFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudflareDeleteZoneFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareDeleteZoneFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareDeleteZoneFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudflareDeleteZoneFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudflareDeleteZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudflareListZonesFail struct{}

func (m *mockCloudflareListZonesFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareListZonesFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudflareListZonesFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareListZonesFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareListZonesFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudflareListZonesFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudflareListZonesFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

func (m *mockCloudflareListZonesFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{}, fmt.Errorf("no zones available")
}

type mockCloudflareCreateRecordsFail struct{}

func (m *mockCloudflareCreateRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, fmt.Errorf("could not create record")
}

func (m *mockCloudflareCreateRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareCreateRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareCreateRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareCreateRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudflareCreateRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudflareCreateRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

func (m *mockCloudflareCreateRecordsFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{}, nil
}

type mockCloudflareDeleteRecordsFail struct{}

func (m *mockCloudflareDeleteRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareDeleteRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareDeleteRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudflareDeleteRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return fmt.Errorf("could not delete record")
}

func (m *mockCloudflareDeleteRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudflareDeleteRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudflareDeleteRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareDeleteRecordsFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{}, nil
}

type mockCloudflareUpdateRecordsFail struct{}

func (m *mockCloudflareUpdateRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudflareUpdateRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareUpdateRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return fmt.Errorf("could not update record")
}

func (m *mockCloudflareUpdateRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudflareUpdateRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudflareUpdateRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudflareUpdateRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudflareUpdateRecordsFail) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return cloudflare.ZonesResponse{}, nil
}

func TestNewCloudflareChanges(t *testing.T) {
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
	changes := newCloudflareChanges(cloudflareCreate, endpoints, true)
	for i, change := range changes {
		assert.Equal(
			t,
			change.ResourceRecordSet[0].TTL,
			expect[i].TTL,
			expect[i].Name)
	}
}

func TestNewCloudflareChangeNoProxied(t *testing.T) {
	change := newCloudflareChange(cloudflareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}}, false)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudflareProxiedAnnotationTrue(t *testing.T) {
	change := newCloudflareChange(cloudflareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}, ProviderSpecific: endpoint.ProviderSpecific{
		endpoint.ProviderSpecificProperty{
			Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
			Value: "true",
		},
	}}, false)
	assert.True(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudflareProxiedAnnotationFalse(t *testing.T) {
	change := newCloudflareChange(cloudflareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}, ProviderSpecific: endpoint.ProviderSpecific{
		endpoint.ProviderSpecificProperty{
			Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
			Value: "false",
		},
	}}, true)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudflareProxiedAnnotationIllegalValue(t *testing.T) {
	change := newCloudflareChange(cloudflareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}, ProviderSpecific: endpoint.ProviderSpecific{
		endpoint.ProviderSpecificProperty{
			Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
			Value: "asdaslkjndaslkdjals",
		},
	}}, false)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestNewCloudflareChangeProxiable(t *testing.T) {
	var cloudflareTypes = []struct {
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

	for _, cloudflareType := range cloudflareTypes {
		change := newCloudflareChange(cloudflareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: cloudflareType.recordType, Targets: endpoint.Targets{"target"}}, true)

		if cloudflareType.proxiable {
			assert.True(t, change.ResourceRecordSet[0].Proxied)
		} else {
			assert.False(t, change.ResourceRecordSet[0].Proxied)
		}
	}

	change := newCloudflareChange(cloudflareCreate, &endpoint.Endpoint{DNSName: "*.foo", RecordType: "A", Targets: endpoint.Targets{"target"}}, true)
	assert.False(t, change.ResourceRecordSet[0].Proxied)
}

func TestCloudflareZones(t *testing.T) {
	provider := &CloudflareProvider{
		Client:       &mockCloudflareClient{},
		domainFilter: NewDomainFilter([]string{"zalando.to."}),
		zoneIDFilter: NewZoneIDFilter([]string{""}),
	}

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}

	validateCloudflareZones(t, zones, []cloudflare.Zone{
		{Name: "ext-dns-test.zalando.to."},
	})
}

func TestRecords(t *testing.T) {
	provider := &CloudflareProvider{
		Client: &mockCloudflareClient{},
	}
	records, err := provider.Records()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	assert.Equal(t, 1, len(records))
	provider.Client = &mockCloudflareDNSRecordsFail{}
	_, err = provider.Records()
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockCloudflareListZonesFail{}
	_, err = provider.Records()
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestNewCloudflareProvider(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	_, err := NewCloudflareProvider(
		NewDomainFilter([]string{"ext-dns-test.zalando.to."}),
		NewZoneIDFilter([]string{""}),
		1,
		false,
		true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("CF_API_KEY")
	_ = os.Unsetenv("CF_API_EMAIL")
	_, err = NewCloudflareProvider(
		NewDomainFilter([]string{"ext-dns-test.zalando.to."}),
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
	provider := &CloudflareProvider{
		Client: &mockCloudflareClient{},
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

func TestCloudflareGetRecordID(t *testing.T) {
	p := &CloudflareProvider{}
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

func validateCloudflareZones(t *testing.T, zones []cloudflare.Zone, expected []cloudflare.Zone) {
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
					TTL:     defaultCloudflareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudflareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudflareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudflareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudflareRecordTTL,
				},
				{
					Name:    "bar.de",
					Type:    "NOT SUPPORTED",
					Content: "10.10.10.1",
					TTL:     defaultCloudflareRecordTTL,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudflareRecordTTL),
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
		assert.Equal(t, groupByNameAndType(tc.Records), tc.ExpectedEndpoints)
	}
}
