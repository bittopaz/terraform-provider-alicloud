package slb

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

// SetServerCertificateName invokes the slb.SetServerCertificateName API synchronously
// api document: https://help.aliyun.com/api/slb/setservercertificatename.html
func (client *Client) SetServerCertificateName(request *SetServerCertificateNameRequest) (response *SetServerCertificateNameResponse, err error) {
	response = CreateSetServerCertificateNameResponse()
	err = client.DoAction(request, response)
	return
}

// SetServerCertificateNameWithChan invokes the slb.SetServerCertificateName API asynchronously
// api document: https://help.aliyun.com/api/slb/setservercertificatename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetServerCertificateNameWithChan(request *SetServerCertificateNameRequest) (<-chan *SetServerCertificateNameResponse, <-chan error) {
	responseChan := make(chan *SetServerCertificateNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetServerCertificateName(request)
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

// SetServerCertificateNameWithCallback invokes the slb.SetServerCertificateName API asynchronously
// api document: https://help.aliyun.com/api/slb/setservercertificatename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetServerCertificateNameWithCallback(request *SetServerCertificateNameRequest, callback func(response *SetServerCertificateNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetServerCertificateNameResponse
		var err error
		defer close(result)
		response, err = client.SetServerCertificateName(request)
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

// SetServerCertificateNameRequest is the request struct for api SetServerCertificateName
type SetServerCertificateNameRequest struct {
	*requests.RpcRequest
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ServerCertificateId   string           `position:"Query" name:"ServerCertificateId"`
	ServerCertificateName string           `position:"Query" name:"ServerCertificateName"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId           string           `position:"Query" name:"access_key_id"`
	Tags                  string           `position:"Query" name:"Tags"`
}

// SetServerCertificateNameResponse is the response struct for api SetServerCertificateName
type SetServerCertificateNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetServerCertificateNameRequest creates a request to invoke SetServerCertificateName API
func CreateSetServerCertificateNameRequest() (request *SetServerCertificateNameRequest) {
	request = &SetServerCertificateNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetServerCertificateName", "slb", "openAPI")
	return
}

// CreateSetServerCertificateNameResponse creates a response to parse from SetServerCertificateName response
func CreateSetServerCertificateNameResponse() (response *SetServerCertificateNameResponse) {
	response = &SetServerCertificateNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
