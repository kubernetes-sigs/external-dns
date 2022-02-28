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
	"encoding/json"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
)

type defaultTencentAPIService struct {
	RetryDefault      int
	TaskCheckInterval time.Duration
	ClientSetService  TencentClientSetService
}

func NewTencentAPIService(region string, rate int, secretId string, secretKey string, internetEndpoint bool) *defaultTencentAPIService {
	tencentAPIService := &defaultTencentAPIService{
		RetryDefault:      3,
		TaskCheckInterval: 3 * time.Second,
		ClientSetService:  NewTencentClientSetService(region, rate, secretId, secretKey, internetEndpoint),
	}
	return tencentAPIService
}

////////////////////////////////////////////////////////////////
// PrivateDns API
////////////////////////////////////////////////////////////////

func (api *defaultTencentAPIService) CreatePrivateZoneRecord(request *privatedns.CreatePrivateZoneRecordRequest) (response *privatedns.CreatePrivateZoneRecordResponse, err error) {
	apiAction := CreatePrivateZoneRecord
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.PrivateDnsCli(apiAction.Name)
		if response, err = client.CreatePrivateZoneRecord(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) DeletePrivateZoneRecord(request *privatedns.DeletePrivateZoneRecordRequest) (response *privatedns.DeletePrivateZoneRecordResponse, err error) {
	apiAction := DeletePrivateZoneRecord
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.PrivateDnsCli(apiAction.Name)
		if response, err = client.DeletePrivateZoneRecord(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) ModifyPrivateZoneRecord(request *privatedns.ModifyPrivateZoneRecordRequest) (response *privatedns.ModifyPrivateZoneRecordResponse, err error) {
	apiAction := ModifyPrivateZoneRecord
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.PrivateDnsCli(apiAction.Name)
		if response, err = client.ModifyPrivateZoneRecord(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) DescribePrivateZoneList(request *privatedns.DescribePrivateZoneListRequest) (response *privatedns.DescribePrivateZoneListResponse, err error) {
	apiAction := DescribePrivateZoneList
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.PrivateDnsCli(apiAction.Name)
		if response, err = client.DescribePrivateZoneList(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) DescribePrivateZoneRecordList(request *privatedns.DescribePrivateZoneRecordListRequest) (response *privatedns.DescribePrivateZoneRecordListResponse, err error) {
	apiAction := DescribePrivateZoneRecordList
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.PrivateDnsCli(apiAction.Name)
		if response, err = client.DescribePrivateZoneRecordList(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

////////////////////////////////////////////////////////////////
// DnsPod API
////////////////////////////////////////////////////////////////

func (api *defaultTencentAPIService) DescribeDomainList(request *dnspod.DescribeDomainListRequest) (response *dnspod.DescribeDomainListResponse, err error) {
	apiAction := DescribeDomainList
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.DnsPodCli(apiAction.Name)
		if response, err = client.DescribeDomainList(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) DescribeRecordList(request *dnspod.DescribeRecordListRequest) (response *dnspod.DescribeRecordListResponse, err error) {
	apiAction := DescribeRecordList
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.DnsPodCli(apiAction.Name)
		if response, err = client.DescribeRecordList(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) CreateRecord(request *dnspod.CreateRecordRequest) (response *dnspod.CreateRecordResponse, err error) {
	apiAction := CreateRecord
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.DnsPodCli(apiAction.Name)
		if response, err = client.CreateRecord(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) DeleteRecord(request *dnspod.DeleteRecordRequest) (response *dnspod.DeleteRecordResponse, err error) {
	apiAction := DeleteRecord
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.DnsPodCli(apiAction.Name)
		if response, err = client.DeleteRecord(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

func (api *defaultTencentAPIService) ModifyRecord(request *dnspod.ModifyRecordRequest) (response *dnspod.ModifyRecordResponse, err error) {
	apiAction := ModifyRecord
	for times := 1; times <= api.RetryDefault; times++ {
		client := api.ClientSetService.DnsPodCli(apiAction.Name)
		if response, err = client.ModifyRecord(request); err != nil {
			requestJson := JsonWrapper(request)
			if retry := dealWithError(apiAction, requestJson, err); retry == false || times == api.RetryDefault {
				APIErrorRecord(apiAction, requestJson, JsonWrapper(response), err)
				return nil, err
			}
			continue
		}
		break
	}
	APIRecord(apiAction, JsonWrapper(request), JsonWrapper(response))
	return response, nil
}

////////////////////////////////////////////////////////////////
// API Error Report
////////////////////////////////////////////////////////////////

func dealWithError(action Action, request string, err error) bool {
	log.Errorf("dealWithError %s/%s request: %s, error: %s.", action.Service, action.Name, request, err.Error())
	if sdkError, ok := err.(*errors.TencentCloudSDKError); ok {
		if sdkError.Code == "RequestLimitExceeded" {
			return true
		} else if sdkError.Code == "InternalError" || sdkError.Code == "ClientError.HttpStatusCodeError" {
			return false
		} else if sdkError.Code == "ClientError.NetworkError" {
			return false
		} else if sdkError.Code == "AuthFailure.UnauthorizedOperation" || sdkError.Code == "UnauthorizedOperation.CamNoAuth" {
			return false
		}
		return false
	}

	if _, ok := err.(net.Error); ok {
		return true
	}

	return false
}

func APIErrorRecord(apiAction Action, request string, response string, err error) {
	log.Infof(fmt.Sprintf("APIError API: %s/%s Request: %s, Response: %s, Error: %s", apiAction.Service, apiAction.Name, request, response, err.Error()))
}

func APIRecord(apiAction Action, request string, response string) {
	message := fmt.Sprintf("APIRecord API: %s/%s Request: %s, Response: %s", apiAction.Service, apiAction.Name, request, response)

	if apiAction.ReadOnly {
		//log.Infof(message)
	} else {
		log.Infof(message)
	}
}

func JsonWrapper(obj interface{}) string {
	if jsonStr, jsonErr := json.Marshal(obj); jsonErr == nil {
		return string(jsonStr)
	}
	return "json_format_error"
}
