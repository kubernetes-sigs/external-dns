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
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var (
	mockProvider                     dnsimpleProvider
	dnsimpleListZonesResponse        dnsimple.ZonesResponse
	dnsimpleListZonesFromEnvResponse dnsimple.ZonesResponse
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
	firstEnvDefinedZone := dnsimple.Zone{
		ID:        0,
		AccountID: 12345,
		Name:      "example-from-env.com",
	}
	envDefinedZones := []dnsimple.Zone{firstEnvDefinedZone}
	dnsimpleListZonesFromEnvResponse = dnsimple.ZonesResponse{
		Response: dnsimple.Response{Pagination: &dnsimple.Pagination{}},
		Data:     envDefinedZones,
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
	fifthRecord := dnsimple.ZoneRecord{
		ID:       5,
		ZoneID:   "example.com",
		ParentID: 0,
		Name:     "example-ipv6",
		Content:  "fd00::1",
		TTL:      3600,
		Priority: 0,
		Type:     "AAAA",
	}

	records := []dnsimple.ZoneRecord{firstRecord, secondRecord, thirdRecord, fourthRecord, fifthRecord}

	// Setup mock services
	// Note: AnythingOfType doesn't work with interfaces https://github.com/stretchr/testify/issues/519
	mockDNS := &mockDnsimpleZoneServiceInterface{}
	mockDNS.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(&dnsimpleListZonesResponse, nil)
	mockDNS.On("ListZones", t.Context(), "2", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(nil, fmt.Errorf("Account ID not found"))

	for _, record := range records {
		recordName := record.Name
		recordType := record.Type
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

		mockDNS.On("ListRecords", t.Context(), "1", record.ZoneID, &dnsimple.ZoneRecordListOptions{Name: &recordName, Type: &recordType, ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(&dnsimpleRecordResponse, nil)
		mockDNS.On("CreateRecord", t.Context(), "1", record.ZoneID, simpleRecord).Return(&dnsimple.ZoneRecordResponse{}, nil)
		mockDNS.On("DeleteRecord", t.Context(), "1", record.ZoneID, record.ID).Return(&dnsimple.ZoneRecordResponse{}, nil)
		mockDNS.On("UpdateRecord", t.Context(), "1", record.ZoneID, record.ID, simpleRecord).Return(&dnsimple.ZoneRecordResponse{}, nil)
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
	ctx := t.Context()
	mockProvider.accountID = "1"
	result, err := mockProvider.Zones(ctx)
	assert.NoError(t, err)
	validateDnsimpleZones(t, result, dnsimpleListZonesResponse.Data)

	mockProvider.accountID = "2"
	_, err = mockProvider.Zones(ctx)
	assert.Error(t, err)

	mockProvider.accountID = "3"
	t.Setenv("DNSIMPLE_ZONES", "example-from-env.com")
	result, err = mockProvider.Zones(ctx)
	assert.NoError(t, err)
	validateDnsimpleZones(t, result, dnsimpleListZonesFromEnvResponse.Data)

	mockProvider.accountID = "2"
	os.Unsetenv("DNSIMPLE_ZONES")
}

// testDnsimpleProviderRecords drives Records() through a table of cases, each
// with a fresh mock, covering the supported record types, dual-stack hosts,
// apex names, and error propagation from ListZones and ListRecords.
func testDnsimpleProviderRecords(t *testing.T) {
	zonesResponse := &dnsimple.ZonesResponse{
		Response: dnsimple.Response{Pagination: &dnsimple.Pagination{TotalPages: 1}},
		Data:     []dnsimple.Zone{{ID: 1, AccountID: 12345, Name: "example.com"}},
	}
	recordsResponse := func(records ...dnsimple.ZoneRecord) *dnsimple.ZoneRecordsResponse {
		return &dnsimple.ZoneRecordsResponse{
			Response: dnsimple.Response{Pagination: &dnsimple.Pagination{TotalPages: 1}},
			Data:     records,
		}
	}

	tests := []struct {
		name    string
		setup   func(m *mockDnsimpleZoneServiceInterface)
		wantErr bool
		want    []*endpoint.Endpoint
	}{
		{
			name: "all supported types including SRV and NS returned together",
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(zonesResponse, nil)
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(recordsResponse(
					dnsimple.ZoneRecord{ID: 1, ZoneID: "example.com", Name: "a", Content: "127.0.0.1", TTL: 3600, Type: "A"},
					dnsimple.ZoneRecord{ID: 2, ZoneID: "example.com", Name: "aaaa", Content: "fd00::1", TTL: 3600, Type: "AAAA"},
					dnsimple.ZoneRecord{ID: 3, ZoneID: "example.com", Name: "cname", Content: "target", TTL: 7200, Type: "CNAME"},
					dnsimple.ZoneRecord{ID: 4, ZoneID: "example.com", Name: "txt", Content: "hello", TTL: 3600, Type: "TXT"},
					dnsimple.ZoneRecord{ID: 5, ZoneID: "example.com", Name: "srv", Content: "1 10 5060 sip.example.com", TTL: 3600, Type: "SRV"},
					dnsimple.ZoneRecord{ID: 6, ZoneID: "example.com", Name: "ns", Content: "ns1.example.com", TTL: 3600, Type: "NS"},
				), nil)
			},
			want: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("a.example.com", endpoint.RecordTypeA, 3600, "127.0.0.1"),
				endpoint.NewEndpointWithTTL("aaaa.example.com", endpoint.RecordTypeAAAA, 3600, "fd00::1"),
				endpoint.NewEndpointWithTTL("cname.example.com", endpoint.RecordTypeCNAME, 7200, "target"),
				endpoint.NewEndpointWithTTL("txt.example.com", endpoint.RecordTypeTXT, 3600, "hello"),
				endpoint.NewEndpointWithTTL("srv.example.com", endpoint.RecordTypeSRV, 3600, "1 10 5060 sip.example.com"),
				endpoint.NewEndpointWithTTL("ns.example.com", endpoint.RecordTypeNS, 3600, "ns1.example.com"),
			},
		},
		{
			name: "dual-stack host returns both A and AAAA",
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(zonesResponse, nil)
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(recordsResponse(
					dnsimple.ZoneRecord{ID: 1, ZoneID: "example.com", Name: "www", Content: "127.0.0.1", TTL: 3600, Type: "A"},
					dnsimple.ZoneRecord{ID: 2, ZoneID: "example.com", Name: "www", Content: "fd00::1", TTL: 3600, Type: "AAAA"},
				), nil)
			},
			want: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, 3600, "127.0.0.1"),
				endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeAAAA, 3600, "fd00::1"),
			},
		},
		{
			name: "apex domain maps empty name to the zone name",
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(zonesResponse, nil)
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(recordsResponse(
					dnsimple.ZoneRecord{ID: 1, ZoneID: "example.com", Name: "", Content: "127.0.0.1", TTL: 3600, Type: "A"},
				), nil)
			},
			want: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, 3600, "127.0.0.1"),
			},
		},
		{
			name: "unsupported types are skipped",
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(zonesResponse, nil)
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(recordsResponse(
					dnsimple.ZoneRecord{ID: 1, ZoneID: "example.com", Name: "a", Content: "127.0.0.1", TTL: 3600, Type: "A"},
					dnsimple.ZoneRecord{ID: 2, ZoneID: "example.com", Name: "soa", Content: "ns.example.com", TTL: 3600, Type: "SOA"},
				), nil)
			},
			want: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("a.example.com", endpoint.RecordTypeA, 3600, "127.0.0.1"),
			},
		},
		{
			name: "ListZones error is propagated",
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(nil, fmt.Errorf("zones boom"))
			},
			wantErr: true,
		},
		{
			name: "ListRecords error is propagated",
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(zonesResponse, nil)
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(nil, fmt.Errorf("records boom"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDNS := &mockDnsimpleZoneServiceInterface{}
			tt.setup(mockDNS)
			p := dnsimpleProvider{client: mockDNS, accountID: "1"}

			result, err := p.Records(t.Context())
			if tt.wantErr {
				assert.Error(t, err)
				mockDNS.AssertExpectations(t)
				return
			}
			require.NoError(t, err)
			assert.True(t, testutils.SameEndpoints(result, tt.want), "expected %v, got %v", tt.want, result)
			mockDNS.AssertExpectations(t)
		})
	}
}

// testDnsimpleProviderApplyChanges drives ApplyChanges through a table of cases,
// asserting the exact DNSimple API calls via AssertExpectations so a no-op would
// fail. The dual-stack cases exercise GetRecordID's name+type disambiguation.
func testDnsimpleProviderApplyChanges(t *testing.T) {
	zonesResponse := &dnsimple.ZonesResponse{
		Response: dnsimple.Response{Pagination: &dnsimple.Pagination{TotalPages: 1}},
		Data:     []dnsimple.Zone{{ID: 1, AccountID: 12345, Name: "example.com"}},
	}

	// listRecords builds the single-record response GetRecordID expects for a
	// ListRecords lookup of the "www" host scoped to the given type.
	listRecords := func(id int64, recordType string) *dnsimple.ZoneRecordsResponse {
		return &dnsimple.ZoneRecordsResponse{
			Response: dnsimple.Response{Pagination: &dnsimple.Pagination{TotalPages: 1}},
			Data: []dnsimple.ZoneRecord{
				{ID: id, ZoneID: "example.com", Name: "www", Type: recordType},
			},
		}
	}

	tests := []struct {
		name    string
		changes *plan.Changes
		setup   func(m *mockDnsimpleZoneServiceInterface)
	}{
		{
			name: "create A and AAAA",
			changes: &plan.Changes{Create: []*endpoint.Endpoint{
				{DNSName: "www.example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "www.example.com", Targets: endpoint.Targets{"fd00::1"}, RecordType: endpoint.RecordTypeAAAA},
			}},
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("CreateRecord", t.Context(), "1", "example.com", dnsimple.ZoneRecordAttributes{Name: new("www"), Type: "A", Content: "127.0.0.1", TTL: defaultTTL}).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
				m.On("CreateRecord", t.Context(), "1", "example.com", dnsimple.ZoneRecordAttributes{Name: new("www"), Type: "AAAA", Content: "fd00::1", TTL: defaultTTL}).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
			},
		},
		{
			name: "create apex strips name to empty",
			changes: &plan.Changes{Create: []*endpoint.Endpoint{
				{DNSName: "example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
			}},
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("CreateRecord", t.Context(), "1", "example.com", dnsimple.ZoneRecordAttributes{Name: new(""), Type: "A", Content: "127.0.0.1", TTL: defaultTTL}).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
			},
		},
		{
			name: "create with custom ttl",
			changes: &plan.Changes{Create: []*endpoint.Endpoint{
				{DNSName: "ttl.example.com", RecordTTL: 60, Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeCNAME},
			}},
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("CreateRecord", t.Context(), "1", "example.com", dnsimple.ZoneRecordAttributes{Name: new("ttl"), Type: "CNAME", Content: "target", TTL: 60}).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
			},
		},
		{
			name: "update looks up id by name and type",
			changes: &plan.Changes{UpdateNew: []*endpoint.Endpoint{
				{DNSName: "www.example.com", Targets: endpoint.Targets{"127.0.0.2"}, RecordType: endpoint.RecordTypeA},
			}},
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{Name: new("www"), Type: new("A"), ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(listRecords(10, "A"), nil).Once()
				m.On("UpdateRecord", t.Context(), "1", "example.com", int64(10), dnsimple.ZoneRecordAttributes{Name: new("www"), Type: "A", Content: "127.0.0.2", TTL: defaultTTL}).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
			},
		},
		{
			name: "delete looks up id by name and type",
			changes: &plan.Changes{Delete: []*endpoint.Endpoint{
				{DNSName: "www.example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
			}},
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{Name: new("www"), Type: new("A"), ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(listRecords(10, "A"), nil).Once()
				m.On("DeleteRecord", t.Context(), "1", "example.com", int64(10)).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
			},
		},
		{
			name: "dual-stack delete resolves the right id per type",
			changes: &plan.Changes{Delete: []*endpoint.Endpoint{
				{DNSName: "www.example.com", Targets: endpoint.Targets{"127.0.0.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "www.example.com", Targets: endpoint.Targets{"fd00::1"}, RecordType: endpoint.RecordTypeAAAA},
			}},
			setup: func(m *mockDnsimpleZoneServiceInterface) {
				// Same name, different types: each lookup must be scoped by type
				// and return that type's distinct record ID.
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{Name: new("www"), Type: new("A"), ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(listRecords(11, "A"), nil).Once()
				m.On("ListRecords", t.Context(), "1", "example.com", &dnsimple.ZoneRecordListOptions{Name: new("www"), Type: new("AAAA"), ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(listRecords(12, "AAAA"), nil).Once()
				m.On("DeleteRecord", t.Context(), "1", "example.com", int64(11)).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
				m.On("DeleteRecord", t.Context(), "1", "example.com", int64(12)).Return(&dnsimple.ZoneRecordResponse{}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDNS := &mockDnsimpleZoneServiceInterface{}
			mockDNS.On("ListZones", t.Context(), "1", &dnsimple.ZoneListOptions{ListOptions: dnsimple.ListOptions{Page: new(1)}}).Return(zonesResponse, nil)
			tt.setup(mockDNS)

			p := dnsimpleProvider{client: mockDNS, accountID: "1"}
			require.NoError(t, p.ApplyChanges(t.Context(), tt.changes))
			mockDNS.AssertExpectations(t)
		})
	}
}

func testDnsimpleProviderApplyChangesSkipsUnknown(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.not-included.com", Targets: endpoint.Targets{"dasd"}, RecordType: endpoint.RecordTypeCNAME},
	}

	mockProvider.accountID = "1"
	err := mockProvider.ApplyChanges(t.Context(), changes)
	if err != nil {
		t.Errorf("Failed to ignore unknown zones: %v", err)
	}
}

func testDnsimpleSuitableZone(t *testing.T) {
	ctx := t.Context()
	mockProvider.accountID = "1"
	zones, err := mockProvider.Zones(ctx)
	require.NoError(t, err)

	zone := dnsimpleSuitableZone("example-beta.example.com", zones)
	assert.Equal(t, "example.com", zone.Name)

	t.Setenv("DNSIMPLE_ZONES", "environment-example.com,example.environment-example.com")
	mockProvider.accountID = "3"
	zones, err = mockProvider.Zones(ctx)
	require.NoError(t, err)

	zone = dnsimpleSuitableZone("hello.example.environment-example.com", zones)
	assert.Equal(t, "example.environment-example.com", zone.Name)

	_ = os.Unsetenv("DNSIMPLE_ZONES")
	mockProvider.accountID = "1"
}

func TestNewProvider(t *testing.T) {
	t.Setenv("DNSIMPLE_OAUTH", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	_, err := newProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true)
	if err == nil {
		t.Errorf("Expected to fail new provider on bad token")
	}

	_ = os.Unsetenv("DNSIMPLE_OAUTH")
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true)
	if err == nil {
		t.Errorf("Expected to fail new provider on empty token")
	}

	t.Setenv("DNSIMPLE_OAUTH", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	t.Setenv("DNSIMPLE_ACCOUNT_ID", "12345678")
	providerTypedProvider, err := newProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true)
	dnsimpleTypedProvider := providerTypedProvider.(*dnsimpleProvider)
	if err != nil {
		t.Errorf("Unexpected error thrown when testing NewDnsimpleProvider with the DNSIMPLE_ACCOUNT_ID environment variable set")
	}
	assert.Equal(t, "12345678", dnsimpleTypedProvider.accountID)
	os.Unsetenv("DNSIMPLE_OAUTH")
	os.Unsetenv("DNSIMPLE_ACCOUNT_ID")
}

func testDnsimpleGetRecordID(t *testing.T) {
	var result int64
	var err error

	mockProvider.accountID = "1"
	result, err = mockProvider.GetRecordID(t.Context(), "example.com", "example", "CNAME")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), result)

	result, err = mockProvider.GetRecordID(t.Context(), "example.com", "example-beta", "A")
	assert.NoError(t, err)
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
