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
	"fmt"
	"strconv"
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// DnsPod For Public Dns

func (p *TencentCloudProvider) dnsRecords() ([]*endpoint.Endpoint, error) {
	recordsList, err := p.recordsForDNS()
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)
	recordMap := groupDomainRecordList(recordsList)
	for _, recordList := range recordMap {
		name := getDnsDomain(*recordList.RecordList[0].Name, *recordList.Domain.Name)
		recordType := *recordList.RecordList[0].Type
		ttl := *recordList.RecordList[0].TTL
		var targets []string
		for _, record := range recordList.RecordList {
			targets = append(targets, *record.Value)
		}
		endpoints = append(endpoints, endpoint.NewEndpointWithTTL(name, recordType, endpoint.TTL(ttl), targets...))
	}
	return endpoints, nil
}

func (p *TencentCloudProvider) recordsForDNS() (map[uint64]*RecordListGroup, error) {
	domainList, err := p.getDomainList()
	if err != nil {
		return nil, err
	}

	recordListGroup := make(map[uint64]*RecordListGroup, 0)
	for _, domain := range domainList {
		records, err := p.getDomainRecordList(*domain.Name)
		if err != nil {
			return nil, err
		}
		for _, record := range records {
			if *record.Type == "TXT" && strings.HasPrefix(*record.Value, "heritage=") {
				record.Value = common.StringPtr(fmt.Sprintf(`"%s"`, *record.Value))
			}
		}
		recordListGroup[*domain.DomainId] = &RecordListGroup{
			Domain:     domain,
			RecordList: records,
		}
	}
	return recordListGroup, nil
}

func (p *TencentCloudProvider) getDomainList() ([]*dnspod.DomainListItem, error) {
	request := dnspod.NewDescribeDomainListRequest()
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(3000)

	domainList := make([]*dnspod.DomainListItem, 0)
	totalCount := int64(100)
	for *request.Offset < totalCount {
		response, err := p.apiService.DescribeDomainList(request)
		if err != nil {
			return nil, err
		}
		if response.Response.DomainList != nil && len(response.Response.DomainList) > 0 {
			if !p.domainFilter.IsConfigured() {
				domainList = append(domainList, response.Response.DomainList...)
			} else {
				for _, domain := range response.Response.DomainList {
					if p.domainFilter.Match(*domain.Name) {
						domainList = append(domainList, domain)
					}
				}
			}
		}
		totalCount = int64(*response.Response.DomainCountInfo.AllTotal)
		request.Offset = common.Int64Ptr(*request.Offset + int64(len(response.Response.DomainList)))
	}
	return domainList, nil
}

func (p *TencentCloudProvider) getDomainRecordList(domain string) ([]*dnspod.RecordListItem, error) {
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(domain)
	request.Offset = common.Uint64Ptr(0)
	request.Limit = common.Uint64Ptr(3000)

	domainList := make([]*dnspod.RecordListItem, 0)
	totalCount := uint64(100)
	for *request.Offset < totalCount {
		response, err := p.apiService.DescribeRecordList(request)
		if err != nil {
			return nil, err
		}
		if response.Response.RecordList != nil && len(response.Response.RecordList) > 0 {
			for _, record := range response.Response.RecordList {
				if *record.Name == "@" && *record.Type == "NS" { // Special Record, Skip it.
					continue
				}
				domainList = append(domainList, record)
			}
		}
		totalCount = *response.Response.RecordCountInfo.TotalCount
		request.Offset = common.Uint64Ptr(*request.Offset + uint64(len(response.Response.RecordList)))
	}
	return domainList, nil
}

type RecordListGroup struct {
	Domain     *dnspod.DomainListItem
	RecordList []*dnspod.RecordListItem
}

func (p *TencentCloudProvider) applyChangesForDNS(changes *plan.Changes) error {
	recordsGroupMap, err := p.recordsForDNS()
	if err != nil {
		return err
	}

	zoneNameIDMapper := provider.ZoneIDName{}
	for _, recordsGroup := range recordsGroupMap {
		if recordsGroup.Domain.DomainId != nil {
			zoneNameIDMapper.Add(strconv.FormatUint(*recordsGroup.Domain.DomainId, 10), *recordsGroup.Domain.Name)
		}
	}

	// Apply Change Delete
	deleteEndpoints := make(map[string][]uint64)
	for _, change := range [][]*endpoint.Endpoint{changes.Delete, changes.UpdateOld} {
		for _, deleteChange := range change {
			if zoneId, _ := zoneNameIDMapper.FindZone(deleteChange.DNSName); zoneId != "" {
				zoneIdString, _ := strconv.ParseUint(zoneId, 10, 64)
				recordListGroup := recordsGroupMap[zoneIdString]
				for _, domainRecord := range recordListGroup.RecordList {
					subDomain := getSubDomain(*recordListGroup.Domain.Name, deleteChange)
					if *domainRecord.Name == subDomain && *domainRecord.Type == deleteChange.RecordType {
						for _, target := range deleteChange.Targets {
							if *domainRecord.Value == target {
								if _, exist := deleteEndpoints[*recordListGroup.Domain.Name]; !exist {
									deleteEndpoints[*recordListGroup.Domain.Name] = make([]uint64, 0)
								}
								deleteEndpoints[*recordListGroup.Domain.Name] = append(deleteEndpoints[*recordListGroup.Domain.Name], *domainRecord.RecordId)
							}
						}
					}
				}
			}
		}
	}

	if err := p.deleteRecords(deleteEndpoints); err != nil {
		return err
	}

	// Apply Change Create
	createEndpoints := make(map[string][]*endpoint.Endpoint)
	for zoneId := range zoneNameIDMapper {
		createEndpoints[zoneId] = make([]*endpoint.Endpoint, 0)
	}
	for _, change := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew} {
		for _, createChange := range change {
			if zoneId, _ := zoneNameIDMapper.FindZone(createChange.DNSName); zoneId != "" {
				createEndpoints[zoneId] = append(createEndpoints[zoneId], createChange)
			}
		}
	}
	if err := p.createRecord(recordsGroupMap, createEndpoints); err != nil {
		return err
	}
	return nil
}

func (p *TencentCloudProvider) createRecord(zoneMap map[uint64]*RecordListGroup, endpointsMap map[string][]*endpoint.Endpoint) error {
	for zoneId, endpoints := range endpointsMap {
		zoneIdString, _ := strconv.ParseUint(zoneId, 10, 64)
		domain := zoneMap[zoneIdString]
		for _, endpoint := range endpoints {
			for _, target := range endpoint.Targets {
				if endpoint.RecordType == "TXT" && strings.HasPrefix(target, `"heritage=`) {
					target = strings.Trim(target, `"`)
				}
				if err := p.createRecords(domain.Domain, endpoint, target); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *TencentCloudProvider) createRecords(domain *dnspod.DomainListItem, endpoint *endpoint.Endpoint, target string) error {
	request := dnspod.NewCreateRecordRequest()

	request.Domain = common.StringPtr(*domain.Name)
	request.RecordType = common.StringPtr(endpoint.RecordType)
	request.Value = common.StringPtr(target)
	request.SubDomain = common.StringPtr(getSubDomain(*domain.Name, endpoint))
	if endpoint.RecordTTL.IsConfigured() {
		request.TTL = common.Uint64Ptr(uint64(endpoint.RecordTTL))
	}
	request.RecordLine = common.StringPtr("默认")

	if _, err := p.apiService.CreateRecord(request); err != nil {
		return err
	}
	return nil
}

func (p *TencentCloudProvider) deleteRecords(RecordIdsMap map[string][]uint64) error {
	for domain, recordIds := range RecordIdsMap {
		if len(recordIds) == 0 {
			continue
		}
		if err := p.deleteRecord(domain, recordIds); err != nil {
			return err
		}
	}
	return nil
}

func (p *TencentCloudProvider) deleteRecord(domain string, recordIds []uint64) error {
	request := dnspod.NewDeleteRecordRequest()
	request.Domain = common.StringPtr(domain)

	for _, recordId := range recordIds {
		request.RecordId = common.Uint64Ptr(recordId)
		if _, err := p.apiService.DeleteRecord(request); err != nil {
			return err
		}
	}
	return nil
}

func groupDomainRecordList(recordListGroup map[uint64]*RecordListGroup) (endpointMap map[string]*RecordListGroup) {
	endpointMap = make(map[string]*RecordListGroup)

	for _, recordGroup := range recordListGroup {
		for _, record := range recordGroup.RecordList {
			key := fmt.Sprintf("%s:%s.%s", *record.Type, *record.Name, *recordGroup.Domain.Name)
			if *record.Name == TencentCloudEmptyPrefix {
				key = fmt.Sprintf("%s:%s", *record.Type, *recordGroup.Domain.Name)
			}
			if _, exist := endpointMap[key]; !exist {
				endpointMap[key] = &RecordListGroup{
					Domain:     recordGroup.Domain,
					RecordList: make([]*dnspod.RecordListItem, 0),
				}
			}
			endpointMap[key].RecordList = append(endpointMap[key].RecordList, record)
		}
	}

	return endpointMap
}
