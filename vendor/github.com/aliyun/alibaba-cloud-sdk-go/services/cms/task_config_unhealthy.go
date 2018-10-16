package cms

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

// TaskConfigUnhealthy invokes the cms.TaskConfigUnhealthy API synchronously
// api document: https://help.aliyun.com/api/cms/taskconfigunhealthy.html
func (client *Client) TaskConfigUnhealthy(request *TaskConfigUnhealthyRequest) (response *TaskConfigUnhealthyResponse, err error) {
	response = CreateTaskConfigUnhealthyResponse()
	err = client.DoAction(request, response)
	return
}

// TaskConfigUnhealthyWithChan invokes the cms.TaskConfigUnhealthy API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfigunhealthy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigUnhealthyWithChan(request *TaskConfigUnhealthyRequest) (<-chan *TaskConfigUnhealthyResponse, <-chan error) {
	responseChan := make(chan *TaskConfigUnhealthyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaskConfigUnhealthy(request)
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

// TaskConfigUnhealthyWithCallback invokes the cms.TaskConfigUnhealthy API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfigunhealthy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigUnhealthyWithCallback(request *TaskConfigUnhealthyRequest, callback func(response *TaskConfigUnhealthyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaskConfigUnhealthyResponse
		var err error
		defer close(result)
		response, err = client.TaskConfigUnhealthy(request)
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

// TaskConfigUnhealthyRequest is the request struct for api TaskConfigUnhealthy
type TaskConfigUnhealthyRequest struct {
	*requests.RpcRequest
	TaskIdList *[]string `position:"Query" name:"TaskIdList"  type:"Repeated"`
}

// TaskConfigUnhealthyResponse is the response struct for api TaskConfigUnhealthy
type TaskConfigUnhealthyResponse struct {
	*responses.BaseResponse
	ErrorCode     int           `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage  string        `json:"ErrorMessage" xml:"ErrorMessage"`
	Success       bool          `json:"Success" xml:"Success"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	UnhealthyList UnhealthyList `json:"UnhealthyList" xml:"UnhealthyList"`
}

// CreateTaskConfigUnhealthyRequest creates a request to invoke TaskConfigUnhealthy API
func CreateTaskConfigUnhealthyRequest() (request *TaskConfigUnhealthyRequest) {
	request = &TaskConfigUnhealthyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigUnhealthy", "cms", "openAPI")
	return
}

// CreateTaskConfigUnhealthyResponse creates a response to parse from TaskConfigUnhealthy response
func CreateTaskConfigUnhealthyResponse() (response *TaskConfigUnhealthyResponse) {
	response = &TaskConfigUnhealthyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
