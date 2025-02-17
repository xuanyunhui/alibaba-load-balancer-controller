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

// ModifyVcoSpec invokes the vpc.ModifyVcoSpec API synchronously
func (client *Client) ModifyVcoSpec(request *ModifyVcoSpecRequest) (response *ModifyVcoSpecResponse, err error) {
	response = CreateModifyVcoSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVcoSpecWithChan invokes the vpc.ModifyVcoSpec API asynchronously
func (client *Client) ModifyVcoSpecWithChan(request *ModifyVcoSpecRequest) (<-chan *ModifyVcoSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyVcoSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVcoSpec(request)
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

// ModifyVcoSpecWithCallback invokes the vpc.ModifyVcoSpec API asynchronously
func (client *Client) ModifyVcoSpecWithCallback(request *ModifyVcoSpecRequest, callback func(response *ModifyVcoSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVcoSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyVcoSpec(request)
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

// ModifyVcoSpecRequest is the request struct for api ModifyVcoSpec
type ModifyVcoSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Spec                 string           `position:"Query" name:"Spec"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpnConnectionId      string           `position:"Query" name:"VpnConnectionId"`
}

// ModifyVcoSpecResponse is the response struct for api ModifyVcoSpec
type ModifyVcoSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVcoSpecRequest creates a request to invoke ModifyVcoSpec API
func CreateModifyVcoSpecRequest() (request *ModifyVcoSpecRequest) {
	request = &ModifyVcoSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVcoSpec", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVcoSpecResponse creates a response to parse from ModifyVcoSpec response
func CreateModifyVcoSpecResponse() (response *ModifyVcoSpecResponse) {
	response = &ModifyVcoSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
