package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// GetBaseProjects support
func (aliyun *Aliyun) GetBaseProjects() (response *cms.DescribeProjectMetaResponse, err error) {
	client, err := aliyun.GetClient("default")
	if err != nil {
		return
	}
	request := cms.CreateDescribeProjectMetaRequest()
	request.Scheme = "https"

	response, err = client.DescribeProjectMeta(request)
	return
}
