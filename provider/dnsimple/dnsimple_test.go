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

package dnsimple

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/dnsimple/dnsimple-go/dnsimple"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var (
	mockProvider                dnsimpleProvider
	dnsimpleListRecordsResponse dnsimple.ZoneRecordsResponse
	dnsimpleListZonesResponse   dnsimple.ZonesResponse
)

func TestDnsimpleServices(t *testing.T) {
	// Setup example responses
	firstZone := dnsimple.Zone{
		ID:        1,
		AccountID: 12345,
		Name:      "example.com",
	}
	secondZone := dnsimple.Zone{
		ID:        2,
		AccountID: 54321,
		Name:      "example-beta.com",
	}
	zones := []dnsimple.Zone{firstZone, secondZone}
	dnsimpleListZonesResponse = dnsimple.ZonesResponse{
		Response: dnsimple.Response{Pagination: &dnsimple.Pagination{}},
		Data:     zones,
	}
	firstRecord := dnsimple.ZoneRecord{
		ID:       2,
		ZoneID:   "example.com",
		ParentID: 0,
		Name:     "example",
		Content:  "target",
		TTL:      3600,
		Priority: 0,
		Type:     "CNAME",
	}
	secondRecord := dnsimple.ZoneRecord{
		ID:       1,
		ZoneID:   "example.com",
		ParentID: 0,
		Name:     "example-beta",
		Content:  "127.0.0.1",
		TTL:      3600,
		Priority: 0,
		Type:     "A",
	}
	thirdRecord := dnsimple.ZoneRecord{
		ID:       3,
		ZoneID:   "example.com",
		ParentID: 0,
		Name:     "custom-ttl",
		Content:  "target",
		TTL:      60,
		Priority: 0,
		Type:     "CNAME",
	}
	fourthRecord := dnsimple.ZoneRecord{
		ID:       4,
		ZoneID:   "example.com",
		ParentID: 0,
		Name:     "", // Apex domain A record
		Content:  "127.0.0.1",
		TTL:      3600,
		Priority: 0,
		Type:     "A",
	}

	records := []dnsimple.ZoneRecord{firstRecord, secondRecord, thirdRecord, fourthRecord}
	dnsimpleListRecordsResponse = dnsimple.ZoneRecordsResponse{
		Response: dnsimple.Response{Pagination: &dnsimple.Pagination{}},
		Data:     records,
	}

	// Setup mock services
	// Note: AnythingOfType doesn't work with interfaces https://github.com/stretchr/testify/issues/519
	mockDNS := &mockDnsimpleZoneServiceInterface{}
	mockDNS.On("ListZones", mock.AnythingOfType("*context.emptyCtx"), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: dnsimple.Int(1)}}).Return(&dnsimpleListZonesResponse, nil)
	mockDNS.On("ListZones", mock.AnythingOfType("*context.emptyCtx"), "2", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: dnsimple.Int(1)}}).Return(nil, fmt.Errorf("Account ID not found"))
	mockDNS.On("ListRecords", mock.AnythingOfType("*context.emptyCtx"), "1", "example.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: dnsimple.Int(1)}}).Return(&dnsimpleListRecordsResponse, nil)
	mockDNS.On("ListRecords", mock.AnythingOfType("*context.emptyCtx"), "1", "example-beta.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: dnsimple.Int(1)}}).Return(&dnsimple.ZoneRecordsResponse{Response: dnsimple.Response{Pagination: &dnsimple.Pagination{}}}, nil)

	for _, record := range records {
		recordName := record.Name
		simpleRecord := dnsimple.ZoneRecordAttributes{
			Name:    &recordName,
			Type:    record.Type,
			Content: record.Content,
			TTL:     record.TTL,
		}

		dnsimpleRecordResponse := dnsimple.ZoneRecordsResponse{
			Response: dnsimple.Response{Pagination: &dnsimple.Pagination{}},
			Data:     []dnsimple.ZoneRecord{record},
		}

		mockDNS.On("ListRecords", mock.AnythingOfType("*context.emptyCtx"), "1", record.ZoneID, &dnsimple.ZoneRecordListOptions{Name: &recordName, ListOptions: dnsimple.ListOptions{Page: dnsimple.Int(1)}}).Return(&dnsimpleRecordResponse, nil)
		mockDNS.On("CreateRecord", mock.AnythingOfType("*context.emptyCtx"), "1", record.ZoneID, simpleRecord).Return(&dnsimple.ZoneRecordResponse{}, nil)
		mockDNS.On("DeleteRecord", mock.AnythingOfType("*context.emptyCtx"), "1", record.ZoneID, record.ID).Return(&dnsimple.ZoneRecordResponse{}, nil)
		mockDNS.On("UpdateRecord", mock.AnythingOfType("*context.emptyCtx"), "1", record.ZoneID, record.ID, simpleRecord).Return(&dnsimple.ZoneRecordResponse{}, nil)
	}

	mockProvider = dnsimpleProvider{client: mockDNS}

	// Run tests on mock services
	t.Run("Zones", testDnsimpleProviderZones)
	t.Run("Records", testDnsimpleProviderRecords)
	t.Run("ApplyChanges", testDnsimpleProviderApplyChanges)
	t.Run("ApplyChanges/SkipUnknownZone", testDnsimpleProviderApplyChangesSkipsUnknown)
	t.Run("SuitableZone", testDnsimpleSuitableZone)
	t.Run("GetRecordID", testDnsimpleGetRecordID)
}

func testDnsimpleProviderZones(t *testing.T) {
	ctx := context.Background()
	mockProvider.accountID = "1"
	result, err := mockProvider.Zones(ctx)
	assert.Nil(t, err)
	validateDnsimpleZones(t, result, dnsimpleListZonesResponse.Data)

	mockProvider.accountID = "2"
	_, err = mockProvider.Zones(ctx)
	assert.NotNil(t, err)
}

func testDnsimpleProviderRecords(t *testing.T) {
	ctx := context.Background()
	mockProvider.accountID = "1"
	result, err := mockProvider.Records(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(dnsimpleListRecordsResponse.Data), len(result))

	mockProvider.accountID = "2"
	_, err = mockProvider.Records(ctx)
	assert.NotNil(t, err)
}

func testDnsimpleProviderApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.example.com", Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "custom-ttl.example.com", RecordTTL: 60, Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "example-beta.example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "example.example.com", Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
	}

	mockProvider.accountID = "1"
	err := mockProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("Failed to apply changes: %v", err)
	}
}

func testDnsimpleProviderApplyChangesSkipsUnknown(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.not-included.com", Targets: endpoint.Targets{"dasd"}, RecordType: endpoint.RecordTypeCNAME},
	}

	mockProvider.accountID = "1"
	err := mockProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("Failed to ignore unknown zones: %v", err)
	}
}

func testDnsimpleSuitableZone(t *testing.T) {
	ctx := context.Background()
	mockProvider.accountID = "1"
	zones, err := mockProvider.Zones(ctx)
	assert.Nil(t, err)

	zone := dnsimpleSuitableZone("example-beta.example.com", zones)
	assert.Equal(t, zone.Name, "example.com")
}

func TestNewDnsimpleProvider(t *testing.T) {
	os.Setenv("DNSIMPLE_OAUTH", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	_, err := NewDnsimpleProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true)
	if err == nil {
		t.Errorf("Expected to fail new provider on bad token")
	}
	os.Unsetenv("DNSIMPLE_OAUTH")
	if err == nil {
		t.Errorf("Expected to fail new provider on empty token")
	}
}

func testDnsimpleGetRecordID(t *testing.T) {
	var result int64
	var err error

	mockProvider.accountID = "1"
	result, err = mockProvider.GetRecordID(context.Background(), "example.com", "example")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), result)

	result, err = mockProvider.GetRecordID(context.Background(), "example.com", "example-beta")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), result)
}

func validateDnsimpleZones(t *testing.T, zones map[string]dnsimple.Zone, expected []dnsimple.Zone) {
	require.Len(t, zones, len(expected))

	for _, e := range expected {
		assert.Equal(t, zones[int64ToString(e.ID)].Name, e.Name)
	}
}

type mockDnsimpleZoneServiceInterface struct {
	mock.Mock
}

func (_m *mockDnsimpleZoneServiceInterface) CreateRecord(ctx context.Context, accountID string, zoneID string, recordAttributes dnsimple.ZoneRecordAttributes) (*dnsimple.ZoneRecordResponse, error) {
	args := _m.Called(ctx, accountID, zoneID, recordAttributes)
	var r0 *dnsimple.ZoneRecordResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dnsimple.ZoneRecordResponse)
	}

	return r0, args.Error(1)
}

func (_m *mockDnsimpleZoneServiceInterface) DeleteRecord(ctx context.Context, accountID string, zoneID string, recordID int64) (*dnsimple.ZoneRecordResponse, error) {
	args := _m.Called(ctx, accountID, zoneID, recordID)
	var r0 *dnsimple.ZoneRecordResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dnsimple.ZoneRecordResponse)
	}

	return r0, args.Error(1)
}

func (_m *mockDnsimpleZoneServiceInterface) ListRecords(ctx context.Context, accountID string, zoneID string, options *dnsimple.ZoneRecordListOptions) (*dnsimple.ZoneRecordsResponse, error) {
	args := _m.Called(ctx, accountID, zoneID, options)
	var r0 *dnsimple.ZoneRecordsResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dnsimple.ZoneRecordsResponse)
	}

	return r0, args.Error(1)
}

func (_m *mockDnsimpleZoneServiceInterface) ListZones(ctx context.Context, accountID string, options *dnsimple.ZoneListOptions) (*dnsimple.ZonesResponse, error) {
	args := _m.Called(ctx, accountID, options)
	var r0 *dnsimple.ZonesResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dnsimple.ZonesResponse)
	}

	return r0, args.Error(1)
}

func (_m *mockDnsimpleZoneServiceInterface) UpdateRecord(ctx context.Context, accountID string, zoneID string, recordID int64, recordAttributes dnsimple.ZoneRecordAttributes) (*dnsimple.ZoneRecordResponse, error) {
	args := _m.Called(ctx, accountID, zoneID, recordID, recordAttributes)
	var r0 *dnsimple.ZoneRecordResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*dnsimple.ZoneRecordResponse)
	}

	return r0, args.Error(1)
}
