package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeStrategy invokes the rds.DescribeStrategy API synchronously
// api document: https://help.aliyun.com/api/rds/describestrategy.html
func (client *Client) DescribeStrategy(request *DescribeStrategyRequest) (response *DescribeStrategyResponse, err error) {
	response = CreateDescribeStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStrategyWithChan invokes the rds.DescribeStrategy API asynchronously
// api document: https://help.aliyun.com/api/rds/describestrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStrategyWithChan(request *DescribeStrategyRequest) (<-chan *DescribeStrategyResponse, <-chan error) {
	responseChan := make(chan *DescribeStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStrategy(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeStrategyWithCallback invokes the rds.DescribeStrategy API asynchronously
// api document: https://help.aliyun.com/api/rds/describestrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStrategyWithCallback(request *DescribeStrategyRequest, callback func(response *DescribeStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStrategyResponse
		var err error
		defer close(result)
		response, err = client.DescribeStrategy(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeStrategyRequest is the request struct for api DescribeStrategy
type DescribeStrategyRequest struct {
	*requests.RpcRequest
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
}

// DescribeStrategyResponse is the response struct for api DescribeStrategy
type DescribeStrategyResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ReplicaId        string `json:"ReplicaId" xml:"ReplicaId"`
	RecoveryMode     string `json:"RecoveryMode" xml:"RecoveryMode"`
	VerificationMode string `json:"VerificationMode" xml:"VerificationMode"`
}

// CreateDescribeStrategyRequest creates a request to invoke DescribeStrategy API
func CreateDescribeStrategyRequest() (request *DescribeStrategyRequest) {
	request = &DescribeStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeStrategy", "rds", "openAPI")
	return
}

// CreateDescribeStrategyResponse creates a response to parse from DescribeStrategy response
func CreateDescribeStrategyResponse() (response *DescribeStrategyResponse) {
	response = &DescribeStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
