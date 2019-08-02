package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// Dimension type
type Dimension struct {
	InstanceID string `json:"instanceId"`
}

// GetMetricListParams type
type GetMetricListParams struct {
	CommonParams
	Dimensions string
	MetricName string
	Namespace  string
	Period     string
	StartTime  string
	EndTime    string
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
	request.StartTime = params.StartTime
	request.EndTime = params.EndTime

	response, err = client.DescribeMetricList(request)

	return
}

// GetMetricReport html
func (aliyun *Aliyun) GetMetricReport(params *GetMetricListParams) (response interface{}, err error) {

	client, err := aliyun.GetClient(params.RegionID)
	if err != nil {
		return
	}

	request := cms.CreateDescribeMetricTopRequest()
	request.Scheme = "https"

	request.Namespace = params.Namespace
	request.Dimensions = params.Dimensions
	request.MetricName = params.MetricName
	request.StartTime = params.StartTime
	request.EndTime = params.EndTime
	request.Orderby = "Maximum"
	request.Length = "1"
	request.OrderDesc = "False"

	response, err = client.DescribeMetricTop(request)

	return
}
