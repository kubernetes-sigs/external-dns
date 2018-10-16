package domain

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

// SaveSingleTaskForTransferProhibitionLock invokes the domain.SaveSingleTaskForTransferProhibitionLock API synchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskfortransferprohibitionlock.html
func (client *Client) SaveSingleTaskForTransferProhibitionLock(request *SaveSingleTaskForTransferProhibitionLockRequest) (response *SaveSingleTaskForTransferProhibitionLockResponse, err error) {
	response = CreateSaveSingleTaskForTransferProhibitionLockResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForTransferProhibitionLockWithChan invokes the domain.SaveSingleTaskForTransferProhibitionLock API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskfortransferprohibitionlock.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForTransferProhibitionLockWithChan(request *SaveSingleTaskForTransferProhibitionLockRequest) (<-chan *SaveSingleTaskForTransferProhibitionLockResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForTransferProhibitionLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForTransferProhibitionLock(request)
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

// SaveSingleTaskForTransferProhibitionLockWithCallback invokes the domain.SaveSingleTaskForTransferProhibitionLock API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskfortransferprohibitionlock.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForTransferProhibitionLockWithCallback(request *SaveSingleTaskForTransferProhibitionLockRequest, callback func(response *SaveSingleTaskForTransferProhibitionLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForTransferProhibitionLockResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForTransferProhibitionLock(request)
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

// SaveSingleTaskForTransferProhibitionLockRequest is the request struct for api SaveSingleTaskForTransferProhibitionLock
type SaveSingleTaskForTransferProhibitionLockRequest struct {
	*requests.RpcRequest
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	DomainName   string           `position:"Query" name:"DomainName"`
	Lang         string           `position:"Query" name:"Lang"`
	Status       requests.Boolean `position:"Query" name:"Status"`
}

// SaveSingleTaskForTransferProhibitionLockResponse is the response struct for api SaveSingleTaskForTransferProhibitionLock
type SaveSingleTaskForTransferProhibitionLockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForTransferProhibitionLockRequest creates a request to invoke SaveSingleTaskForTransferProhibitionLock API
func CreateSaveSingleTaskForTransferProhibitionLockRequest() (request *SaveSingleTaskForTransferProhibitionLockRequest) {
	request = &SaveSingleTaskForTransferProhibitionLockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForTransferProhibitionLock", "", "")
	return
}

// CreateSaveSingleTaskForTransferProhibitionLockResponse creates a response to parse from SaveSingleTaskForTransferProhibitionLock response
func CreateSaveSingleTaskForTransferProhibitionLockResponse() (response *SaveSingleTaskForTransferProhibitionLockResponse) {
	response = &SaveSingleTaskForTransferProhibitionLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
