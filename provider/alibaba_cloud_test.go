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
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"

	"sigs.k8s.io/external-dns/endpoint"

	"sigs.k8s.io/external-dns/plan"
)

type MockAlibabaCloudDNSAPI struct {
	records []alidns.Record
}

func NewMockAlibabaCloudDNSAPI() *MockAlibabaCloudDNSAPI {
	api := MockAlibabaCloudDNSAPI{}
	api.records = []alidns.Record{
		{
			RecordId:   "1",
			DomainName: "container-service.top",
			Type:       "A",
			TTL:        300,
			RR:         "abc",
			Value:      "1.2.3.4",
		},
		{
			RecordId:   "2",
			DomainName: "container-service.top",
			Type:       "TXT",
			TTL:        300,
			RR:         "abc",
			Value:      "heritage=external-dns;external-dns/owner=default",
		},
	}
	return &api
}

func (m *MockAlibabaCloudDNSAPI) AddDomainRecord(request *alidns.AddDomainRecordRequest) (response *alidns.AddDomainRecordResponse, err error) {
	ttl, _ := request.TTL.GetValue()
	m.records = append(m.records, alidns.Record{
		RecordId:   "3",
		DomainName: request.DomainName,
		Type:       request.Type,
		TTL:        ttl,
		RR:         request.RR,
		Value:      request.Value,
	})
	response = alidns.CreateAddDomainRecordResponse()
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) DeleteDomainRecord(request *alidns.DeleteDomainRecordRequest) (response *alidns.DeleteDomainRecordResponse, err error) {
	var result []alidns.Record
	for _, record := range m.records {
		if record.RecordId != request.RecordId {
			result = append(result, record)
		}
	}
	m.records = result
	response = alidns.CreateDeleteDomainRecordResponse()
	response.RecordId = request.RecordId
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) UpdateDomainRecord(request *alidns.UpdateDomainRecordRequest) (response *alidns.UpdateDomainRecordResponse, err error) {
	ttl, _ := request.TTL.GetValue()
	for i := range m.records {
		if m.records[i].RecordId == request.RecordId {
			m.records[i].TTL = ttl
		}
	}
	response = alidns.CreateUpdateDomainRecordResponse()
	response.RecordId = request.RecordId
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) DescribeDomains(request *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error) {
	var result alidns.Domains
	for _, record := range m.records {
		domain := alidns.Domain{}
		domain.DomainName = record.DomainName
	}
	response = alidns.CreateDescribeDomainsResponse()
	response.Domains = result
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) DescribeDomainRecords(request *alidns.DescribeDomainRecordsRequest) (response *alidns.DescribeDomainRecordsResponse, err error) {
	var result []alidns.Record
	for _, record := range m.records {
		if record.DomainName == request.DomainName {
			result = append(result, record)
		}
	}
	response = alidns.CreateDescribeDomainRecordsResponse()
	response.DomainRecords.Record = result
	return response, nil
}

type MockAlibabaCloudPrivateZoneAPI struct {
	zone    pvtz.Zone
	records []pvtz.Record
}

func NewMockAlibabaCloudPrivateZoneAPI() *MockAlibabaCloudPrivateZoneAPI {
	vpc := pvtz.Vpc{
		RegionId: "cn-beijing",
		VpcId:    "vpc-xxxxxx",
	}

	api := MockAlibabaCloudPrivateZoneAPI{zone: pvtz.Zone{
		ZoneId:   "test-zone",
		ZoneName: "container-service.top",
		Vpcs: pvtz.Vpcs{
			Vpc: []pvtz.Vpc{vpc},
		},
	}}

	api.records = []pvtz.Record{
		{
			RecordId: 1,
			Type:     "A",
			Ttl:      300,
			Rr:       "abc",
			Value:    "1.2.3.4",
		},
		{
			RecordId: 2,
			Type:     "TXT",
			Ttl:      300,
			Rr:       "abc",
			Value:    "heritage=external-dns;external-dns/owner=default",
		},
	}
	return &api
}

func (m *MockAlibabaCloudPrivateZoneAPI) AddZoneRecord(request *pvtz.AddZoneRecordRequest) (response *pvtz.AddZoneRecordResponse, err error) {
	ttl, _ := request.Ttl.GetValue()
	m.records = append(m.records, pvtz.Record{
		RecordId: 3,
		Type:     request.Type,
		Ttl:      ttl,
		Rr:       request.Rr,
		Value:    request.Value,
	})
	response = pvtz.CreateAddZoneRecordResponse()
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DeleteZoneRecord(request *pvtz.DeleteZoneRecordRequest) (response *pvtz.DeleteZoneRecordResponse, err error) {
	recordID, _ := request.RecordId.GetValue()

	var result []pvtz.Record
	for _, record := range m.records {
		if record.RecordId != recordID {
			result = append(result, record)
		}
	}
	m.records = result
	response = pvtz.CreateDeleteZoneRecordResponse()
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) UpdateZoneRecord(request *pvtz.UpdateZoneRecordRequest) (response *pvtz.UpdateZoneRecordResponse, err error) {
	recordID, _ := request.RecordId.GetValue()
	ttl, _ := request.Ttl.GetValue()
	for i := range m.records {
		if m.records[i].RecordId == recordID {
			m.records[i].Ttl = ttl
		}
	}
	response = pvtz.CreateUpdateZoneRecordResponse()
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DescribeZoneRecords(request *pvtz.DescribeZoneRecordsRequest) (response *pvtz.DescribeZoneRecordsResponse, err error) {
	response = pvtz.CreateDescribeZoneRecordsResponse()
	response.Records.Record = append(response.Records.Record, m.records...)
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DescribeZones(request *pvtz.DescribeZonesRequest) (response *pvtz.DescribeZonesResponse, err error) {
	response = pvtz.CreateDescribeZonesResponse()
	response.Zones.Zone = append(response.Zones.Zone, m.zone)
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DescribeZoneInfo(request *pvtz.DescribeZoneInfoRequest) (response *pvtz.DescribeZoneInfoResponse, err error) {
	response = pvtz.CreateDescribeZoneInfoResponse()
	response.ZoneId = m.zone.ZoneId
	response.ZoneName = m.zone.ZoneName
	response.BindVpcs = pvtz.BindVpcs{Vpc: m.zone.Vpcs.Vpc}
	return response, nil
}

func newTestAlibabaCloudProvider(private bool) *AlibabaCloudProvider {
	cfg := alibabaCloudConfig{
		RegionID:        "cn-beijing",
		AccessKeyID:     "xxxxxx",
		AccessKeySecret: "xxxxxx",
		VPCID:           "vpc-xxxxxx",
	}
	//
	//dnsClient, _ := alidns.NewClientWithAccessKey(
	//	cfg.RegionID,
	//	cfg.AccessKeyID,
	//	cfg.AccessKeySecret,
	//)
	//
	//pvtzClient, _ := pvtz.NewClientWithAccessKey(
	//	"cn-hangzhou",
	//	cfg.AccessKeyID,
	//	cfg.AccessKeySecret,
	//)
	domainFilterTest := endpoint.NewDomainFilter([]string{"container-service.top.", "example.org"})

	return &AlibabaCloudProvider{
		domainFilter: domainFilterTest,
		vpcID:        cfg.VPCID,
		dryRun:       false,
		dnsClient:    NewMockAlibabaCloudDNSAPI(),
		pvtzClient:   NewMockAlibabaCloudPrivateZoneAPI(),
		privateZone:  private,
	}
}

func TestAlibabaCloudPrivateProvider_Records(t *testing.T) {
	p := newTestAlibabaCloudProvider(true)
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

func TestAlibabaCloudProvider_Records(t *testing.T) {
	p := newTestAlibabaCloudProvider(false)
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

func TestAlibabaCloudProvider_ApplyChanges(t *testing.T) {
	p := newTestAlibabaCloudProvider(false)
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "xyz.container-service.top",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "abc.container-service.top",
				RecordType: "A",
				RecordTTL:  500,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "abc.container-service.top",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	ctx := context.Background()
	p.ApplyChanges(ctx, &changes)
	endpoints, err := p.Records(ctx)
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

func TestAlibabaCloudProvider_Records_PrivateZone(t *testing.T) {
	p := newTestAlibabaCloudProvider(true)
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

func TestAlibabaCloudProvider_ApplyChanges_PrivateZone(t *testing.T) {
	p := newTestAlibabaCloudProvider(true)
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "xyz.container-service.top",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "abc.container-service.top",
				RecordType: "A",
				RecordTTL:  500,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "abc.container-service.top",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	ctx := context.Background()
	p.ApplyChanges(ctx, &changes)
	endpoints, err := p.Records(ctx)
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

func TestAlibabaCloudProvider_splitDNSName(t *testing.T) {
	p := newTestAlibabaCloudProvider(false)
	endpoint := &endpoint.Endpoint{}
	endpoint.DNSName = "www.example.org"
	rr, domain := p.splitDNSName(endpoint)
	if rr != "www" || domain != "example.org" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = ".example.org"
	rr, domain = p.splitDNSName(endpoint)
	if rr != "@" || domain != "example.org" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "www"
	rr, domain = p.splitDNSName(endpoint)
	if rr != "www" || domain != "" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = ""
	rr, domain = p.splitDNSName(endpoint)
	if rr != "@" || domain != "" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "_30000._tcp.container-service.top"
	rr, domain = p.splitDNSName(endpoint)
	if rr != "_30000._tcp" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "container-service.top"
	rr, domain = p.splitDNSName(endpoint)
	if rr != "@" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "a.b.container-service.top"
	rr, domain = p.splitDNSName(endpoint)
	if rr != "a.b" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "a.b.c.container-service.top"
	rr, domain = p.splitDNSName(endpoint)
	if rr != "a.b.c" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
}

func TestAlibabaCloudProvider_TXTEndpoint(t *testing.T) {

	p := newTestAlibabaCloudProvider(false)
	const recordValue = "heritage=external-dns,external-dns/owner=default"
	const endpointTarget = "\"heritage=external-dns,external-dns/owner=default\""

	if p.escapeTXTRecordValue(endpointTarget) != endpointTarget {
		t.Errorf("Failed to escapeTXTRecordValue: %s", p.escapeTXTRecordValue(endpointTarget))
	}
	if p.unescapeTXTRecordValue(recordValue) != endpointTarget {
		t.Errorf("Failed to unescapeTXTRecordValue: %s", p.unescapeTXTRecordValue(recordValue))
	}
}

//TestAlibabaCloudProvider_TXTEndpoint_PrivateZone
func TestAlibabaCloudProvider_TXTEndpoint_PrivateZone(t *testing.T) {

	p := newTestAlibabaCloudProvider(true)
	const recordValue = "heritage=external-dns,external-dns/owner=default"
	const endpointTarget = "\"heritage=external-dns,external-dns/owner=default\""

	if p.escapeTXTRecordValue(endpointTarget) != endpointTarget {
		t.Errorf("Failed to escapeTXTRecordValue: %s", p.escapeTXTRecordValue(endpointTarget))
	}
	if p.unescapeTXTRecordValue(recordValue) != endpointTarget {
		t.Errorf("Failed to unescapeTXTRecordValue: %s", p.unescapeTXTRecordValue(recordValue))
	}
}
