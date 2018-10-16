package live

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

// DescribeLiveStreamsControlHistory invokes the live.DescribeLiveStreamsControlHistory API synchronously
// api document: https://help.aliyun.com/api/live/describelivestreamscontrolhistory.html
func (client *Client) DescribeLiveStreamsControlHistory(request *DescribeLiveStreamsControlHistoryRequest) (response *DescribeLiveStreamsControlHistoryResponse, err error) {
	response = CreateDescribeLiveStreamsControlHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamsControlHistoryWithChan invokes the live.DescribeLiveStreamsControlHistory API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamscontrolhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamsControlHistoryWithChan(request *DescribeLiveStreamsControlHistoryRequest) (<-chan *DescribeLiveStreamsControlHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamsControlHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamsControlHistory(request)
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

// DescribeLiveStreamsControlHistoryWithCallback invokes the live.DescribeLiveStreamsControlHistory API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamscontrolhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamsControlHistoryWithCallback(request *DescribeLiveStreamsControlHistoryRequest, callback func(response *DescribeLiveStreamsControlHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamsControlHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamsControlHistory(request)
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

// DescribeLiveStreamsControlHistoryRequest is the request struct for api DescribeLiveStreamsControlHistory
type DescribeLiveStreamsControlHistoryRequest struct {
	*requests.RpcRequest
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamsControlHistoryResponse is the response struct for api DescribeLiveStreamsControlHistory
type DescribeLiveStreamsControlHistoryResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ControlInfo ControlInfo `json:"ControlInfo" xml:"ControlInfo"`
}

// CreateDescribeLiveStreamsControlHistoryRequest creates a request to invoke DescribeLiveStreamsControlHistory API
func CreateDescribeLiveStreamsControlHistoryRequest() (request *DescribeLiveStreamsControlHistoryRequest) {
	request = &DescribeLiveStreamsControlHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsControlHistory", "live", "openAPI")
	return
}

// CreateDescribeLiveStreamsControlHistoryResponse creates a response to parse from DescribeLiveStreamsControlHistory response
func CreateDescribeLiveStreamsControlHistoryResponse() (response *DescribeLiveStreamsControlHistoryResponse) {
	response = &DescribeLiveStreamsControlHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
