package push

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

// QueryUniqueDeviceStat invokes the push.QueryUniqueDeviceStat API synchronously
// api document: https://help.aliyun.com/api/push/queryuniquedevicestat.html
func (client *Client) QueryUniqueDeviceStat(request *QueryUniqueDeviceStatRequest) (response *QueryUniqueDeviceStatResponse, err error) {
	response = CreateQueryUniqueDeviceStatResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUniqueDeviceStatWithChan invokes the push.QueryUniqueDeviceStat API asynchronously
// api document: https://help.aliyun.com/api/push/queryuniquedevicestat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUniqueDeviceStatWithChan(request *QueryUniqueDeviceStatRequest) (<-chan *QueryUniqueDeviceStatResponse, <-chan error) {
	responseChan := make(chan *QueryUniqueDeviceStatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUniqueDeviceStat(request)
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

// QueryUniqueDeviceStatWithCallback invokes the push.QueryUniqueDeviceStat API asynchronously
// api document: https://help.aliyun.com/api/push/queryuniquedevicestat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUniqueDeviceStatWithCallback(request *QueryUniqueDeviceStatRequest, callback func(response *QueryUniqueDeviceStatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUniqueDeviceStatResponse
		var err error
		defer close(result)
		response, err = client.QueryUniqueDeviceStat(request)
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

// QueryUniqueDeviceStatRequest is the request struct for api QueryUniqueDeviceStat
type QueryUniqueDeviceStatRequest struct {
	*requests.RpcRequest
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
	StartTime   string           `position:"Query" name:"StartTime"`
	EndTime     string           `position:"Query" name:"EndTime"`
	Granularity string           `position:"Query" name:"Granularity"`
}

// QueryUniqueDeviceStatResponse is the response struct for api QueryUniqueDeviceStat
type QueryUniqueDeviceStatResponse struct {
	*responses.BaseResponse
	RequestId      string                                `json:"RequestId" xml:"RequestId"`
	AppDeviceStats AppDeviceStatsInQueryUniqueDeviceStat `json:"AppDeviceStats" xml:"AppDeviceStats"`
}

// CreateQueryUniqueDeviceStatRequest creates a request to invoke QueryUniqueDeviceStat API
func CreateQueryUniqueDeviceStatRequest() (request *QueryUniqueDeviceStatRequest) {
	request = &QueryUniqueDeviceStatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryUniqueDeviceStat", "", "")
	return
}

// CreateQueryUniqueDeviceStatResponse creates a response to parse from QueryUniqueDeviceStat response
func CreateQueryUniqueDeviceStatResponse() (response *QueryUniqueDeviceStatResponse) {
	response = &QueryUniqueDeviceStatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}