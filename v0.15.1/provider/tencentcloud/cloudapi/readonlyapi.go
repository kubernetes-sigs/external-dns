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

type readonlyAPIService struct {
	defaultTencentAPIService
}

func NewReadOnlyAPIService(region string, rate int, secretId string, secretKey string, internetEndpoint bool) *readonlyAPIService {
	apiService := NewTencentAPIService(region, rate, secretId, secretKey, internetEndpoint)
	tencentAPIService := &readonlyAPIService{
		*apiService,
	}
	return tencentAPIService
}

////////////////////////////////////////////////////////////////
// PrivateDns API
////////////////////////////////////////////////////////////////

func (api *readonlyAPIService) CreatePrivateZoneRecord(request *privatedns.CreatePrivateZoneRecordRequest) (response *privatedns.CreatePrivateZoneRecordResponse, err error) {
	apiAction := CreatePrivateZoneRecord
	APIRecord(apiAction, JsonWrapper(request), "dryRun")
	return response, nil
}

func (api *readonlyAPIService) DeletePrivateZoneRecord(request *privatedns.DeletePrivateZoneRecordRequest) (response *privatedns.DeletePrivateZoneRecordResponse, err error) {
	apiAction := DeletePrivateZoneRecord
	APIRecord(apiAction, JsonWrapper(request), "dryRun")
	return response, nil
}

func (api *readonlyAPIService) ModifyPrivateZoneRecord(request *privatedns.ModifyPrivateZoneRecordRequest) (response *privatedns.ModifyPrivateZoneRecordResponse, err error) {
	apiAction := ModifyPrivateZoneRecord
	APIRecord(apiAction, JsonWrapper(request), "dryRun")
	return response, nil
}

////////////////////////////////////////////////////////////////
// DnsPod API
////////////////////////////////////////////////////////////////

func (api *readonlyAPIService) CreateRecord(request *dnspod.CreateRecordRequest) (response *dnspod.CreateRecordResponse, err error) {
	apiAction := CreateRecord
	APIRecord(apiAction, JsonWrapper(request), "dryRun")
	return response, nil
}

func (api *readonlyAPIService) DeleteRecord(request *dnspod.DeleteRecordRequest) (response *dnspod.DeleteRecordResponse, err error) {
	apiAction := DeleteRecord
	APIRecord(apiAction, JsonWrapper(request), "dryRun")
	return response, nil
}

func (api *readonlyAPIService) ModifyRecord(request *dnspod.ModifyRecordRequest) (response *dnspod.ModifyRecordResponse, err error) {
	apiAction := ModifyRecord
	APIRecord(apiAction, JsonWrapper(request), "dryRun")
	return response, nil
}
