package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// GetMonitorNamespaces support
func (aliyun *Aliyun) GetMonitorNamespaces() (response *cms.DescribeProjectMetaResponse, err error) {
	client, err := aliyun.GetClient("default")
	if err != nil {
		return
	}
	request := cms.CreateDescribeProjectMetaRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(999)

	response, err = client.DescribeProjectMeta(request)
	return
}
