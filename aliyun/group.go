package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// GetGroupListParams type
type GetGroupListParams struct {
	RegionID string `json:"RegionId"`
	Keyword  string
}

// GetGroupList func
func (aliyun *Aliyun) GetGroupList(params *GetGroupListParams) (response *cms.DescribeMonitorGroupsResponse, err error) {

	client, err := cms.NewClientWithAccessKey(params.RegionID, aliyun.AccessKey, aliyun.AccessKeySecret)

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
