package aliyun

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GetMetricReportParams type
type GetMetricReportParams struct {
	StartTime string ``
	EndTime   string ``
	GroupID   int    `json:"GroupId"` // 应用分组的 ID
}

// DimensionReport of instance
type DimensionReport struct {
	Name        string
	DisplayName string
	Max         string
	Avg         string
}

// MetricReport type
type MetricReport struct {
	GroupResource
	Report []DimensionReport
}

// MetricReportResponse type
type MetricReportResponse struct {
	Errors []string
	Report map[string][]MetricReport
}

// ReportDimension type
type ReportDimension struct {
	DisplayName string
	Name        string
}

// DefaultReportDimensions type
var DefaultReportDimensions = map[string][]ReportDimension{
	"ECS": {
		{"CPU", "cpu_total"},
		{"内存", "memory_usedutilization"},
		{"连接数", "concurrentConnections"},
	},
	"RDS": {
		{"CPU 使用率", ""},
		{"内存使用率", ""},
		{"连接数使用率", ""},
	},
	"KVSTORE": {
		{"CPU 使用率", ""},
		{"内存使用率", ""},
		{"连接数使用率", ""},
	},
	"SLB": {
		{"CPU 使用率", ""},
		{"内存使用率", ""},
		{"连接数使用率", ""},
	},
	"CDN": {
		{"宽带峰值", ""},
		{"下行流量", ""},
		{"QPS", ""},
	},
}

type getMetricReportParams struct {
	*GetMetricReportParams
	category  string
	resources []*GroupResource
	report    chan MetricReport
	err       chan error
}

func getNamespace(category string) string {
	return "acs_" + strings.ToLower(category)
}

func (aliyun *Aliyun) getMetricReport(params getMetricReportParams) {

	resp := map[string]map[string]DimensionReport{}
	dimensions := make([]Dimension, len(params.resources))

	for i, r := range params.resources {
		dimensions[i] = Dimension{InstanceID: r.InstanceID}
		resp[r.InstanceID] = map[string]DimensionReport{}
		for _, d := range DefaultReportDimensions[params.category] {
			resp[r.InstanceID][d.Name] = DimensionReport{}
		}
	}

	dimensionsBytes, jsonError := json.Marshal(dimensions)
	if jsonError != nil {
		params.err <- jsonError
		return
	}

	type PeakData struct {
		ReportDimension
		Datapoints []Datapoint
	}
	maxChan, avgChan, errChan := make(chan PeakData), make(chan PeakData), make(chan error)

	namespace := getNamespace(params.category)
	getPeakDataOrderBy := func(d ReportDimension, orderBy string, data chan PeakData) {

		datapoints, err := aliyun.GetMetricTop(&GetMetricTopParams{
			GetMetricListParams: GetMetricListParams{
				Dimensions: string(dimensionsBytes),
				StartTime:  params.StartTime,
				EndTime:    params.EndTime,
				Namespace:  namespace,
				MetricName: d.Name,
			},
			Orderby: orderBy,
		})
		if err != nil {
			errChan <- err
			return
		}
		data <- PeakData{d, datapoints}
	}

	for _, d := range DefaultReportDimensions[params.category] {
		go getPeakDataOrderBy(d, "Maximum", maxChan)
		go getPeakDataOrderBy(d, "Average", avgChan)
	}

	var errs string
	select {
	case max := <-maxChan:
		for _, datapoint := range max.Datapoints {
			v := resp[datapoint.InstanceID][max.Name]
			v.Max = fmt.Sprintf("%v", datapoint.Maximum)
		}
	case avg := <-avgChan:
		for _, datapoint := range avg.Datapoints {
			v := resp[datapoint.InstanceID][avg.Name]
			v.Max = fmt.Sprintf("%v", datapoint.Average)
		}
	case e := <-errChan:
		errs += e.Error()
	}
	fmt.Println(resp)
	params.err <- fmt.Errorf(errs)
}

// GetMetricReport html
func (aliyun *Aliyun) GetMetricReport(params *GetMetricReportParams) (html string, err error) {
	if params.StartTime == "" || params.EndTime == "" {
		err = fmt.Errorf("开始时间和结束时间必须要指定, 格式支持 2000-01-01 00:00:00, 并且结束时间不能大于开始时间")
		return
	}
	resources, err := aliyun.GetGroupResources(&GetGroupResourcesParams{GroupID: params.GroupID})
	if err != nil {
		return
	}
	errs := make(chan error, len(resources))
	report := make(chan MetricReport, len(resources))
	quit := make(chan int)
	d := map[string][]*GroupResource{}
	for _, item := range resources {
		if item.InstanceID == "" {
			continue
		}
		if d[item.Category] == nil {
			d[item.Category] = []*GroupResource{}
		}
		d[item.Category] = append(d[item.Category], item)
	}
	for category, y := range d {
		go aliyun.getMetricReport(getMetricReportParams{params, category, y, report, errs})
		break
	}
	resp := &MetricReportResponse{
		Report: map[string][]MetricReport{},
	}
	cursor := 0
	finishOne := func() {
		cursor++
		if cursor == len(resources) {
			quit <- 0
		}
	}
	select {
	case r := <-report:
		if resp.Report[r.Category] == nil {
			resp.Report[r.Category] = []MetricReport{}
		}
		resp.Report[r.Category] = append(resp.Report[r.Category], r)
		finishOne()
	case e := <-errs:
		resp.Errors = append(resp.Errors, e.Error())
		finishOne()
	case <-quit:
		break
	}
	return
}
