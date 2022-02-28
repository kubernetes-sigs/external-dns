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

package cloudapi

import (
	"math/rand"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
)

type mockAPIService struct {
	privateZones       []*privatedns.PrivateZone
	privateZoneRecords map[string][]*privatedns.PrivateZoneRecord

	dnspodDomains []*dnspod.DomainListItem
	dnspodRecords map[string][]*dnspod.RecordListItem
}

func NewMockService(privateZones []*privatedns.PrivateZone, privateZoneRecords map[string][]*privatedns.PrivateZoneRecord, dnspodDomains []*dnspod.DomainListItem, dnspodRecords map[string][]*dnspod.RecordListItem) *mockAPIService {
	rand.Seed(time.Now().Unix())
	return &mockAPIService{
		privateZones:       privateZones,
		privateZoneRecords: privateZoneRecords,
		dnspodDomains:      dnspodDomains,
		dnspodRecords:      dnspodRecords,
	}
}

////////////////////////////////////////////////////////////////
// PrivateDns API
////////////////////////////////////////////////////////////////

func (api *mockAPIService) CreatePrivateZoneRecord(request *privatedns.CreatePrivateZoneRecordRequest) (response *privatedns.CreatePrivateZoneRecordResponse, err error) {
	randomRecordId := RandStringRunes(8)
	if _, exist := api.privateZoneRecords[*request.ZoneId]; !exist {
		api.privateZoneRecords[*request.ZoneId] = make([]*privatedns.PrivateZoneRecord, 0)
	}
	if request.TTL == nil {
		request.TTL = common.Int64Ptr(300)
	}
	api.privateZoneRecords[*request.ZoneId] = append(api.privateZoneRecords[*request.ZoneId], &privatedns.PrivateZoneRecord{
		RecordId:    common.StringPtr(randomRecordId),
		ZoneId:      request.ZoneId,
		SubDomain:   request.SubDomain,
		RecordType:  request.RecordType,
		RecordValue: request.RecordValue,
		TTL:         request.TTL,
	})
	return response, nil
}

func (api *mockAPIService) DeletePrivateZoneRecord(request *privatedns.DeletePrivateZoneRecordRequest) (response *privatedns.DeletePrivateZoneRecordResponse, err error) {
	result := make([]*privatedns.PrivateZoneRecord, 0)
	if _, exist := api.privateZoneRecords[*request.ZoneId]; !exist {
		return response, nil
	}
	for _, privateZoneRecord := range api.privateZoneRecords[*request.ZoneId] {
		deleteflag := false
		if request.RecordIdSet != nil && len(request.RecordIdSet) != 0 {
			for _, recordId := range request.RecordIdSet {
				if *privateZoneRecord.RecordId == *recordId {
					deleteflag = true
					break
				}
			}
		}
		if request.RecordId != nil && *request.RecordId == *privateZoneRecord.RecordId {
			deleteflag = true
		}
		if !deleteflag {
			result = append(result, privateZoneRecord)
		}
	}
	api.privateZoneRecords[*request.ZoneId] = result
	return response, nil
}

func (api *mockAPIService) ModifyPrivateZoneRecord(request *privatedns.ModifyPrivateZoneRecordRequest) (response *privatedns.ModifyPrivateZoneRecordResponse, err error) {
	if _, exist := api.privateZoneRecords[*request.ZoneId]; !exist {
		return response, nil
	}
	for _, privateZoneRecord := range api.privateZoneRecords[*request.ZoneId] {
		if *privateZoneRecord.RecordId != *request.RecordId {
			continue
		}
		privateZoneRecord.ZoneId = request.ZoneId
		privateZoneRecord.SubDomain = request.SubDomain
		privateZoneRecord.RecordType = request.RecordType
		privateZoneRecord.RecordValue = request.RecordValue
		privateZoneRecord.TTL = request.TTL
	}
	return response, nil
}

func (api *mockAPIService) DescribePrivateZoneList(request *privatedns.DescribePrivateZoneListRequest) (response *privatedns.DescribePrivateZoneListResponse, err error) {
	response = privatedns.NewDescribePrivateZoneListResponse()
	response.Response = &struct {
		TotalCount     *int64                    `json:"TotalCount,omitempty" name:"TotalCount"`
		PrivateZoneSet []*privatedns.PrivateZone `json:"PrivateZoneSet,omitempty" name:"PrivateZoneSet"`
		RequestId      *string                   `json:"RequestId,omitempty" name:"RequestId"`
	}{
		TotalCount:     common.Int64Ptr(int64(len(api.privateZones))),
		PrivateZoneSet: api.privateZones,
	}
	return response, nil
}

func (api *mockAPIService) DescribePrivateZoneRecordList(request *privatedns.DescribePrivateZoneRecordListRequest) (response *privatedns.DescribePrivateZoneRecordListResponse, err error) {
	response = privatedns.NewDescribePrivateZoneRecordListResponse()
	response.Response = &struct {
		TotalCount *int64                          `json:"TotalCount,omitempty" name:"TotalCount"`
		RecordSet  []*privatedns.PrivateZoneRecord `json:"RecordSet,omitempty" name:"RecordSet"`
		RequestId  *string                         `json:"RequestId,omitempty" name:"RequestId"`
	}{}
	if _, exist := api.privateZoneRecords[*request.ZoneId]; !exist {
		response.Response.TotalCount = common.Int64Ptr(0)
		response.Response.RecordSet = make([]*privatedns.PrivateZoneRecord, 0)
		return response, nil
	}
	response.Response.TotalCount = common.Int64Ptr(int64(len(api.privateZoneRecords[*request.ZoneId])))
	response.Response.RecordSet = api.privateZoneRecords[*request.ZoneId]
	return response, nil
}

////////////////////////////////////////////////////////////////
// DnsPod API
////////////////////////////////////////////////////////////////

func (api *mockAPIService) DescribeDomainList(request *dnspod.DescribeDomainListRequest) (response *dnspod.DescribeDomainListResponse, err error) {
	response = dnspod.NewDescribeDomainListResponse()
	response.Response = &struct {
		DomainCountInfo *dnspod.DomainCountInfo  `json:"DomainCountInfo,omitempty" name:"DomainCountInfo"`
		DomainList      []*dnspod.DomainListItem `json:"DomainList,omitempty" name:"DomainList"`
		RequestId       *string                  `json:"RequestId,omitempty" name:"RequestId"`
	}{}
	response.Response.DomainList = api.dnspodDomains
	response.Response.DomainCountInfo = &dnspod.DomainCountInfo{
		AllTotal: common.Uint64Ptr(uint64(len(api.dnspodDomains))),
	}
	return response, nil
}

func (api *mockAPIService) DescribeRecordList(request *dnspod.DescribeRecordListRequest) (response *dnspod.DescribeRecordListResponse, err error) {
	response = dnspod.NewDescribeRecordListResponse()
	response.Response = &struct {
		RecordCountInfo *dnspod.RecordCountInfo  `json:"RecordCountInfo,omitempty" name:"RecordCountInfo"`
		RecordList      []*dnspod.RecordListItem `json:"RecordList,omitempty" name:"RecordList"`
		RequestId       *string                  `json:"RequestId,omitempty" name:"RequestId"`
	}{}
	if _, exist := api.dnspodRecords[*request.Domain]; !exist {
		response.Response.RecordList = make([]*dnspod.RecordListItem, 0)
		response.Response.RecordCountInfo = &dnspod.RecordCountInfo{
			TotalCount: common.Uint64Ptr(uint64(0)),
		}
		return response, nil
	}
	response.Response.RecordList = api.dnspodRecords[*request.Domain]
	response.Response.RecordCountInfo = &dnspod.RecordCountInfo{
		TotalCount: common.Uint64Ptr(uint64(len(api.dnspodRecords[*request.Domain]))),
	}
	return response, nil
}

func (api *mockAPIService) CreateRecord(request *dnspod.CreateRecordRequest) (response *dnspod.CreateRecordResponse, err error) {
	randomRecordId := RandUint64()
	if _, exist := api.dnspodRecords[*request.Domain]; !exist {
		api.dnspodRecords[*request.Domain] = make([]*dnspod.RecordListItem, 0)
	}
	if request.TTL == nil {
		request.TTL = common.Uint64Ptr(300)
	}
	api.dnspodRecords[*request.Domain] = append(api.dnspodRecords[*request.Domain], &dnspod.RecordListItem{
		RecordId: common.Uint64Ptr(randomRecordId),
		Value:    request.Value,
		TTL:      request.TTL,
		Name:     request.SubDomain,
		Line:     request.RecordLine,
		LineId:   request.RecordLineId,
		Type:     request.RecordType,
	})
	return response, nil
}

func (api *mockAPIService) DeleteRecord(request *dnspod.DeleteRecordRequest) (response *dnspod.DeleteRecordResponse, err error) {
	result := make([]*dnspod.RecordListItem, 0)
	if _, exist := api.dnspodRecords[*request.Domain]; !exist {
		return response, nil
	}
	for _, zoneRecord := range api.dnspodRecords[*request.Domain] {
		deleteflag := false
		if request.RecordId != nil && *request.RecordId == *zoneRecord.RecordId {
			deleteflag = true
		}
		if !deleteflag {
			result = append(result, zoneRecord)
		}
	}
	api.dnspodRecords[*request.Domain] = result
	return response, nil
}

func (api *mockAPIService) ModifyRecord(request *dnspod.ModifyRecordRequest) (response *dnspod.ModifyRecordResponse, err error) {
	if _, exist := api.dnspodRecords[*request.Domain]; !exist {
		return response, nil
	}
	for _, zoneRecord := range api.dnspodRecords[*request.Domain] {
		if *zoneRecord.RecordId != *request.RecordId {
			continue
		}
		zoneRecord.Type = request.RecordType
		zoneRecord.Name = request.SubDomain
		zoneRecord.Value = request.Value
		zoneRecord.TTL = request.TTL
	}
	return response, nil
}

var letterRunes = []byte("abcdefghijklmnopqrstuvwxyz")

func RandStringRunes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandUint64() uint64 {
	return rand.Uint64()
}
