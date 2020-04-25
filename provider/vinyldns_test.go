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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinyldns/go-vinyldns/vinyldns"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockVinyldnsZoneInterface struct {
	mock.Mock
}

var mockVinylDNSProvider vinyldnsProvider

var vinylDNSZones []vinyldns.Zone
var vinylDNSRecords []vinyldns.RecordSet
var vinylDNSRecordSetUpdateResponse *vinyldns.RecordSetUpdateResponse

func TestVinylDNSServices(t *testing.T) {
	firstZone := vinyldns.Zone{
		ID:   "0",
		Name: "example.com.",
	}
	secondZone := vinyldns.Zone{
		ID:   "1",
		Name: "example-beta.com.",
	}
	vinylDNSZones = []vinyldns.Zone{firstZone, secondZone}

	firstRecord := vinyldns.RecordSet{
		ZoneID: "0",
		Name:   "example.com.",
		TTL:    300,
		Type:   "CNAME",
		Records: []vinyldns.Record{
			{
				CName: "vinyldns.com",
			},
		},
	}
	vinylDNSRecords = []vinyldns.RecordSet{firstRecord}

	vinylDNSRecordSetUpdateResponse = &vinyldns.RecordSetUpdateResponse{
		Zone:      firstZone,
		RecordSet: firstRecord,
		ChangeID:  "123",
		Status:    "Active",
	}

	mockVinylDNS := &mockVinyldnsZoneInterface{}
	mockVinylDNS.On("Zones").Return(vinylDNSZones, nil)
	mockVinylDNS.On("RecordSets", "0").Return(vinylDNSRecords, nil)
	mockVinylDNS.On("RecordSets", "1").Return(nil, nil)
	mockVinylDNS.On("RecordSets", "2").Return(nil, fmt.Errorf("Record not found"))
	mockVinylDNS.On("RecordSetCreate", &firstRecord).Return(vinylDNSRecordSetUpdateResponse, nil)
	mockVinylDNS.On("RecordSetUpdate", &firstRecord).Return(vinylDNSRecordSetUpdateResponse, nil)
	mockVinylDNS.On("RecordSetDelete", "0", "").Return(nil, nil)

	mockVinylDNSProvider = vinyldnsProvider{client: mockVinylDNS}

	// Run tests on mock services
	t.Run("Records", testVinylDNSProviderRecords)
	t.Run("ApplyChanges", testVinylDNSProviderApplyChanges)
	t.Run("SuitableZone", testVinylDNSSuitableZone)
	t.Run("GetRecordID", testVinylDNSFindRecordSetID)
}

func testVinylDNSProviderRecords(t *testing.T) {
	ctx := context.Background()

	mockVinylDNSProvider.domainFilter = endpoint.NewDomainFilter([]string{"example.com"})
	result, err := mockVinylDNSProvider.Records(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(vinylDNSRecords), len(result))

	mockVinylDNSProvider.zoneFilter = NewZoneIDFilter([]string{"0"})
	result, err = mockVinylDNSProvider.Records(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(vinylDNSRecords), len(result))

	mockVinylDNSProvider.zoneFilter = NewZoneIDFilter([]string{"1"})
	result, err = mockVinylDNSProvider.Records(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func testVinylDNSProviderApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.com", Targets: endpoint.Targets{"vinyldns.com"}, RecordType: endpoint.RecordTypeCNAME},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "example.com", Targets: endpoint.Targets{"vinyldns.com"}, RecordType: endpoint.RecordTypeCNAME},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "example.com", Targets: endpoint.Targets{"vinyldns.com"}, RecordType: endpoint.RecordTypeCNAME}}

	mockVinylDNSProvider.zoneFilter = NewZoneIDFilter([]string{"1"})
	err := mockVinylDNSProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("Failed to apply changes: %v", err)
	}
}

func testVinylDNSSuitableZone(t *testing.T) {
	mockVinylDNSProvider.zoneFilter = NewZoneIDFilter([]string{"0"})

	zone := vinyldnsSuitableZone("example.com", vinylDNSZones)
	assert.Equal(t, zone.Name, "example.com.")
}

func TestNewVinylDNSProvider(t *testing.T) {
	os.Setenv("VINYLDNS_ACCESS_KEY", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	_, err := NewVinylDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), NewZoneIDFilter([]string{"0"}), true)
	assert.Nil(t, err)

	os.Unsetenv("VINYLDNS_ACCESS_KEY")
	_, err = NewVinylDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), NewZoneIDFilter([]string{"0"}), true)
	assert.NotNil(t, err)
	if err == nil {
		t.Errorf("Expected to fail new provider on empty token")
	}
}

func testVinylDNSFindRecordSetID(t *testing.T) {
	mockVinylDNSProvider.zoneFilter = NewZoneIDFilter([]string{"0"})
	result, err := mockVinylDNSProvider.findRecordSetID("0", "example.com.")
	assert.Nil(t, err)
	assert.Equal(t, "", result)

	_, err = mockVinylDNSProvider.findRecordSetID("2", "example-beta")
	assert.NotNil(t, err)
}

func (m *mockVinyldnsZoneInterface) Zones() ([]vinyldns.Zone, error) {
	args := m.Called()
	var r0 []vinyldns.Zone

	if args.Get(0) != nil {
		r0 = args.Get(0).([]vinyldns.Zone)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSet(zoneID, recordSet string) (vinyldns.RecordSet, error) {
	args := m.Called(zoneID, recordSet)
	var r0 vinyldns.RecordSet

	if args.Get(0) != nil {
		r0 = args.Get(0).(vinyldns.RecordSet)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSets(id string) ([]vinyldns.RecordSet, error) {
	args := m.Called(id)
	var r0 []vinyldns.RecordSet

	if args.Get(0) != nil {
		r0 = args.Get(0).([]vinyldns.RecordSet)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSetCreate(rs *vinyldns.RecordSet) (*vinyldns.RecordSetUpdateResponse, error) {
	args := m.Called(rs)
	var r0 *vinyldns.RecordSetUpdateResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*vinyldns.RecordSetUpdateResponse)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSetUpdate(rs *vinyldns.RecordSet) (*vinyldns.RecordSetUpdateResponse, error) {
	args := m.Called(rs)
	var r0 *vinyldns.RecordSetUpdateResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*vinyldns.RecordSetUpdateResponse)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSetDelete(zoneID, recordSetID string) (*vinyldns.RecordSetUpdateResponse, error) {
	args := m.Called(zoneID, recordSetID)
	var r0 *vinyldns.RecordSetUpdateResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*vinyldns.RecordSetUpdateResponse)
	}

	return r0, args.Error(1)
}
