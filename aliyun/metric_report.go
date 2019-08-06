package aliyun

import (
	"fmt"

	"encoding/json"

	"github.com/thoas/go-funk"
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
	Max         []string
	Avg         []string
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
		{"内存", "mem.usedutilization"},
		{"连接数", "cpu_total"},
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

func (aliyun *Aliyun) getMetricReport(resource []*GroupResource, report chan MetricReport, err chan error, quit chan int) {
	dimensions := funk.Map(resource, func(r *GroupResource) Dimension {
		return Dimension{InstanceID: r.InstanceID}
	}).([]Dimension)
	dimensionsStr, jsonError := json.Marshal(dimensions)
	if jsonError != nil {
		err <- jsonError
		return
	}
	max, avg := make(chan []Datapoint), make(chan []Datapoint)
	maxData, avgData := <-max, <-avg
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
	for _, y := range d {
		go aliyun.getMetricReport(y, report, errs, quit)
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
