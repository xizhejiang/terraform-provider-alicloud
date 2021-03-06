package emr

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

// MetastoreCreateDataResource invokes the emr.MetastoreCreateDataResource API synchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatedataresource.html
func (client *Client) MetastoreCreateDataResource(request *MetastoreCreateDataResourceRequest) (response *MetastoreCreateDataResourceResponse, err error) {
	response = CreateMetastoreCreateDataResourceResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreCreateDataResourceWithChan invokes the emr.MetastoreCreateDataResource API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatedataresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreCreateDataResourceWithChan(request *MetastoreCreateDataResourceRequest) (<-chan *MetastoreCreateDataResourceResponse, <-chan error) {
	responseChan := make(chan *MetastoreCreateDataResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreCreateDataResource(request)
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

// MetastoreCreateDataResourceWithCallback invokes the emr.MetastoreCreateDataResource API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatedataresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreCreateDataResourceWithCallback(request *MetastoreCreateDataResourceRequest, callback func(response *MetastoreCreateDataResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreCreateDataResourceResponse
		var err error
		defer close(result)
		response, err = client.MetastoreCreateDataResource(request)
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

// MetastoreCreateDataResourceRequest is the request struct for api MetastoreCreateDataResource
type MetastoreCreateDataResourceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Default         requests.Boolean `position:"Query" name:"Default"`
	AccessType      string           `position:"Query" name:"AccessType"`
	Name            string           `position:"Query" name:"Name"`
	Description     string           `position:"Query" name:"Description"`
	MetaType        string           `position:"Query" name:"MetaType"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// MetastoreCreateDataResourceResponse is the response struct for api MetastoreCreateDataResource
type MetastoreCreateDataResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateMetastoreCreateDataResourceRequest creates a request to invoke MetastoreCreateDataResource API
func CreateMetastoreCreateDataResourceRequest() (request *MetastoreCreateDataResourceRequest) {
	request = &MetastoreCreateDataResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreCreateDataResource", "emr", "openAPI")
	return
}

// CreateMetastoreCreateDataResourceResponse creates a response to parse from MetastoreCreateDataResource response
func CreateMetastoreCreateDataResourceResponse() (response *MetastoreCreateDataResourceResponse) {
	response = &MetastoreCreateDataResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
