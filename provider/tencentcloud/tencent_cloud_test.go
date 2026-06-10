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

package tencentcloud

import (
	"context"
	strRand "crypto/rand"
	"fmt"
	numRand "math/rand"
	"strconv"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type MockDNSPodAPI struct {
	regionId string
	domains  []*dnspod.DomainListItem
	records  map[string][]*dnspod.RecordListItem
}

type MockPrivateDNSAPI struct {
	regionId string
	zones    []*privatedns.PrivateZone
	records  map[string][]*privatedns.PrivateZoneRecord
}

func (api *MockDNSPodAPI) CreateRecord(request *dnspod.CreateRecordRequest) (*dnspod.CreateRecordResponse, error) {
	if request == nil || request.DomainId == nil {
		return dnspod.NewCreateRecordResponse(), fmt.Errorf("invalid request")
	}

	domainId := strconv.FormatUint(*request.DomainId, 10)
	if _, exist := api.records[domainId]; !exist {
		return dnspod.NewCreateRecordResponse(), fmt.Errorf("domain not found")
	}
	ttl := request.TTL
	if ttl == nil {
		ttl = common.Uint64Ptr(defaultDnsTTL)
	}
	api.records[domainId] = append(api.records[domainId], &dnspod.RecordListItem{
		RecordId: common.Uint64Ptr(numRand.Uint64()),
		Value:    request.Value,
		Name:     request.SubDomain,
		Line:     request.RecordLine,
		LineId:   request.RecordLineId,
		Type:     request.RecordType,
		TTL:      ttl,
	})

	response := dnspod.NewCreateRecordResponse()
	response.Response = &dnspod.CreateRecordResponseParams{}
	return response, nil
}

func (api *MockDNSPodAPI) CreateRecordBatch(request *dnspod.CreateRecordBatchRequest) (*dnspod.CreateRecordBatchResponse, error) {
	for _, domainID := range request.DomainIdList {
		if _, exist := api.records[*domainID]; !exist {
			continue
		}
		for _, record := range request.RecordList {
			ttl := record.TTL
			if ttl == nil {
				ttl = common.Uint64Ptr(defaultDnsTTL)
			}
			api.records[*domainID] = append(api.records[*domainID], &dnspod.RecordListItem{
				RecordId: common.Uint64Ptr(numRand.Uint64()),
				Value:    record.Value,
				Name:     record.SubDomain,
				Line:     record.RecordLine,
				LineId:   record.RecordLineId,
				Type:     record.RecordType,
				TTL:      ttl,
			})
		}
	}

	response := dnspod.NewCreateRecordBatchResponse()
	response.Response = &dnspod.CreateRecordBatchResponseParams{}
	return response, nil
}

func (api *MockDNSPodAPI) DeleteRecord(request *dnspod.DeleteRecordRequest) (*dnspod.DeleteRecordResponse, error) {
	if request == nil || request.DomainId == nil || request.RecordId == nil {
		return dnspod.NewDeleteRecordResponse(), fmt.Errorf("invalid request")
	}

	domainId := strconv.FormatUint(*request.DomainId, 10)

	if _, exist := api.records[domainId]; !exist {
		return dnspod.NewDeleteRecordResponse(), fmt.Errorf("domain not found")
	}

	recordId := *request.RecordId
	records := api.records[domainId]
	remaining := make([]*dnspod.RecordListItem, 0, len(records))
	for _, record := range records {
		if *record.RecordId != recordId {
			remaining = append(remaining, record)
		}
	}
	api.records[domainId] = remaining

	response := dnspod.NewDeleteRecordResponse()
	response.Response = &dnspod.DeleteRecordResponseParams{}
	return response, nil
}

func (api *MockDNSPodAPI) DeleteRecordBatch(request *dnspod.DeleteRecordBatchRequest) (*dnspod.DeleteRecordBatchResponse, error) {
	recordIDs := make(map[uint64]struct{}, len(request.RecordIdList))
	for _, recordID := range request.RecordIdList {
		recordIDs[*recordID] = struct{}{}
	}

	for domain, records := range api.records {
		remaining := make([]*dnspod.RecordListItem, 0, len(records))
		for _, record := range records {
			if _, deleteRecord := recordIDs[*record.RecordId]; !deleteRecord {
				remaining = append(remaining, record)
			}
		}
		api.records[domain] = remaining
	}

	response := dnspod.NewDeleteRecordBatchResponse()
	response.Response = &dnspod.DeleteRecordBatchResponseParams{}
	return response, nil
}

func (api *MockDNSPodAPI) ModifyRecord(request *dnspod.ModifyRecordRequest) (*dnspod.ModifyRecordResponse, error) {
	if request == nil || request.DomainId == nil || request.RecordId == nil {
		return dnspod.NewModifyRecordResponse(), fmt.Errorf("invalid request")
	}

	domainId := strconv.FormatUint(*request.DomainId, 10)

	if _, exist := api.records[domainId]; !exist {
		return dnspod.NewModifyRecordResponse(), fmt.Errorf("domain not found")
	}

	recordId := *request.RecordId
	records := api.records[domainId]
	for _, record := range records {
		if *record.RecordId == recordId {
			if request.Value != nil {
				record.Value = request.Value
			}
			if request.SubDomain != nil {
				record.Name = request.SubDomain
			}
			if request.RecordLine != nil {
				record.Line = request.RecordLine
			}
			if request.RecordLineId != nil {
				record.LineId = request.RecordLineId
			}
			if request.RecordType != nil {
				record.Type = request.RecordType
			}
			if request.TTL != nil {
				record.TTL = request.TTL
			}
			break
		}
	}

	response := dnspod.NewModifyRecordResponse()
	response.Response = &dnspod.ModifyRecordResponseParams{}
	return response, nil
}

func (api *MockDNSPodAPI) ModifyRecordBatchV3(request *dnspod.ModifyRecordBatchV3Request) (*dnspod.ModifyRecordBatchV3Response, error) {
	modifyRecords := make(map[uint64]*dnspod.ModifyRecordItem, len(request.ModifyRecordList))
	for _, record := range request.ModifyRecordList {
		if record.RecordId != nil {
			modifyRecords[*record.RecordId] = record
		}
	}

	for _, records := range api.records {
		for _, record := range records {
			if modifyRecord, exists := modifyRecords[*record.RecordId]; exists {
				if modifyRecord.Value != nil {
					record.Value = modifyRecord.Value
				}
				if modifyRecord.SubDomain != nil {
					record.Name = modifyRecord.SubDomain
				}
				if modifyRecord.RecordLine != nil {
					record.Line = modifyRecord.RecordLine
				}
				if modifyRecord.RecordType != nil {
					record.Type = modifyRecord.RecordType
				}
				if modifyRecord.TTL != nil {
					record.TTL = modifyRecord.TTL
				}
			}
		}
	}

	response := dnspod.NewModifyRecordBatchV3Response()
	response.Response = &dnspod.ModifyRecordBatchV3ResponseParams{}
	return response, nil
}

func (api *MockDNSPodAPI) DescribeDomainList(request *dnspod.DescribeDomainListRequest) (*dnspod.DescribeDomainListResponse, error) {
	response := dnspod.NewDescribeDomainListResponse()
	response.Response = &dnspod.DescribeDomainListResponseParams{
		DomainCountInfo: &dnspod.DomainCountInfo{
			AllTotal: common.Uint64Ptr(uint64(len(api.domains))),
		},
		DomainList: api.domains,
	}
	return response, nil
}

func (api *MockDNSPodAPI) DescribeRecordList(request *dnspod.DescribeRecordListRequest) (*dnspod.DescribeRecordListResponse, error) {
	domainId := ""
	if request.Domain != nil {
		for _, domain := range api.domains {
			if domain.Name != nil && *domain.Name == *request.Domain {
				domainId = strconv.FormatUint(*domain.DomainId, 10)
				break
			}
		}
	} else if request.DomainId != nil {
		domainId = strconv.FormatUint(*request.DomainId, 10)
	}

	if domainId == "" {
		return dnspod.NewDescribeRecordListResponse(), nil
	}

	records := api.records[domainId]
	response := dnspod.NewDescribeRecordListResponse()
	response.Response = &dnspod.DescribeRecordListResponseParams{
		RecordCountInfo: &dnspod.RecordCountInfo{
			TotalCount: common.Uint64Ptr(uint64(len(records))),
		},
		RecordList: records,
	}
	return response, nil
}

func (api *MockPrivateDNSAPI) CreatePrivateZoneRecord(request *privatedns.CreatePrivateZoneRecordRequest) (*privatedns.CreatePrivateZoneRecordResponse, error) {
	if request.ZoneId == nil {
		return privatedns.NewCreatePrivateZoneRecordResponse(), fmt.Errorf("invalid request")
	}
	zoneID := *request.ZoneId
	if _, exist := api.records[zoneID]; !exist {
		api.records[zoneID] = make([]*privatedns.PrivateZoneRecord, 0)
	}
	ttl := request.TTL
	if ttl == nil {
		ttl = common.Int64Ptr(defaultPvtTTL)
	}
	api.records[zoneID] = append(api.records[zoneID], &privatedns.PrivateZoneRecord{
		RecordId:    common.StringPtr(strRand.Text()),
		ZoneId:      &zoneID,
		SubDomain:   request.SubDomain,
		RecordType:  request.RecordType,
		RecordValue: request.RecordValue,
		TTL:         ttl,
		MX:          request.MX,
		Weight:      request.Weight,
	})
	return privatedns.NewCreatePrivateZoneRecordResponse(), nil
}

func (api *MockPrivateDNSAPI) CreatePrivateZoneRecordList(request *privatedns.CreatePrivateZoneRecordListRequest) (*privatedns.CreatePrivateZoneRecordListResponse, error) {
	for _, zoneID := range request.ZoneIds {
		if _, exist := api.records[*zoneID]; !exist {
			api.records[*zoneID] = make([]*privatedns.PrivateZoneRecord, 0)
		}
		for _, record := range request.RecordsInfo {
			ttl := record.TTL
			if ttl == nil {
				ttl = common.Int64Ptr(defaultPvtTTL)
			}
			api.records[*zoneID] = append(api.records[*zoneID], &privatedns.PrivateZoneRecord{
				RecordId:    common.StringPtr(strRand.Text()),
				ZoneId:      zoneID,
				SubDomain:   record.SubDomain,
				RecordType:  record.RecordType,
				RecordValue: record.RecordValue,
				TTL:         ttl,
				MX:          record.MX,
				Weight:      record.Weight,
			})
		}
	}

	response := privatedns.NewCreatePrivateZoneRecordListResponse()
	response.Response = &privatedns.CreatePrivateZoneRecordListResponseParams{}
	return response, nil
}

func (api *MockPrivateDNSAPI) DeletePrivateZoneRecord(request *privatedns.DeletePrivateZoneRecordRequest) (*privatedns.DeletePrivateZoneRecordResponse, error) {
	if request.ZoneId == nil {
		return privatedns.NewDeletePrivateZoneRecordResponse(), nil
	}

	recordIDs := make(map[string]struct{}, len(request.RecordIdSet)+1)
	if request.RecordId != nil {
		recordIDs[*request.RecordId] = struct{}{}
	}
	for _, recordID := range request.RecordIdSet {
		recordIDs[*recordID] = struct{}{}
	}

	records := api.records[*request.ZoneId]
	remaining := make([]*privatedns.PrivateZoneRecord, 0, len(records))
	for _, record := range records {
		if _, deleteRecord := recordIDs[*record.RecordId]; !deleteRecord {
			remaining = append(remaining, record)
		}
	}
	api.records[*request.ZoneId] = remaining

	response := privatedns.NewDeletePrivateZoneRecordResponse()
	response.Response = &privatedns.DeletePrivateZoneRecordResponseParams{}
	return response, nil
}

func (api *MockPrivateDNSAPI) ModifyPrivateZoneRecord(request *privatedns.ModifyPrivateZoneRecordRequest) (*privatedns.ModifyPrivateZoneRecordResponse, error) {
	if request == nil || request.ZoneId == nil || request.RecordId == nil {
		return privatedns.NewModifyPrivateZoneRecordResponse(), fmt.Errorf("invalid request")
	}
	records := api.records[*request.ZoneId]

	for _, record := range records {
		if *record.RecordId == *request.RecordId {
			if request.SubDomain != nil {
				record.SubDomain = request.SubDomain
			}

			if request.RecordType != nil {
				record.RecordType = request.RecordType
			}

			if request.RecordValue != nil {
				record.RecordValue = request.RecordValue
			}

			if request.TTL != nil {
				record.TTL = request.TTL
			}

			if request.MX != nil {
				record.MX = request.MX
			}

			if request.Weight != nil {
				record.Weight = request.Weight
			}
		}
	}
	return privatedns.NewModifyPrivateZoneRecordResponse(), nil
}

func (api *MockPrivateDNSAPI) DescribePrivateZoneList(request *privatedns.DescribePrivateZoneListRequest) (*privatedns.DescribePrivateZoneListResponse, error) {
	response := privatedns.NewDescribePrivateZoneListResponse()
	response.Response = &privatedns.DescribePrivateZoneListResponseParams{
		TotalCount:     common.Int64Ptr(int64(len(api.zones))),
		PrivateZoneSet: api.zones,
	}
	return response, nil
}

func (api *MockPrivateDNSAPI) DescribePrivateZoneRecordList(request *privatedns.DescribePrivateZoneRecordListRequest) (*privatedns.DescribePrivateZoneRecordListResponse, error) {
	records := api.records[*request.ZoneId]
	response := privatedns.NewDescribePrivateZoneRecordListResponse()
	response.Response = &privatedns.DescribePrivateZoneRecordListResponseParams{
		TotalCount: common.Int64Ptr(int64(len(records))),
		RecordSet:  records,
	}
	return response, nil
}

func newMockTencentCloudProvider(domainFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, privateZone bool) *TencentCloudProvider {
	cfg := tencentCloudConfig{
		RegionId: "ap-shanghai",
		VPCId:    "vpc-abcdefg",
	}

	zoneId1 := common.StringPtr(strRand.Text())
	privateZones := []*privatedns.PrivateZone{
		{
			ZoneId: zoneId1,
			Domain: common.StringPtr("external-dns-test.com"),
			VpcSet: []*privatedns.VpcInfo{
				{
					UniqVpcId: common.StringPtr("vpc-abcdefg"),
					Region:    common.StringPtr("ap-shanghai"),
				},
			},
		},
	}

	zoneRecordId1 := common.StringPtr(strRand.Text())
	zoneRecordId2 := common.StringPtr(strRand.Text())
	privateZoneRecords := map[string][]*privatedns.PrivateZoneRecord{
		*zoneId1: {
			{
				ZoneId:      zoneId1,
				RecordId:    zoneRecordId1,
				SubDomain:   common.StringPtr("nginx"),
				RecordType:  common.StringPtr("TXT"),
				RecordValue: common.StringPtr("heritage=external-dns,external-dns/owner=default"),
				TTL:         common.Int64Ptr(300),
			},
			{
				ZoneId:      zoneId1,
				RecordId:    zoneRecordId2,
				SubDomain:   common.StringPtr("nginx"),
				RecordType:  common.StringPtr("A"),
				RecordValue: common.StringPtr("10.10.10.10"),
				TTL:         common.Int64Ptr(300),
			},
		},
	}

	dnsDomainId1 := common.Uint64Ptr(numRand.Uint64())
	dnsDomainId1Str := strconv.FormatUint(*dnsDomainId1, 10)
	dnsPodDomains := []*dnspod.DomainListItem{
		{
			DomainId: dnsDomainId1,
			Name:     common.StringPtr("external-dns-test.com"),
		},
	}

	dnsDomainRecordId1 := common.Uint64Ptr(numRand.Uint64())
	dnsDomainRecordId2 := common.Uint64Ptr(numRand.Uint64())
	dnspodRecords := map[string][]*dnspod.RecordListItem{
		dnsDomainId1Str: {
			{
				RecordId: dnsDomainRecordId1,
				Value:    common.StringPtr("heritage=external-dns,external-dns/owner=default"),
				Name:     common.StringPtr("nginx"),
				Type:     common.StringPtr("TXT"),
				TTL:      common.Uint64Ptr(300),
			},
			{
				RecordId: dnsDomainRecordId2,
				Name:     common.StringPtr("nginx"),
				Type:     common.StringPtr("A"),
				Value:    common.StringPtr("10.10.10.10"),
				TTL:      common.Uint64Ptr(300),
			},
		},
	}

	dnsApi := &MockDNSPodAPI{
		regionId: cfg.RegionId,
		domains:  dnsPodDomains,
		records:  dnspodRecords,
	}

	zoneApi := &MockPrivateDNSAPI{
		regionId: cfg.RegionId,
		zones:    privateZones,
		records:  privateZoneRecords,
	}

	tencentCloudProvider := &TencentCloudProvider{
		domainFilter: domainFilter,
		zoneIDFilter: &zoneIDFilter,
		dryRun:       false,
		dnsApi:       dnsApi,
		pvtApi:       zoneApi,
		vpcID:        cfg.VPCId,
		privateZone:  privateZone,
	}

	return tencentCloudProvider
}

func TestTencentCloudProvider_PrivateDNS_Records(t *testing.T) {
	p := newMockTencentCloudProvider(endpoint.NewDomainFilter([]string{"external-dns-test.com"}), provider.NewZoneIDFilter([]string{}), true)
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}
}

func TestTencentCloudProvider_PrivateDNS_ApplyChanges(t *testing.T) {
	p := newMockTencentCloudProvider(endpoint.NewDomainFilter([]string{"external-dns-test.com"}), provider.NewZoneIDFilter([]string{}), true)

	// Test for Create、UpdateOld、UpdateNew、Delete
	// The base record will be created.
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "redis.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("10.10.10.10"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete one target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete another target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Create new records
	changes = &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 4 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete new records
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}
}

func TestTencentCloudProvider_PublicDNS_Records(t *testing.T) {
	p := newMockTencentCloudProvider(endpoint.NewDomainFilter([]string{"external-dns-test.com"}), provider.NewZoneIDFilter([]string{}), false)
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}
}

func TestTencentCloudProvider_PublicDNS_ApplyChanges(t *testing.T) {
	p := newMockTencentCloudProvider(endpoint.NewDomainFilter([]string{"external-dns-test.com"}), provider.NewZoneIDFilter([]string{}), false)
	// Test for Create、UpdateOld、UpdateNew、Delete
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "redis.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("10.10.10.10"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete one target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete another target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 1 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Create new records
	changes = &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete new records
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
			{
				DNSName:    "new.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 1 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}
}
