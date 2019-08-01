package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

type MonitorNamespace struct {
	Namespace   string
	Description string
}

// GetMonitorNamespaces support
func (aliyun *Aliyun) GetMonitorNamespaces() (namespaces []*MonitorNamespace, err error) {
	client, err := aliyun.GetClient("default")
	if err != nil {
		return
	}
	request := cms.CreateDescribeProjectMetaRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(999)

	response, err := client.DescribeProjectMeta(request)
	if err != nil {
		return
	}

	for _, item := range response.Resources.Resource {
		namespace := &MonitorNamespace{item.Namespace, item.Description}
		namespaces = append(namespaces, namespace)
	}

	return
}
