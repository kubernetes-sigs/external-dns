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
	"fmt"
	"os"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"

	cloudflare "github.com/cloudflare/cloudflare-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

type mockCloudFlareUserDetailsFail struct{}

func (m *mockCloudFlareUserDetailsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareUserDetailsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareUserDetailsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareUserDetailsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareUserDetailsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, fmt.Errorf("could not get ID from zone name")
}

func (m *mockCloudFlareUserDetailsFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudFlareUserDetailsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareCreateZoneFail struct{}

func (m *mockCloudFlareCreateZoneFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareCreateZoneFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareCreateZoneFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareCreateZoneFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareCreateZoneFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareCreateZoneFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudFlareCreateZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
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

type mockCloudFlareZoneIDByNameFail struct{}

func (m *mockCloudFlareZoneIDByNameFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareZoneIDByNameFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareZoneIDByNameFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareZoneIDByNameFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) ZoneIDByName(zoneName string) (string, error) {
	return "", fmt.Errorf("no ID for zone found")
}

func (m *mockCloudFlareZoneIDByNameFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareDeleteZoneFail struct{}

func (m *mockCloudFlareDeleteZoneFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareDeleteZoneFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareDeleteZoneFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareDeleteZoneFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareDeleteZoneFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudFlareDeleteZoneFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareDeleteZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
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

type mockCloudFlareCreateRecordsFail struct{}

func (m *mockCloudFlareCreateRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, fmt.Errorf("could not create record")
}

func (m *mockCloudFlareCreateRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareCreateRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareCreateRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareCreateRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareCreateRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareCreateRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

type mockCloudFlareDeleteRecordsFail struct{}

func (m *mockCloudFlareDeleteRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareDeleteRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDeleteRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareDeleteRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return fmt.Errorf("could not delete record")
}

func (m *mockCloudFlareDeleteRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareDeleteRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareDeleteRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareUpdateRecordsFail struct{}

func (m *mockCloudFlareUpdateRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareUpdateRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareUpdateRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return fmt.Errorf("could not update record")
}

func (m *mockCloudFlareUpdateRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareUpdateRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareUpdateRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareUpdateRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
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
			change.ResourceRecordSet.TTL,
			expect[i].TTL,
			expect[i].Name)
	}
}

func TestNewCloudFlareChangeNoProxied(t *testing.T) {
	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "new", RecordType: "A", Targets: endpoint.Targets{"target"}}, false)
	assert.False(t, change.ResourceRecordSet.Proxied)
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
			assert.True(t, change.ResourceRecordSet.Proxied)
		} else {
			assert.False(t, change.ResourceRecordSet.Proxied)
		}
	}

	change := newCloudFlareChange(cloudFlareCreate, &endpoint.Endpoint{DNSName: "*.foo", RecordType: "A", Targets: endpoint.Targets{"target"}}, true)
	assert.False(t, change.ResourceRecordSet.Proxied)
}

func TestCloudFlareZones(t *testing.T) {
	provider := &CloudFlareProvider{
		Client:       &mockCloudFlareClient{},
		domainFilter: NewDomainFilter([]string{"zalando.to."}),
		zoneIDFilter: NewZoneIDFilter([]string{""}),
	}

	zones, err := provider.Zones()
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
	records, err := provider.Records()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	assert.Equal(t, 1, len(records))
	provider.Client = &mockCloudFlareDNSRecordsFail{}
	_, err = provider.Records()
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockCloudFlareListZonesFail{}
	_, err = provider.Records()
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestNewCloudFlareProvider(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	_, err := NewCloudFlareProvider(NewDomainFilter([]string{"ext-dns-test.zalando.to."}), NewZoneIDFilter([]string{""}), false, true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("CF_API_KEY")
	_ = os.Unsetenv("CF_API_EMAIL")
	_, err = NewCloudFlareProvider(NewDomainFilter([]string{"ext-dns-test.zalando.to."}), NewZoneIDFilter([]string{""}), false, true)
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
	err := provider.ApplyChanges(changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	// empty changes
	changes.Create = []*endpoint.Endpoint{}
	changes.Delete = []*endpoint.Endpoint{}
	changes.UpdateOld = []*endpoint.Endpoint{}
	changes.UpdateNew = []*endpoint.Endpoint{}

	err = provider.ApplyChanges(changes)
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

	assert.Equal(t, "", p.getRecordID(records, cloudflare.DNSRecord{
		Name: "foo.com",
		Type: endpoint.RecordTypeA,
	}))
	assert.Equal(t, "2", p.getRecordID(records, cloudflare.DNSRecord{
		Name: "bar.de",
		Type: endpoint.RecordTypeA,
	}))
}

func validateCloudFlareZones(t *testing.T, zones []cloudflare.Zone, expected []cloudflare.Zone) {
	require.Len(t, zones, len(expected))

	for i, zone := range zones {
		assert.Equal(t, expected[i].Name, zone.Name)
	}
}
