package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// Namespaces of aliyun monitor product
var Namespaces = struct {
	CDN   string
	SLB   string
	Redis string
	ECS   string
	MySQL string
}{
	CDN:   "acs_cdn",
	SLB:   "acs_slb_dashboard",
	Redis: "acs_kvstore",
	ECS:   "acs_ecs_dashboard",
	MySQL: "acs_rds_dashboard",
}

// GetMetricListParams type
type GetMetricListParams struct {
	CommonParams
	Dimensions string
	MetricName string
	Namespace  string
	Period     string
}

// GetMetricList data
func (aliyun *Aliyun) GetMetricList(params *GetMetricListParams) (response *cms.DescribeMetricListResponse, err error) {

	client, err := aliyun.GetClient(params.RegionID)
	if err != nil {
		return
	}

	request := cms.CreateDescribeMetricListRequest()
	request.Scheme = "https"

	request.MetricName = params.MetricName
	request.Namespace = params.Namespace
	request.Period = params.Period
	request.Dimensions = params.Dimensions

	response, err = client.DescribeMetricList(request)

	return
}
