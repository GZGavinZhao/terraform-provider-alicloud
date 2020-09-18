package dcdn

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

// PreloadDcdnObjectCaches invokes the dcdn.PreloadDcdnObjectCaches API synchronously
// api document: https://help.aliyun.com/api/dcdn/preloaddcdnobjectcaches.html
func (client *Client) PreloadDcdnObjectCaches(request *PreloadDcdnObjectCachesRequest) (response *PreloadDcdnObjectCachesResponse, err error) {
	response = CreatePreloadDcdnObjectCachesResponse()
	err = client.DoAction(request, response)
	return
}

// PreloadDcdnObjectCachesWithChan invokes the dcdn.PreloadDcdnObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/dcdn/preloaddcdnobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreloadDcdnObjectCachesWithChan(request *PreloadDcdnObjectCachesRequest) (<-chan *PreloadDcdnObjectCachesResponse, <-chan error) {
	responseChan := make(chan *PreloadDcdnObjectCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreloadDcdnObjectCaches(request)
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

// PreloadDcdnObjectCachesWithCallback invokes the dcdn.PreloadDcdnObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/dcdn/preloaddcdnobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreloadDcdnObjectCachesWithCallback(request *PreloadDcdnObjectCachesRequest, callback func(response *PreloadDcdnObjectCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreloadDcdnObjectCachesResponse
		var err error
		defer close(result)
		response, err = client.PreloadDcdnObjectCaches(request)
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

// PreloadDcdnObjectCachesRequest is the request struct for api PreloadDcdnObjectCaches
type PreloadDcdnObjectCachesRequest struct {
	*requests.RpcRequest
	Area          string           `position:"Query" name:"Area"`
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// PreloadDcdnObjectCachesResponse is the response struct for api PreloadDcdnObjectCaches
type PreloadDcdnObjectCachesResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	PreloadTaskId string `json:"PreloadTaskId" xml:"PreloadTaskId"`
}

// CreatePreloadDcdnObjectCachesRequest creates a request to invoke PreloadDcdnObjectCaches API
func CreatePreloadDcdnObjectCachesRequest() (request *PreloadDcdnObjectCachesRequest) {
	request = &PreloadDcdnObjectCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "PreloadDcdnObjectCaches", "", "")
	request.Method = requests.POST
	return
}

// CreatePreloadDcdnObjectCachesResponse creates a response to parse from PreloadDcdnObjectCaches response
func CreatePreloadDcdnObjectCachesResponse() (response *PreloadDcdnObjectCachesResponse) {
	response = &PreloadDcdnObjectCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}