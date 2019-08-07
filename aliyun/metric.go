package aliyun

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// Datapoint type
type Datapoint struct {
	Timestamp  int
	InstanceID string `json:"instanceId"`
	Minimum    float64
	Average    float64
	Maximum    float64
	// Order      int
	// UserID     string `json:"userId"`
	// Count      int `json:"_count"`
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
	Express    string
	Length     string
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
	request.Express = params.Express
	request.Length = params.Length

	response, err = client.DescribeMetricList(request)

	return
}

// GetMetricTopParams type
type GetMetricTopParams struct {
	GetMetricListParams
	Orderby string // Maximum Average
}

// GetMetricTop data
func (aliyun *Aliyun) GetMetricTop(params *GetMetricTopParams) (datapoints []Datapoint, err error) {

	orderBy := params.Orderby
	if orderBy == "" {
		orderBy = "Maximum"
	}

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
	request.Express = fmt.Sprintf(`{"orderby":"%v"}`, orderBy)
	request.Length = "1"

	response, err := client.DescribeMetricList(request)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(response.Datapoints), &datapoints); err != nil {
		return
	}

	return
}

func (aliyun *Aliyun) notWorkGetMetricTop(params *GetMetricTopParams) (datapoints []Datapoint, err error) {

	orderBy := params.Orderby
	if orderBy == "" {
		orderBy = "Maximum"
	}

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
	request.Orderby = orderBy
	request.Length = "1"
	request.OrderDesc = "False"

	response, err := client.DescribeMetricTop(request)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(response.Datapoints), &datapoints); err != nil {
		return
	}

	return
}
