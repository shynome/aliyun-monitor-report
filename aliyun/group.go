package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// GetGroupListParams type
type GetGroupListParams struct {
	CommonParams
	Keyword string
}

// GetGroupList func
func (aliyun *Aliyun) GetGroupList(params *GetGroupListParams) (response *cms.DescribeMonitorGroupsResponse, err error) {

	client, err := aliyun.GetClient(params.RegionID)
	if err != nil {
		return
	}

	request := cms.CreateDescribeMonitorGroupsRequest()
	request.Scheme = "https"
	request.Keyword = params.Keyword
	request.PageSize = requests.NewInteger(99)

	response, err = client.DescribeMonitorGroups(request)
	if err != nil {
		return
	}
	return

}

// GetGroupDetailsParams type
type GetGroupDetailsParams struct {
	CommonParams
	GroupID  int    `json:"GroupId"`
	Category string ``
	Keyword  string ``
}

// GetGroupDetails by id
func (aliyun *Aliyun) GetGroupDetails(params *GetGroupDetailsParams) (response *cms.DescribeMonitorGroupInstancesResponse, err error) {

	client, err := aliyun.GetClient(params.RegionID)
	if err != nil {
		return
	}

	request := cms.CreateDescribeMonitorGroupInstancesRequest()
	request.Scheme = "https"
	request.GroupId = requests.NewInteger(params.GroupID)
	request.Keyword = params.Keyword
	request.Category = params.Category
	request.PageSize = requests.NewInteger(99)

	response, err = client.DescribeMonitorGroupInstances(request)

	return
}
