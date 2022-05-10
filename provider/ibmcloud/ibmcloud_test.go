/*
Copyright 2022 The Kubernetes Authors.

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

package ibmcloud

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"

	"github.com/IBM/networking-go-sdk/dnssvcsv1"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

func NewMockIBMCloudDNSAPI() *mockIbmcloudClientInterface {
	//  Setup public example responses
	firstPublicRecord := dnsrecordsv1.DnsrecordDetails{
		ID:      core.StringPtr("123"),
		Name:    core.StringPtr("test.example.com"),
		Type:    core.StringPtr("A"),
		Content: core.StringPtr("1.2.3.4"),
		Proxied: core.BoolPtr(true),
		TTL:     core.Int64Ptr(int64(120)),
	}
	secondPublicRecord := dnsrecordsv1.DnsrecordDetails{
		ID:      core.StringPtr("456"),
		Name:    core.StringPtr("test.example.com"),
		Type:    core.StringPtr("TXT"),
		Proxied: core.BoolPtr(false),
		Content: core.StringPtr("\"heritage=external-dns,external-dns/owner=tower-pdns\""),
		TTL:     core.Int64Ptr(int64(120)),
	}
	publicRecordsResult := []dnsrecordsv1.DnsrecordDetails{firstPublicRecord, secondPublicRecord}
	publicRecordsResultInfo := &dnsrecordsv1.ResultInfo{
		Page:       core.Int64Ptr(int64(1)),
		TotalCount: core.Int64Ptr(int64(1)),
	}

	publicRecordsResp := &dnsrecordsv1.ListDnsrecordsResp{
		Result:     publicRecordsResult,
		ResultInfo: publicRecordsResultInfo,
	}
	// Setup private example responses
	firstPrivateZone := dnssvcsv1.Dnszone{
		ID:    core.StringPtr("123"),
		Name:  core.StringPtr("example.com"),
		State: core.StringPtr(zoneStatePendingNetwork),
	}

	secondPrivateZone := dnssvcsv1.Dnszone{
		ID:    core.StringPtr("456"),
		Name:  core.StringPtr("example1.com"),
		State: core.StringPtr(zoneStateActive),
	}
	privateZones := []dnssvcsv1.Dnszone{firstPrivateZone, secondPrivateZone}
	listZonesResp := &dnssvcsv1.ListDnszones{
		Dnszones: privateZones,
	}
	firstPrivateRecord := dnssvcsv1.ResourceRecord{
		ID:    core.StringPtr("123"),
		Name:  core.StringPtr("test.example.com"),
		Type:  core.StringPtr("A"),
		Rdata: map[string]interface{}{"ip": "1.2.3.4"},
		TTL:   core.Int64Ptr(int64(120)),
	}
	secondPrivateRecord := dnssvcsv1.ResourceRecord{
		ID:    core.StringPtr("456"),
		Name:  core.StringPtr("testCNAME.example.com"),
		Type:  core.StringPtr("CNAME"),
		Rdata: map[string]interface{}{"cname": "test.example.com"},
		TTL:   core.Int64Ptr(int64(120)),
	}
	thirdPrivateRecord := dnssvcsv1.ResourceRecord{
		ID:    core.StringPtr("789"),
		Name:  core.StringPtr("test.example.com"),
		Type:  core.StringPtr("TXT"),
		Rdata: map[string]interface{}{"text": "\"heritage=external-dns,external-dns/owner=tower-pdns\""},
		TTL:   core.Int64Ptr(int64(120)),
	}
	privateRecords := []dnssvcsv1.ResourceRecord{firstPrivateRecord, secondPrivateRecord, thirdPrivateRecord}
	privateRecordsResop := &dnssvcsv1.ListResourceRecords{
		ResourceRecords: privateRecords,
		Offset:          core.Int64Ptr(int64(0)),
		TotalCount:      core.Int64Ptr(int64(1)),
	}

	// Setup record rData
	inputARecord := &dnssvcsv1.ResourceRecordInputRdataRdataARecord{
		Ip: core.StringPtr("1.2.3.4"),
	}
	inputCnameRecord := &dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord{
		Cname: core.StringPtr("test.example.com"),
	}
	inputTxtRecord := &dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord{
		Text: core.StringPtr("test"),
	}

	updateARecord := &dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord{
		Ip: core.StringPtr("1.2.3.4"),
	}
	updateCnameRecord := &dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord{
		Cname: core.StringPtr("test.example.com"),
	}
	updateTxtRecord := &dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord{
		Text: core.StringPtr("test"),
	}

	// Setup mock services
	mockDNSClient := &mockIbmcloudClientInterface{}
	mockDNSClient.On("CreateDNSRecordWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("UpdateDNSRecordWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("DeleteDNSRecordWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("ListAllDDNSRecordsWithContext", mock.Anything, mock.Anything).Return(publicRecordsResp, nil, nil)
	mockDNSClient.On("ListDnszonesWithContext", mock.Anything, mock.Anything).Return(listZonesResp, nil, nil)
	mockDNSClient.On("GetDnszoneWithContext", mock.Anything, mock.Anything).Return(&firstPrivateZone, nil, nil)
	mockDNSClient.On("ListResourceRecordsWithContext", mock.Anything, mock.Anything).Return(privateRecordsResop, nil, nil)
	mockDNSClient.On("CreatePermittedNetworkWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("CreateResourceRecordWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("DeleteResourceRecordWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("UpdateResourceRecordWithContext", mock.Anything, mock.Anything).Return(nil, nil, nil)
	mockDNSClient.On("NewResourceRecordInputRdataRdataARecord", mock.Anything).Return(inputARecord, nil)
	mockDNSClient.On("NewResourceRecordInputRdataRdataCnameRecord", mock.Anything).Return(inputCnameRecord, nil)
	mockDNSClient.On("NewResourceRecordInputRdataRdataTxtRecord", mock.Anything).Return(inputTxtRecord, nil)
	mockDNSClient.On("NewResourceRecordUpdateInputRdataRdataARecord", mock.Anything).Return(updateARecord, nil)
	mockDNSClient.On("NewResourceRecordUpdateInputRdataRdataCnameRecord", mock.Anything).Return(updateCnameRecord, nil)
	mockDNSClient.On("NewResourceRecordUpdateInputRdataRdataTxtRecord", mock.Anything).Return(updateTxtRecord, nil)

	return mockDNSClient
}

func newTestIBMCloudProvider(private bool) *IBMCloudProvider {
	mockSource := &mockSource{}
	endpoints := []*endpoint.Endpoint{
		{
			DNSName: "new.example.com",
			Targets: endpoint.Targets{"4.3.2.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  "ibmcloud-vpc",
					Value: "crn:v1:staging:public:is:us-south:a/0821fa9f9ebcc7b7c9a0d6e9bf9442a4::vpc:be33cdad-9a03-4bfa-82ca-eadb9f1de688",
				},
			},
		},
	}
	mockSource.On("Endpoints", mock.Anything).Return(endpoints, nil, nil)

	domainFilterTest := endpoint.NewDomainFilter([]string{"example.com"})

	return &IBMCloudProvider{
		Client:       NewMockIBMCloudDNSAPI(),
		source:       mockSource,
		domainFilter: domainFilterTest,
		DryRun:       false,
		instanceID:   "test123",
		privateZone:  private,
	}
}

func TestPublic_Records(t *testing.T) {
	p := newTestIBMCloudProvider(false)
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %++v", *endpoint)
		}
	}
}

func TestPrivate_Records(t *testing.T) {
	p := newTestIBMCloudProvider(true)
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %++v", *endpoint)
		}
	}
}

func TestPublic_ApplyChanges(t *testing.T) {
	p := newTestIBMCloudProvider(false)

	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "newA.example.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  "ibmcloud-proxied",
						Value: "false",
					},
				},
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: "A",
				RecordTTL:  180,
				Targets:    endpoint.NewTargets("1.2.3.4"),
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  "ibmcloud-proxied",
						Value: "false",
					},
				},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: "A",
				RecordTTL:  180,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  "ibmcloud-proxied",
						Value: "true",
					},
				},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=tower-pdns\""),
			},
		},
	}
	ctx := context.Background()
	err := p.ApplyChanges(ctx, &changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestPrivate_ApplyChanges(t *testing.T) {
	p := newTestIBMCloudProvider(true)

	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "newA.example.com",
				RecordType: "A",
				RecordTTL:  120,
				Targets:    endpoint.NewTargets("4.3.2.1"),
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  "ibmcloud-vpc",
						Value: "crn:v1:staging:public:is:us-south:a/0821fa9f9ebcc7b7c9a0d6e9bf9442a4::vpc:be33cdad-9a03-4bfa-82ca-eadb9f1de688",
					},
				},
			},
			{
				DNSName:    "newCNAME.example.com",
				RecordType: "CNAME",
				RecordTTL:  180,
				Targets:    endpoint.NewTargets("newA.example.com"),
			},
			{
				DNSName:    "newTXT.example.com",
				RecordType: "TXT",
				RecordTTL:  240,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=tower-pdns\""),
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: "A",
				RecordTTL:  180,
				Targets:    endpoint.NewTargets("1.2.3.4"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: "A",
				RecordTTL:  180,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=tower-pdns\""),
			},
		},
	}
	ctx := context.Background()
	err := p.ApplyChanges(ctx, &changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestAdjustEndpoints(t *testing.T) {
	p := newTestIBMCloudProvider(false)
	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "test.example.com",
			Targets:    endpoint.Targets{"1.2.3.4"},
			RecordType: endpoint.RecordTypeA,
			RecordTTL:  300,
			Labels:     endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  "ibmcloud-proxied",
					Value: "true",
				},
			},
		},
	}

	ep := p.AdjustEndpoints(endpoints)

	assert.Equal(t, endpoint.TTL(0), ep[0].RecordTTL)
	assert.Equal(t, "test.example.com", ep[0].DNSName)
}

func TestPrivateZone_withFilterID(t *testing.T) {
	p := newTestIBMCloudProvider(true)
	p.zoneIDFilter = provider.NewZoneIDFilter([]string{"123", "456"})

	zones, err := p.privateZones(context.Background())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	} else {
		if len(zones) != 2 {
			t.Errorf("Incorrect number of zones: %d", len(zones))
		}
		for _, zone := range zones {
			t.Logf("zone %s", *zone.ID)
		}
	}
}

func TestPublicConfig_Validate(t *testing.T) {
	// mock http server
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer GinkgoRecover()
		time.Sleep(0)

		// Set mock response
		res.Header().Set("Content-type", "application/json")
		res.WriteHeader(200)
		fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "123", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"]}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
	}))
	zoneIDFilterTest := provider.NewZoneIDFilter([]string{"123"})
	domainFilterTest := endpoint.NewDomainFilter([]string{"example.com"})
	cfg := &ibmcloudConfig{
		Endpoint: testServer.URL,
		CRN:      "crn:v1:bluemix:public:internet-svcs:global:a/bcf1865e99742d38d2d5fc3fb80a5496:a6338168-9510-4951-9d67-425612de96f0::",
	}
	crn := cfg.CRN
	authenticator := &core.NoAuthAuthenticator{}
	service, isPrivate, err := cfg.Validate(authenticator, domainFilterTest, provider.NewZoneIDFilter([]string{""}))
	assert.NoError(t, err)
	assert.Equal(t, false, isPrivate)
	assert.Equal(t, crn, *service.publicRecordsService.Crn)
	assert.Equal(t, "123", *service.publicRecordsService.ZoneIdentifier)

	service, isPrivate, err = cfg.Validate(authenticator, endpoint.NewDomainFilter([]string{""}), zoneIDFilterTest)
	assert.NoError(t, err)
	assert.Equal(t, false, isPrivate)
	assert.Equal(t, crn, *service.publicRecordsService.Crn)
	assert.Equal(t, "123", *service.publicRecordsService.ZoneIdentifier)

	testServer.Close()
}

func TestPrivateConfig_Validate(t *testing.T) {
	zoneIDFilterTest := provider.NewZoneIDFilter([]string{"123"})
	domainFilterTest := endpoint.NewDomainFilter([]string{"example.com"})
	authenticator := &core.NoAuthAuthenticator{}
	cfg := &ibmcloudConfig{
		Endpoint: "XXX",
		CRN:      "crn:v1:bluemix:public:dns-svcs:global:a/bcf1865e99742d38d2d5fc3fb80a5496:a6338168-9510-4951-9d67-425612de96f0::",
	}
	_, isPrivate, err := cfg.Validate(authenticator, domainFilterTest, zoneIDFilterTest)
	assert.NoError(t, err)
	assert.Equal(t, true, isPrivate)
}

// mockIbmcloudClientInterface is an autogenerated mock type for the ibmcloudClient type
type mockIbmcloudClientInterface struct {
	mock.Mock
}

// CreateDNSRecordWithContext provides a mock function with given fields: ctx, createDnsRecordOptions
func (_m *mockIbmcloudClientInterface) CreateDNSRecordWithContext(ctx context.Context, createDnsRecordOptions *dnsrecordsv1.CreateDnsRecordOptions) (*dnsrecordsv1.DnsrecordResp, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, createDnsRecordOptions)

	var r0 *dnsrecordsv1.DnsrecordResp
	if rf, ok := ret.Get(0).(func(context.Context, *dnsrecordsv1.CreateDnsRecordOptions) *dnsrecordsv1.DnsrecordResp); ok {
		r0 = rf(ctx, createDnsRecordOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnsrecordsv1.DnsrecordResp)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnsrecordsv1.CreateDnsRecordOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, createDnsRecordOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnsrecordsv1.CreateDnsRecordOptions) error); ok {
		r2 = rf(ctx, createDnsRecordOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreatePermittedNetworkWithContext provides a mock function with given fields: ctx, createPermittedNetworkOptions
func (_m *mockIbmcloudClientInterface) CreatePermittedNetworkWithContext(ctx context.Context, createPermittedNetworkOptions *dnssvcsv1.CreatePermittedNetworkOptions) (*dnssvcsv1.PermittedNetwork, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, createPermittedNetworkOptions)

	var r0 *dnssvcsv1.PermittedNetwork
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.CreatePermittedNetworkOptions) *dnssvcsv1.PermittedNetwork); ok {
		r0 = rf(ctx, createPermittedNetworkOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.PermittedNetwork)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.CreatePermittedNetworkOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, createPermittedNetworkOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnssvcsv1.CreatePermittedNetworkOptions) error); ok {
		r2 = rf(ctx, createPermittedNetworkOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateResourceRecordWithContext provides a mock function with given fields: ctx, createResourceRecordOptions
func (_m *mockIbmcloudClientInterface) CreateResourceRecordWithContext(ctx context.Context, createResourceRecordOptions *dnssvcsv1.CreateResourceRecordOptions) (*dnssvcsv1.ResourceRecord, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, createResourceRecordOptions)

	var r0 *dnssvcsv1.ResourceRecord
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.CreateResourceRecordOptions) *dnssvcsv1.ResourceRecord); ok {
		r0 = rf(ctx, createResourceRecordOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecord)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.CreateResourceRecordOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, createResourceRecordOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnssvcsv1.CreateResourceRecordOptions) error); ok {
		r2 = rf(ctx, createResourceRecordOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteDNSRecordWithContext provides a mock function with given fields: ctx, deleteDnsRecordOptions
func (_m *mockIbmcloudClientInterface) DeleteDNSRecordWithContext(ctx context.Context, deleteDnsRecordOptions *dnsrecordsv1.DeleteDnsRecordOptions) (*dnsrecordsv1.DeleteDnsrecordResp, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, deleteDnsRecordOptions)

	var r0 *dnsrecordsv1.DeleteDnsrecordResp
	if rf, ok := ret.Get(0).(func(context.Context, *dnsrecordsv1.DeleteDnsRecordOptions) *dnsrecordsv1.DeleteDnsrecordResp); ok {
		r0 = rf(ctx, deleteDnsRecordOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnsrecordsv1.DeleteDnsrecordResp)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnsrecordsv1.DeleteDnsRecordOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, deleteDnsRecordOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnsrecordsv1.DeleteDnsRecordOptions) error); ok {
		r2 = rf(ctx, deleteDnsRecordOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteResourceRecordWithContext provides a mock function with given fields: ctx, deleteResourceRecordOptions
func (_m *mockIbmcloudClientInterface) DeleteResourceRecordWithContext(ctx context.Context, deleteResourceRecordOptions *dnssvcsv1.DeleteResourceRecordOptions) (*core.DetailedResponse, error) {
	ret := _m.Called(ctx, deleteResourceRecordOptions)

	var r0 *core.DetailedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.DeleteResourceRecordOptions) *core.DetailedResponse); ok {
		r0 = rf(ctx, deleteResourceRecordOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.DetailedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.DeleteResourceRecordOptions) error); ok {
		r1 = rf(ctx, deleteResourceRecordOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDnszoneWithContext provides a mock function with given fields: ctx, getDnszoneOptions
func (_m *mockIbmcloudClientInterface) GetDnszoneWithContext(ctx context.Context, getDnszoneOptions *dnssvcsv1.GetDnszoneOptions) (*dnssvcsv1.Dnszone, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, getDnszoneOptions)

	var r0 *dnssvcsv1.Dnszone
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.GetDnszoneOptions) *dnssvcsv1.Dnszone); ok {
		r0 = rf(ctx, getDnszoneOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.Dnszone)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.GetDnszoneOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, getDnszoneOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnssvcsv1.GetDnszoneOptions) error); ok {
		r2 = rf(ctx, getDnszoneOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAllDDNSRecordsWithContext provides a mock function with given fields: ctx, listAllDnsRecordsOptions
func (_m *mockIbmcloudClientInterface) ListAllDDNSRecordsWithContext(ctx context.Context, listAllDnsRecordsOptions *dnsrecordsv1.ListAllDnsRecordsOptions) (*dnsrecordsv1.ListDnsrecordsResp, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, listAllDnsRecordsOptions)

	var r0 *dnsrecordsv1.ListDnsrecordsResp
	if rf, ok := ret.Get(0).(func(context.Context, *dnsrecordsv1.ListAllDnsRecordsOptions) *dnsrecordsv1.ListDnsrecordsResp); ok {
		r0 = rf(ctx, listAllDnsRecordsOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnsrecordsv1.ListDnsrecordsResp)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnsrecordsv1.ListAllDnsRecordsOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, listAllDnsRecordsOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnsrecordsv1.ListAllDnsRecordsOptions) error); ok {
		r2 = rf(ctx, listAllDnsRecordsOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListDnszonesWithContext provides a mock function with given fields: ctx, listDnszonesOptions
func (_m *mockIbmcloudClientInterface) ListDnszonesWithContext(ctx context.Context, listDnszonesOptions *dnssvcsv1.ListDnszonesOptions) (*dnssvcsv1.ListDnszones, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, listDnszonesOptions)

	var r0 *dnssvcsv1.ListDnszones
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.ListDnszonesOptions) *dnssvcsv1.ListDnszones); ok {
		r0 = rf(ctx, listDnszonesOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ListDnszones)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.ListDnszonesOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, listDnszonesOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnssvcsv1.ListDnszonesOptions) error); ok {
		r2 = rf(ctx, listDnszonesOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListResourceRecordsWithContext provides a mock function with given fields: ctx, listResourceRecordsOptions
func (_m *mockIbmcloudClientInterface) ListResourceRecordsWithContext(ctx context.Context, listResourceRecordsOptions *dnssvcsv1.ListResourceRecordsOptions) (*dnssvcsv1.ListResourceRecords, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, listResourceRecordsOptions)

	var r0 *dnssvcsv1.ListResourceRecords
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.ListResourceRecordsOptions) *dnssvcsv1.ListResourceRecords); ok {
		r0 = rf(ctx, listResourceRecordsOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ListResourceRecords)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.ListResourceRecordsOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, listResourceRecordsOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnssvcsv1.ListResourceRecordsOptions) error); ok {
		r2 = rf(ctx, listResourceRecordsOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewResourceRecordInputRdataRdataARecord provides a mock function with given fields: ip
func (_m *mockIbmcloudClientInterface) NewResourceRecordInputRdataRdataARecord(ip string) (*dnssvcsv1.ResourceRecordInputRdataRdataARecord, error) {
	ret := _m.Called(ip)

	var r0 *dnssvcsv1.ResourceRecordInputRdataRdataARecord
	if rf, ok := ret.Get(0).(func(string) *dnssvcsv1.ResourceRecordInputRdataRdataARecord); ok {
		r0 = rf(ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecordInputRdataRdataARecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResourceRecordInputRdataRdataCnameRecord provides a mock function with given fields: cname
func (_m *mockIbmcloudClientInterface) NewResourceRecordInputRdataRdataCnameRecord(cname string) (*dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord, error) {
	ret := _m.Called(cname)

	var r0 *dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord
	if rf, ok := ret.Get(0).(func(string) *dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord); ok {
		r0 = rf(cname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResourceRecordInputRdataRdataTxtRecord provides a mock function with given fields: text
func (_m *mockIbmcloudClientInterface) NewResourceRecordInputRdataRdataTxtRecord(text string) (*dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord, error) {
	ret := _m.Called(text)

	var r0 *dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord
	if rf, ok := ret.Get(0).(func(string) *dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord); ok {
		r0 = rf(text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResourceRecordUpdateInputRdataRdataARecord provides a mock function with given fields: ip
func (_m *mockIbmcloudClientInterface) NewResourceRecordUpdateInputRdataRdataARecord(ip string) (*dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord, error) {
	ret := _m.Called(ip)

	var r0 *dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord
	if rf, ok := ret.Get(0).(func(string) *dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord); ok {
		r0 = rf(ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResourceRecordUpdateInputRdataRdataCnameRecord provides a mock function with given fields: cname
func (_m *mockIbmcloudClientInterface) NewResourceRecordUpdateInputRdataRdataCnameRecord(cname string) (*dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord, error) {
	ret := _m.Called(cname)

	var r0 *dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord
	if rf, ok := ret.Get(0).(func(string) *dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord); ok {
		r0 = rf(cname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResourceRecordUpdateInputRdataRdataTxtRecord provides a mock function with given fields: text
func (_m *mockIbmcloudClientInterface) NewResourceRecordUpdateInputRdataRdataTxtRecord(text string) (*dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord, error) {
	ret := _m.Called(text)

	var r0 *dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord
	if rf, ok := ret.Get(0).(func(string) *dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord); ok {
		r0 = rf(text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDNSRecordWithContext provides a mock function with given fields: ctx, updateDnsRecordOptions
func (_m *mockIbmcloudClientInterface) UpdateDNSRecordWithContext(ctx context.Context, updateDnsRecordOptions *dnsrecordsv1.UpdateDnsRecordOptions) (*dnsrecordsv1.DnsrecordResp, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, updateDnsRecordOptions)

	var r0 *dnsrecordsv1.DnsrecordResp
	if rf, ok := ret.Get(0).(func(context.Context, *dnsrecordsv1.UpdateDnsRecordOptions) *dnsrecordsv1.DnsrecordResp); ok {
		r0 = rf(ctx, updateDnsRecordOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnsrecordsv1.DnsrecordResp)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnsrecordsv1.UpdateDnsRecordOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, updateDnsRecordOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnsrecordsv1.UpdateDnsRecordOptions) error); ok {
		r2 = rf(ctx, updateDnsRecordOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateResourceRecordWithContext provides a mock function with given fields: ctx, updateResourceRecordOptions
func (_m *mockIbmcloudClientInterface) UpdateResourceRecordWithContext(ctx context.Context, updateResourceRecordOptions *dnssvcsv1.UpdateResourceRecordOptions) (*dnssvcsv1.ResourceRecord, *core.DetailedResponse, error) {
	ret := _m.Called(ctx, updateResourceRecordOptions)

	var r0 *dnssvcsv1.ResourceRecord
	if rf, ok := ret.Get(0).(func(context.Context, *dnssvcsv1.UpdateResourceRecordOptions) *dnssvcsv1.ResourceRecord); ok {
		r0 = rf(ctx, updateResourceRecordOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnssvcsv1.ResourceRecord)
		}
	}

	var r1 *core.DetailedResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dnssvcsv1.UpdateResourceRecordOptions) *core.DetailedResponse); ok {
		r1 = rf(ctx, updateResourceRecordOptions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.DetailedResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dnssvcsv1.UpdateResourceRecordOptions) error); ok {
		r2 = rf(ctx, updateResourceRecordOptions)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockSource struct {
	mock.Mock
}

// Endpoints provides a mock function with given fields: ctx
func (_m *mockSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	ret := _m.Called(ctx)

	var r0 []*endpoint.Endpoint
	if rf, ok := ret.Get(0).(func(context.Context) []*endpoint.Endpoint); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*endpoint.Endpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddEventHandler provides a mock function with given fields: _a0, _a1
func (_m *mockSource) AddEventHandler(_a0 context.Context, _a1 func()) {
	_m.Called(_a0, _a1)
}
