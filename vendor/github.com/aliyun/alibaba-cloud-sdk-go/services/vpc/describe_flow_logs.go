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

// DescribeFlowLogs invokes the vpc.DescribeFlowLogs API synchronously
func (client *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
	response = CreateDescribeFlowLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowLogsWithChan invokes the vpc.DescribeFlowLogs API asynchronously
func (client *Client) DescribeFlowLogsWithChan(request *DescribeFlowLogsRequest) (<-chan *DescribeFlowLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowLogs(request)
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

// DescribeFlowLogsWithCallback invokes the vpc.DescribeFlowLogs API asynchronously
func (client *Client) DescribeFlowLogsWithCallback(request *DescribeFlowLogsRequest, callback func(response *DescribeFlowLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowLogs(request)
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

// DescribeFlowLogsRequest is the request struct for api DescribeFlowLogs
type DescribeFlowLogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer        `position:"Query" name:"ResourceOwnerId"`
	Description          string                  `position:"Query" name:"Description"`
	PageNumber           requests.Integer        `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                  `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer        `position:"Query" name:"PageSize"`
	ResourceId           string                  `position:"Query" name:"ResourceId"`
	ProjectName          string                  `position:"Query" name:"ProjectName"`
	LogStoreName         string                  `position:"Query" name:"LogStoreName"`
	ResourceOwnerAccount string                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer        `position:"Query" name:"OwnerId"`
	ResourceType         string                  `position:"Query" name:"ResourceType"`
	Tags                 *[]DescribeFlowLogsTags `position:"Query" name:"Tags"  type:"Repeated"`
	VpcId                string                  `position:"Query" name:"VpcId"`
	TrafficType          string                  `position:"Query" name:"TrafficType"`
	FlowLogId            string                  `position:"Query" name:"FlowLogId"`
	FlowLogName          string                  `position:"Query" name:"FlowLogName"`
	Status               string                  `position:"Query" name:"Status"`
}

// DescribeFlowLogsTags is a repeated param struct in DescribeFlowLogsRequest
type DescribeFlowLogsTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeFlowLogsResponse is the response struct for api DescribeFlowLogs
type DescribeFlowLogsResponse struct {
	*responses.BaseResponse
	PageSize   string   `json:"PageSize" xml:"PageSize"`
	PageNumber string   `json:"PageNumber" xml:"PageNumber"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount string   `json:"TotalCount" xml:"TotalCount"`
	Success    string   `json:"Success" xml:"Success"`
	FlowLogs   FlowLogs `json:"FlowLogs" xml:"FlowLogs"`
}

// CreateDescribeFlowLogsRequest creates a request to invoke DescribeFlowLogs API
func CreateDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
	request = &DescribeFlowLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeFlowLogs", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFlowLogsResponse creates a response to parse from DescribeFlowLogs response
func CreateDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
	response = &DescribeFlowLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
