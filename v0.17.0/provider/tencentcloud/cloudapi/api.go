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
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
)

type Action struct {
	Service  string `json:"service"`
	Name     string `json:"name"`
	ReadOnly bool   `json:"readOnly"`
}

var (
	/* PrivateDNS */
	CreatePrivateZoneRecord       = Action{Service: "PrivateDns", Name: "CreatePrivateZoneRecord", ReadOnly: false}
	DeletePrivateZoneRecord       = Action{Service: "PrivateDns", Name: "DeletePrivateZoneRecord", ReadOnly: false}
	ModifyPrivateZoneRecord       = Action{Service: "PrivateDns", Name: "ModifyPrivateZoneRecord", ReadOnly: false}
	DescribePrivateZoneList       = Action{Service: "PrivateDns", Name: "DescribePrivateZoneList", ReadOnly: true}
	DescribePrivateZoneRecordList = Action{Service: "PrivateDns", Name: "DescribePrivateZoneRecordList", ReadOnly: true}

	/* DNSPod */
	DescribeDomainList = Action{Service: "DnsPod", Name: "DescribeDomainList", ReadOnly: true}
	DescribeRecordList = Action{Service: "DnsPod", Name: "DescribeRecordList", ReadOnly: true}
	CreateRecord       = Action{Service: "DnsPod", Name: "CreateRecord", ReadOnly: false}
	DeleteRecord       = Action{Service: "DnsPod", Name: "DeleteRecord", ReadOnly: false}
	ModifyRecord       = Action{Service: "DnsPod", Name: "ModifyRecord", ReadOnly: false}
)

type TencentAPIService interface {
	// PrivateDNS
	CreatePrivateZoneRecord(request *privatedns.CreatePrivateZoneRecordRequest) (response *privatedns.CreatePrivateZoneRecordResponse, err error)
	DeletePrivateZoneRecord(request *privatedns.DeletePrivateZoneRecordRequest) (response *privatedns.DeletePrivateZoneRecordResponse, err error)
	ModifyPrivateZoneRecord(request *privatedns.ModifyPrivateZoneRecordRequest) (response *privatedns.ModifyPrivateZoneRecordResponse, err error)
	DescribePrivateZoneList(request *privatedns.DescribePrivateZoneListRequest) (response *privatedns.DescribePrivateZoneListResponse, err error)
	DescribePrivateZoneRecordList(request *privatedns.DescribePrivateZoneRecordListRequest) (response *privatedns.DescribePrivateZoneRecordListResponse, err error)

	// DNSPod
	DescribeDomainList(request *dnspod.DescribeDomainListRequest) (response *dnspod.DescribeDomainListResponse, err error)
	DescribeRecordList(request *dnspod.DescribeRecordListRequest) (response *dnspod.DescribeRecordListResponse, err error)
	CreateRecord(request *dnspod.CreateRecordRequest) (response *dnspod.CreateRecordResponse, err error)
	DeleteRecord(request *dnspod.DeleteRecordRequest) (response *dnspod.DeleteRecordResponse, err error)
	ModifyRecord(request *dnspod.ModifyRecordRequest) (response *dnspod.ModifyRecordResponse, err error)
}
