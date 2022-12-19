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
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// PrivateZone For Internal Dns

func (p *TencentCloudProvider) privateZoneRecords() ([]*endpoint.Endpoint, error) {
	privateZones, err := p.recordForPrivateZone()
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)
	recordMap := groupPrivateZoneRecords(privateZones)
	for _, recordList := range recordMap {
		name := getDnsDomain(*recordList.RecordList[0].SubDomain, *recordList.Zone.Domain)
		recordType := *recordList.RecordList[0].RecordType
		ttl := *recordList.RecordList[0].TTL
		var targets []string
		for _, record := range recordList.RecordList {
			targets = append(targets, *record.RecordValue)
		}
		endpoints = append(endpoints, endpoint.NewEndpointWithTTL(name, recordType, endpoint.TTL(ttl), targets...))
	}
	return endpoints, nil
}

func (p *TencentCloudProvider) recordForPrivateZone() (map[string]*PrivateZoneRecordListGroup, error) {
	privateZones, err := p.getPrivateZones()
	if err != nil {
		return nil, err
	}

	recordListGroup := make(map[string]*PrivateZoneRecordListGroup, 0)
	for _, zone := range privateZones {
		records, err := p.getPrivateZoneRecords(*zone.ZoneId)
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			if *record.RecordType == "TXT" && strings.HasPrefix(*record.RecordValue, "heritage=") {
				record.RecordValue = common.StringPtr(fmt.Sprintf("\"%s\"", *record.RecordValue))
			}
		}
		recordListGroup[*zone.ZoneId] = &PrivateZoneRecordListGroup{
			Zone:       zone,
			RecordList: records,
		}
	}

	return recordListGroup, nil
}

func (p *TencentCloudProvider) getPrivateZones() ([]*privatedns.PrivateZone, error) {
	filters := make([]*privatedns.Filter, 1)
	filters[0] = &privatedns.Filter{
		Name: common.StringPtr("Vpc"),
		Values: []*string{
			common.StringPtr(p.vpcID),
		},
	}

	if p.zoneIDFilter.IsConfigured() {
		zoneIDs := make([]*string, len(p.zoneIDFilter.ZoneIDs))
		for index, zoneId := range p.zoneIDFilter.ZoneIDs {
			zoneIDs[index] = common.StringPtr(zoneId)
		}
		filters = append(filters, &privatedns.Filter{
			Name:   common.StringPtr("ZoneId"),
			Values: zoneIDs,
		})
	}

	request := privatedns.NewDescribePrivateZoneListRequest()
	request.Filters = filters
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(100)

	privateZones := make([]*privatedns.PrivateZone, 0)
	totalCount := int64(100)
	for *request.Offset < totalCount {
		response, err := p.apiService.DescribePrivateZoneList(request)
		if err != nil {
			return nil, err
		}
		if response.Response.PrivateZoneSet != nil && len(response.Response.PrivateZoneSet) > 0 {
			privateZones = append(privateZones, response.Response.PrivateZoneSet...)
		}
		totalCount = *response.Response.TotalCount
		request.Offset = common.Int64Ptr(*request.Offset + int64(len(response.Response.PrivateZoneSet)))
	}

	privateZonesFilter := make([]*privatedns.PrivateZone, 0)
	for _, privateZone := range privateZones {
		if p.domainFilter.IsConfigured() && !p.domainFilter.Match(*privateZone.Domain) {
			continue
		}
		privateZonesFilter = append(privateZonesFilter, privateZone)
	}
	return privateZonesFilter, nil
}

func (p *TencentCloudProvider) getPrivateZoneRecords(zoneId string) ([]*privatedns.PrivateZoneRecord, error) {
	request := privatedns.NewDescribePrivateZoneRecordListRequest()
	request.ZoneId = common.StringPtr(zoneId)
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(100)

	privateZoneRecords := make([]*privatedns.PrivateZoneRecord, 0)
	totalCount := int64(100)
	for *request.Offset < totalCount {
		response, err := p.apiService.DescribePrivateZoneRecordList(request)
		if err != nil {
			return nil, err
		}
		if response.Response.RecordSet != nil && len(response.Response.RecordSet) > 0 {
			privateZoneRecords = append(privateZoneRecords, response.Response.RecordSet...)
		}
		totalCount = *response.Response.TotalCount
		request.Offset = common.Int64Ptr(*request.Offset + int64(len(response.Response.RecordSet)))
	}
	return privateZoneRecords, nil
}

type PrivateZoneRecordListGroup struct {
	Zone       *privatedns.PrivateZone
	RecordList []*privatedns.PrivateZoneRecord
}

// Returns nil if the operation was successful or an error if the operation failed.
func (p *TencentCloudProvider) applyChangesForPrivateZone(changes *plan.Changes) error {
	zoneGroups, err := p.recordForPrivateZone()
	if err != nil {
		return err
	}

	// In PrivateDns Service. A Zone has at least one record. The last rule cannot be deleted.
	for _, zoneGroup := range zoneGroups {
		if !containsBaseRecord(zoneGroup.RecordList) {
			err := p.createPrivateZoneRecord(zoneGroup.Zone, &endpoint.Endpoint{
				DNSName:    *zoneGroup.Zone.Domain,
				RecordType: "TXT",
			}, "tencent_provider_record")
			if err != nil {
				return err
			}
		}
	}

	zoneNameIDMapper := provider.ZoneIDName{}
	for _, zoneGroup := range zoneGroups {
		if zoneGroup.Zone.ZoneId != nil {
			zoneNameIDMapper.Add(*zoneGroup.Zone.ZoneId, *zoneGroup.Zone.Domain)
		}
	}

	// Apply Change Delete
	deleteEndpoints := make(map[string][]string)
	for _, change := range [][]*endpoint.Endpoint{changes.Delete, changes.UpdateOld} {
		for _, deleteChange := range change {
			if zoneId, _ := zoneNameIDMapper.FindZone(deleteChange.DNSName); zoneId != "" {
				zoneGroup := zoneGroups[zoneId]
				for _, zoneRecord := range zoneGroup.RecordList {
					subDomain := getSubDomain(*zoneGroup.Zone.Domain, deleteChange)
					if *zoneRecord.SubDomain == subDomain && *zoneRecord.RecordType == deleteChange.RecordType {
						for _, target := range deleteChange.Targets {
							if *zoneRecord.RecordValue == target {
								if _, exist := deleteEndpoints[zoneId]; !exist {
									deleteEndpoints[zoneId] = make([]string, 0)
								}
								deleteEndpoints[zoneId] = append(deleteEndpoints[zoneId], *zoneRecord.RecordId)
							}
						}
					}
				}
			}
		}
	}

	if err := p.deletePrivateZoneRecords(deleteEndpoints); err != nil {
		return err
	}

	// Apply Change Create
	createEndpoints := make(map[string][]*endpoint.Endpoint)
	for _, change := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew} {
		for _, createChange := range change {
			if zoneId, _ := zoneNameIDMapper.FindZone(createChange.DNSName); zoneId != "" {
				if _, exist := createEndpoints[zoneId]; !exist {
					createEndpoints[zoneId] = make([]*endpoint.Endpoint, 0)
				}
				createEndpoints[zoneId] = append(createEndpoints[zoneId], createChange)
			}
		}
	}
	if err := p.createPrivateZoneRecords(zoneGroups, createEndpoints); err != nil {
		return err
	}
	return nil
}

func containsBaseRecord(records []*privatedns.PrivateZoneRecord) bool {
	for _, record := range records {
		if *record.SubDomain == TencentCloudEmptyPrefix && *record.RecordType == "TXT" && *record.RecordValue == "tencent_provider_record" {
			return true
		}
	}
	return false
}

func (p *TencentCloudProvider) createPrivateZoneRecords(zoneGroups map[string]*PrivateZoneRecordListGroup, endpointsMap map[string][]*endpoint.Endpoint) error {
	for zoneId, endpoints := range endpointsMap {
		zoneGroup := zoneGroups[zoneId]
		for _, endpoint := range endpoints {
			for _, target := range endpoint.Targets {
				if endpoint.RecordType == "TXT" && strings.HasPrefix(target, "\"heritage=") {
					target = strings.Trim(target, "\"")
				}
				if err := p.createPrivateZoneRecord(zoneGroup.Zone, endpoint, target); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *TencentCloudProvider) deletePrivateZoneRecords(zoneRecordIdsMap map[string][]string) error {
	for zoneId, zoneRecordIds := range zoneRecordIdsMap {
		if len(zoneRecordIds) == 0 {
			continue
		}
		if err := p.deletePrivateZoneRecord(zoneId, zoneRecordIds); err != nil {
			return err
		}
	}
	return nil
}

func (p *TencentCloudProvider) createPrivateZoneRecord(zone *privatedns.PrivateZone, endpoint *endpoint.Endpoint, target string) error {
	request := privatedns.NewCreatePrivateZoneRecordRequest()
	request.ZoneId = common.StringPtr(*zone.ZoneId)
	request.RecordType = common.StringPtr(endpoint.RecordType)
	request.RecordValue = common.StringPtr(target)
	request.SubDomain = common.StringPtr(getSubDomain(*zone.Domain, endpoint))
	if endpoint.RecordTTL.IsConfigured() {
		request.TTL = common.Int64Ptr(int64(endpoint.RecordTTL))
	}

	if _, err := p.apiService.CreatePrivateZoneRecord(request); err != nil {
		return err
	}
	return nil
}

func (p *TencentCloudProvider) deletePrivateZoneRecord(zoneId string, zoneRecordIds []string) error {
	recordIds := make([]*string, len(zoneRecordIds))
	for index, recordId := range zoneRecordIds {
		recordIds[index] = common.StringPtr(recordId)
	}

	request := privatedns.NewDeletePrivateZoneRecordRequest()
	request.ZoneId = common.StringPtr(zoneId)
	request.RecordIdSet = recordIds

	if _, err := p.apiService.DeletePrivateZoneRecord(request); err != nil {
		return err
	}
	return nil
}

func groupPrivateZoneRecords(zoneRecords map[string]*PrivateZoneRecordListGroup) (endpointMap map[string]*PrivateZoneRecordListGroup) {
	endpointMap = make(map[string]*PrivateZoneRecordListGroup)

	for _, recordGroup := range zoneRecords {
		for _, record := range recordGroup.RecordList {
			key := fmt.Sprintf("%s:%s.%s", *record.RecordType, *record.SubDomain, *recordGroup.Zone.Domain)
			if *record.SubDomain == TencentCloudEmptyPrefix {
				key = fmt.Sprintf("%s:%s", *record.RecordType, *recordGroup.Zone.Domain)
			}
			if _, exist := endpointMap[key]; !exist {
				endpointMap[key] = &PrivateZoneRecordListGroup{
					Zone:       recordGroup.Zone,
					RecordList: make([]*privatedns.PrivateZoneRecord, 0),
				}
			}
			endpointMap[key].RecordList = append(endpointMap[key].RecordList, record)
		}
	}

	return endpointMap
}
