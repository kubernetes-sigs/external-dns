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

	"github.com/cloudflare/cloudflare-go"
)

//type mockManagedZonesClient struct{}
//
//func (m *mockManagedZonesClient) Create(project string, managedzone *dns.ManagedZone) managedZonesCreateCallInterface {
//	return &mockManagedZonesCreateCall{}
//}
//
//func (m *mockManagedZonesClient) Delete(project string, managedZone string) managedZonesDeleteCallInterface {
//	return &mockManagedZonesDeleteCall{}
//}
//
//func (m *mockManagedZonesClient) List(project string) managedZonesListCallInterface {
//	return &mockManagedZonesListCall{}
//}

type mockCloudFlareClient struct{}

func (m *mockCloudFlareClient) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareClient) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{cloudflare.DNSRecord{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
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

func (m *mockCloudFlareClient) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareClient) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareClient) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareClient) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{Name: "ext-dns-test.zalando.to."}, nil
}

func (m *mockCloudFlareClient) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{ID: "1234567890"}, nil
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

func (m *mockCloudFlareUserDetailsFail) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareUserDetailsFail) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareUserDetailsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareUserDetailsFail) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{Name: "ext-dns-test.zalando.to."}, nil
}

func (m *mockCloudFlareUserDetailsFail) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{ID: "1234567890"}, nil
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

func (m *mockCloudFlareCreateZoneFail) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareCreateZoneFail) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareCreateZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareCreateZoneFail) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{}, fmt.Errorf("could not create new zone")
}

func (m *mockCloudFlareCreateZoneFail) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{ID: "1234567890"}, nil
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

func (m *mockCloudFlareDNSRecordsFail) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDNSRecordsFail) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDNSRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDNSRecordsFail) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{}, nil
}

func (m *mockCloudFlareDNSRecordsFail) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{ID: "1234567890"}, nil
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

func (m *mockCloudFlareZoneIDByNameFail) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{ID: "1234567890"}, nil
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

func (m *mockCloudFlareDeleteZoneFail) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDeleteZoneFail) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDeleteZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDeleteZoneFail) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{}, nil
}

func (m *mockCloudFlareDeleteZoneFail) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{}, fmt.Errorf("could not delete zone")
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

func (m *mockCloudFlareListZonesFail) Zones() ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareListZonesFail) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareListZonesFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

func (m *mockCloudFlareListZonesFail) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return cloudflare.Zone{}, nil
}

func (m *mockCloudFlareListZonesFail) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return cloudflare.ZoneID{}, fmt.Errorf("could not delete zone")
}
func TestCloudFlareCreateZone(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	_, err := provider.CreateZone("ext-dns-test.zalando.to.")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudFlareDeleteZone(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	_, err := provider.DeleteZone("ext-dns-test.zalando.to.")
	if err != nil {
		t.Fatal(err)
	}
	provider.Client = &mockCloudFlareZoneIDByNameFail{}
	_, err = provider.DeleteZone("ext-dns-test.zalando.to.")
	if err == nil {
		t.Errorf("exptected error")
	}
	provider.Client = &mockCloudFlareDeleteZoneFail{}
	_, err = provider.DeleteZone("ext-dns-test.zalando.to.")
	if err == nil {
		t.Errorf("exptected error")
	}
}

func TestCloudFlareCreateZoneFail(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareUserDetailsFail{},
	}
	_, err := provider.CreateZone("ext-dns-test.zalando.to.")
	if err == nil {
		t.Errorf("exptected error")
	}
	provider.Client = &mockCloudFlareCreateZoneFail{}
	_, err = provider.CreateZone("ext-dns-test.zalando.to.")
	if err == nil {
		t.Errorf("exptected error")
	}
}

func TestCloudFlareGetRecordID(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareDNSRecordsFail{},
	}
	zoneID := "12345656790"
	record := cloudflare.DNSRecord{Name: "foobar.ext-dns-test.zalando.to"}
	_, err := provider.getRecordID(zoneID, record)
	if err == nil {
		t.Errorf("exptected error")
	}
	provider.Client = &mockCloudFlareClient{}
	_, err = provider.getRecordID(zoneID, record)
	if err == nil {
		t.Errorf("exptected error")
	}
	record = cloudflare.DNSRecord{Name: "foobar.ext-dns-test.zalando.to."}
	_, err = provider.getRecordID(zoneID, record)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestZones(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	_, err := provider.Zones()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockCloudFlareListZonesFail{}
	_, err = provider.Zones()
	if err != nil {
		//https://github.com/cloudflare/cloudflare-go/blob/ea9272e4235ff7a9aa37e2c7c7a8debe22b3d696/zone.go#L270
		t.Errorf("should not fail, %s", err)
	}

}

func TestZone(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	_, err := provider.Zone("ext-dns-test.zalando.to.")
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockCloudFlareZoneIDByNameFail{}
	_, err = provider.Zone("ext-dns-test.zalando.to.")
	if err == nil {
		t.Errorf("exptected error")
	}
	provider.Client = &mockCloudFlareListZonesFail{}
	_, err = provider.Zone("ext-dns-test.zalando.to.")
	if err != nil {
		//https://github.com/cloudflare/cloudflare-go/blob/ea9272e4235ff7a9aa37e2c7c7a8debe22b3d696/zone.go#L270
		t.Errorf("should not fail, %s", err)
	}
}

func TestCreateRecords(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	zone := "ext-dns-test.zalando.to"
	endpoints := []*endpoint.Endpoint{
		&endpoint.Endpoint{DNSName: "new", Target: "target"},
	}
	err := provider.CreateRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.DryRun = true
	err = provider.CreateRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockCloudFlareZoneIDByNameFail{}
	err = provider.CreateRecords(zone, endpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}

	endpoints = []*endpoint.Endpoint{}
	err = provider.CreateRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestDeleteRecords(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	zone := "ext-dns-test.zalando.to"
	endpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target"}}
	err := provider.DeleteRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.DryRun = true
	err = provider.DeleteRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	endpoints = []*endpoint.Endpoint{}
	err = provider.DeleteRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestUpdateRecords(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	zone := "ext-dns-test.zalando.to"
	oldCNameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "old-target"}}
	newCNameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "new-target"}}

	oldANameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "8.8.8.8"}}
	newANameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "7.7.7.7"}}
	err := provider.UpdateRecords(zone, newCNameEndpoints, oldCNameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	err = provider.UpdateRecords(zone, newANameEndpoints, oldANameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.DryRun = true
	err = provider.UpdateRecords(zone, newCNameEndpoints, oldCNameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	err = provider.UpdateRecords(zone, newANameEndpoints, oldANameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestNewCloudFlareChanges(t *testing.T) {
	action := CloudFlareCreate
	endpoints := []*endpoint.Endpoint{{DNSName: "new", Target: "target"}}
	_ = newCloudFlareChanges(action, endpoints)
}

func TestRecords(t *testing.T) {
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	zone := "ext-dns-test.zalando.to"
	_, err := provider.Records(zone)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockCloudFlareDNSRecordsFail{}
	_, err = provider.Records(zone)
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockCloudFlareZoneIDByNameFail{}
	_, err = provider.Records(zone)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestNewCloudFlareProvider(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	_, err := NewCloudFlareProvider(true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	provider := &CloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	zone := "ext-dns-test.zalando.to"
	changes.Create = []*endpoint.Endpoint{{DNSName: "new.ext-dns-test.zalando.to.", Target: "target"}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target"}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target-old"}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target-new"}}
	err := provider.ApplyChanges(zone, changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}
