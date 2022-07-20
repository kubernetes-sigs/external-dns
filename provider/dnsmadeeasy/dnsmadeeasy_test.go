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

package dnsmadeeasy

import (
	"context"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var mockProvider dmeProvider
var dnsMadeEasyListRecordsResponse dmeListRecordResponse
var dnsMadeEasyListZonesResponse dmeListZoneResponse

func TestDnsMadeEasyServices(t *testing.T) {

	firstZone := Zone{
		Id: "1119443",
		DomainAttribute: models.DomainAttribute{
			Name: "example.com",
		},
	}

	secondZone := Zone{
		Id: "1119433",
		DomainAttribute: models.DomainAttribute{
			Name: "example-beta.com",
		},
	}

	dnsMadeEasyListZonesResponse = dmeListZoneResponse{
		Zones: []Zone{firstZone, secondZone},
		dmePaginatedResponse: dmePaginatedResponse{
			Page:         1,
			TotalPages:   1,
			TotalRecords: 2,
		},
	}

	firstRecord := Record{
		Id:   "2",
		Zone: &firstZone,
		ManagedDNSRecordActions: models.ManagedDNSRecordActions{
			Name:     "example",
			Value:    "target",
			Type:     "CNAME",
			Ttl:      "1800",
			Priority: "0",
		},
	}
	secondRecord := Record{
		Id:   "1",
		Zone: &firstZone,
		ManagedDNSRecordActions: models.ManagedDNSRecordActions{
			Name:     "example-beta",
			Value:    "127.0.0.1",
			Type:     "A",
			Ttl:      "1800",
			Priority: "0",
		},
	}
	thirdRecord := Record{
		Id:   "3",
		Zone: &firstZone,
		ManagedDNSRecordActions: models.ManagedDNSRecordActions{
			Name:     "custom-ttl",
			Value:    "target",
			Type:     "CNAME",
			Ttl:      "60",
			Priority: "0",
		},
	}
	fourthRecord := Record{
		Id:   "4",
		Zone: &firstZone,
		ManagedDNSRecordActions: models.ManagedDNSRecordActions{
			Name:     "", // Apex domain A record
			Value:    "127.0.0.1",
			Type:     "A",
			Ttl:      "1800",
			Priority: "0",
		},
	}
	fifthRecord := Record{
		Id:   "5",
		Zone: &firstZone,
		ManagedDNSRecordActions: models.ManagedDNSRecordActions{
			Name:     "cname-target",
			Value:    "target.anotherdomain.com.",
			Type:     "CNAME",
			Ttl:      "1800",
			Priority: "0",
		},
	}

	records := []Record{firstRecord, secondRecord, thirdRecord, fourthRecord, fifthRecord}

	dnsMadeEasyListRecordsResponse = dmeListRecordResponse{
		Records: records,
		dmePaginatedResponse: dmePaginatedResponse{
			Page:         1,
			TotalPages:   1,
			TotalRecords: len(records),
		},
	}

	// Setup mock services
	// Note: AnythingOfType doesn't work with interfaces https://github.com/stretchr/testify/issues/519
	mockDNS := &mockDmeService{}
	mockDNS.On("ListZones", mock.AnythingOfType("*context.emptyCtx"), 1).Return(&dnsMadeEasyListZonesResponse, nil)
	mockDNS.On("ListRecords", mock.AnythingOfType("*context.emptyCtx"), "1119443", 1).Return(&dnsMadeEasyListRecordsResponse, nil)
	mockDNS.On("ListRecords", mock.AnythingOfType("*context.emptyCtx"), "1119433", 1).Return(&dmeListRecordResponse{Records: []Record{}}, nil)

	for _, record := range records {
		recordName := record.Name
		simpleRecord := Record{
			ManagedDNSRecordActions: models.ManagedDNSRecordActions{
				Name:  recordName,
				Type:  record.Type,
				Value: record.Value,
				Ttl:   record.Ttl,
			},
		}

		dmeRecordResponse := dmeListRecordResponse{
			dmePaginatedResponse: dmePaginatedResponse{
				Page:         1,
				TotalPages:   1,
				TotalRecords: 1,
			},
			Records: []Record{record},
		}

		mockDNS.On("ListRecords", mock.AnythingOfType("*context.emptyCtx"), record.Zone.Id, 1).Return(&dmeRecordResponse, nil)
		mockDNS.On("DeleteRecord", mock.AnythingOfType("*context.emptyCtx"), record.Zone.Id, record.Id).Return(nil)
		mockDNS.On("UpdateRecord", mock.AnythingOfType("*context.emptyCtx"), record.Zone.Id, record.Id, simpleRecord).Return(Record{}, nil)
		recordWithoutId := record
		recordWithoutId.Id = ""
		mockDNS.On("CreateRecord", mock.AnythingOfType("*context.emptyCtx"), record.Zone.Id, simpleRecord).Return(&Record{}, nil)
	}

	mockProvider = dmeProvider{service: mockDNS}

	// Run tests on mock services
	t.Run("Zones", testDnsMadeEasyProviderZones)
	t.Run("Records", testDnsMadeEasyProviderRecords)
	t.Run("ApplyChanges", testDnsMadeEasyProviderApplyChanges)
	t.Run("ApplyChanges/SkipUnknownZone", testDnsMadeEasyProviderApplyChangesSkipsUnknown)
	t.Run("SuitableZone", testDnsMadeEasySuitableZone)
	t.Run("GetRecordID", testDnsMadeEasyGetRecordID)
}

func testDnsMadeEasyProviderZones(t *testing.T) {
	ctx := context.Background()
	result, err := mockProvider.Zones(ctx)
	assert.Nil(t, err)
	validateDnsMadeEasyZones(t, result, dnsMadeEasyListZonesResponse.Zones)
}

func testDnsMadeEasyProviderRecords(t *testing.T) {
	ctx := context.Background()
	result, err := mockProvider.Records(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(dnsMadeEasyListRecordsResponse.Records), len(result))
}

func testDnsMadeEasyProviderApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.example.com", Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "custom-ttl.example.com", RecordTTL: 60, Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "cname-target.example.com", Targets: endpoint.Targets{"target.anotherdomain.com"}, RecordType: endpoint.RecordTypeCNAME},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "example-beta.example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "example.example.com", Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
	}
	ctx := context.Background()

	err := mockProvider.ApplyChanges(ctx, changes)
	if err != nil {
		t.Errorf("Failed to apply changes: %v", err)
	}
}

func testDnsMadeEasyProviderApplyChangesSkipsUnknown(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.not-included.com", Targets: endpoint.Targets{"dasd"}, RecordType: endpoint.RecordTypeCNAME},
	}

	err := mockProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("Failed to ignore unknown zones: %v", err)
	}
}

func testDnsMadeEasySuitableZone(t *testing.T) {
	ctx := context.Background()
	zones, err := mockProvider.Zones(ctx)
	assert.Nil(t, err)

	_, zone := dnsmadeeasySuitableZone("example-beta.example.com", zones)
	assert.Equal(t, zone.Name, "example.com")
}

func TestNewDnsMadeEasyProvider(t *testing.T) {
	os.Unsetenv("dme_apikey")
	os.Unsetenv("dme_secretkey")
	_, err := NewDmeProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true)
	if err == nil {
		t.Errorf("Expected to fail new provider on bad token")
	}

}

func testDnsMadeEasyGetRecordID(t *testing.T) {
	var result string
	var err error

	result, err = mockProvider.GetRecordID(context.Background(), "1119443", "example")
	assert.Nil(t, err)
	assert.Equal(t, "2", result)

	result, err = mockProvider.GetRecordID(context.Background(), "1119443", "example-beta")
	assert.Nil(t, err)
	assert.Equal(t, "1", result)
}

func validateDnsMadeEasyZones(t *testing.T, zones map[string]Zone, expected []Zone) {
	require.Len(t, zones, len(expected))

	for _, e := range expected {
		assert.Equal(t, zones[e.Id].Name, e.Name)
	}
}

type mockDmeService struct {
	mock.Mock
}

func (_m *mockDmeService) CreateRecord(ctx context.Context, zoneID string, recordAttributes Record) (*Record, error) {
	args := _m.Called(ctx, zoneID, recordAttributes)
	var r0 *Record

	if args.Get(0) != nil {
		r0 = args.Get(0).(*Record)
	}

	return r0, args.Error(1)
}

func (_m *mockDmeService) DeleteRecord(ctx context.Context, zoneID string, recordID string) error {
	args := _m.Called(ctx, zoneID, recordID)
	return args.Error(0)
}

func (_m *mockDmeService) ListRecords(ctx context.Context, zoneID string, page int) (response *dmeListRecordResponse, _ error) {
	args := _m.Called(ctx, zoneID, page)
	var r0 *dmeListRecordResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dmeListRecordResponse)
	}

	return r0, args.Error(1)
}

func (_m *mockDmeService) ListZones(ctx context.Context, page int) (response *dmeListZoneResponse, _ error) {
	args := _m.Called(ctx, page)
	var r0 *dmeListZoneResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dmeListZoneResponse)
	}

	return r0, args.Error(1)
}

func (_m *mockDmeService) GetZoneByName(ctx context.Context, zoneName string) (response *Zone, _ error) {
	args := _m.Called(ctx, zoneName)
	var r0 *Zone

	if args.Get(0) != nil {
		r0 = args.Get(0).(*Zone)
	}

	return r0, args.Error(1)
}

func (_m *mockDmeService) UpdateRecord(ctx context.Context, zoneID string, recordID string, recordAttributes Record) (*Record, error) {
	args := _m.Called(ctx, zoneID, recordID, recordAttributes)
	var r0 Record

	if args.Get(0) != nil {
		r0 = args.Get(0).(Record)
	}

	return &r0, args.Error(1)
}
