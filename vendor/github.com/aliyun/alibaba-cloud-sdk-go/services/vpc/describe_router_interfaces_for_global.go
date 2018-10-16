package vpc

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

// DescribeRouterInterfacesForGlobal invokes the vpc.DescribeRouterInterfacesForGlobal API synchronously
// api document: https://help.aliyun.com/api/vpc/describerouterinterfacesforglobal.html
func (client *Client) DescribeRouterInterfacesForGlobal(request *DescribeRouterInterfacesForGlobalRequest) (response *DescribeRouterInterfacesForGlobalResponse, err error) {
	response = CreateDescribeRouterInterfacesForGlobalResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRouterInterfacesForGlobalWithChan invokes the vpc.DescribeRouterInterfacesForGlobal API asynchronously
// api document: https://help.aliyun.com/api/vpc/describerouterinterfacesforglobal.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouterInterfacesForGlobalWithChan(request *DescribeRouterInterfacesForGlobalRequest) (<-chan *DescribeRouterInterfacesForGlobalResponse, <-chan error) {
	responseChan := make(chan *DescribeRouterInterfacesForGlobalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouterInterfacesForGlobal(request)
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

// DescribeRouterInterfacesForGlobalWithCallback invokes the vpc.DescribeRouterInterfacesForGlobal API asynchronously
// api document: https://help.aliyun.com/api/vpc/describerouterinterfacesforglobal.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouterInterfacesForGlobalWithCallback(request *DescribeRouterInterfacesForGlobalRequest, callback func(response *DescribeRouterInterfacesForGlobalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouterInterfacesForGlobalResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouterInterfacesForGlobal(request)
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

// DescribeRouterInterfacesForGlobalRequest is the request struct for api DescribeRouterInterfacesForGlobal
type DescribeRouterInterfacesForGlobalRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeRouterInterfacesForGlobalResponse is the response struct for api DescribeRouterInterfacesForGlobal
type DescribeRouterInterfacesForGlobalResponse struct {
	*responses.BaseResponse
	RequestId          string                                                `json:"RequestId" xml:"RequestId"`
	Code               string                                                `json:"Code" xml:"Code"`
	Message            string                                                `json:"Message" xml:"Message"`
	Desc               string                                                `json:"desc" xml:"desc"`
	Success            bool                                                  `json:"Success" xml:"Success"`
	PageSize           int                                                   `json:"PageSize" xml:"PageSize"`
	PageNumber         int                                                   `json:"PageNumber" xml:"PageNumber"`
	TotalCount         int                                                   `json:"TotalCount" xml:"TotalCount"`
	RouterInterfaceSet RouterInterfaceSetInDescribeRouterInterfacesForGlobal `json:"RouterInterfaceSet" xml:"RouterInterfaceSet"`
}

// CreateDescribeRouterInterfacesForGlobalRequest creates a request to invoke DescribeRouterInterfacesForGlobal API
func CreateDescribeRouterInterfacesForGlobalRequest() (request *DescribeRouterInterfacesForGlobalRequest) {
	request = &DescribeRouterInterfacesForGlobalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfacesForGlobal", "vpc", "openAPI")
	return
}

// CreateDescribeRouterInterfacesForGlobalResponse creates a response to parse from DescribeRouterInterfacesForGlobal response
func CreateDescribeRouterInterfacesForGlobalResponse() (response *DescribeRouterInterfacesForGlobalResponse) {
	response = &DescribeRouterInterfacesForGlobalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
