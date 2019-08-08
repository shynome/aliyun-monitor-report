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
	ReportDimension
	Error error
	Max   float64
	Avg   float64
}

// MetricReport type
type MetricReport struct {
	*GroupResource
	Dimensions []*DimensionReport
	Error      error
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
		{"CPU 使用率", "CpuUsage"},
		{"内存使用率", "MemoryUsage"},
		{"连接数使用率", "ConnectionUsage"},
	},
	"KVSTORE": {
		{"CPU 使用率", "CpuUsage"},
		{"内存使用率", "MemoryUsage"},
		{"连接数使用率", "ConnectionUsage"},
	},
	"SLB": {
		{"流入带宽", "TrafficRXNew"},
		{"活跃连接数", "ActiveConnection"},
		{"并发连接数", "MaxConnection"},
	},
	// "CDN": {
	// 	{"宽带峰值", ""},
	// 	{"下行流量", ""},
	// 	{"QPS", ""},
	// },
}

type getMetricReportParams struct {
	*GetMetricReportParams
	category  string
	resources []*GroupResource
	report    chan MetricReport
}

func getNamespace(category string) string {
	category = "acs_" + strings.ToLower(category)
	switch category {
	case "acs_slb":
		category = "acs_slb_dashboard"
	case "acs_ecs":
		category = "acs_ecs_dashboard"
	case "acs_rds":
		category = "acs_rds_dashboard"
	case "acs_apigateway":
		category = "acs_apigateway_dashboard"
	case "acs_containerservice":
		category = "acs_containerservice_dashboard"
	case "acs_ehpc":
		category = "acs_ehpc_dashboard"
	case "acs_ess":
		category = "acs_ess_dashboard"
	case "acs_oss":
		category = "acs_oss_dashboard"
	case "acs_sls":
		category = "acs_sls_dashboard"
	}
	return category
}

var (
	orderByMax = "Maximum"
	orderByAvg = "Average"
)

func (aliyun *Aliyun) getMetricReport(params getMetricReportParams) {

	resp := map[string]map[string]*DimensionReport{}
	dimensions := make([]Dimension, len(params.resources))
	finish := func(err error) {
		if err != nil {
			for _, r := range params.resources {
				params.report <- MetricReport{
					GroupResource: r,
					Error:         err,
				}
			}
			return
		}
		for _, r := range params.resources {
			d := make([]*DimensionReport, len(dimensions))
			for _, v := range resp[r.InstanceID] {
				d = append(d, v)
			}
			params.report <- MetricReport{
				GroupResource: r,
				Dimensions:    d,
			}
		}
	}

	for i, r := range params.resources {
		dimensions[i] = Dimension{InstanceID: r.InstanceID}
		resp[r.InstanceID] = map[string]*DimensionReport{}
		for _, d := range DefaultReportDimensions[params.category] {
			resp[r.InstanceID][d.Name] = &DimensionReport{ReportDimension: d}
		}
	}

	dimensionsBytes, jsonError := json.Marshal(dimensions)
	if jsonError != nil {
		finish(jsonError)
		return
	}

	type PeakData struct {
		ReportDimension
		Datapoints []Datapoint
		Error      error
	}

	namespace := getNamespace(params.category)
	getPeakData := func(d ReportDimension, orderBy string) PeakData {
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
		peakData := PeakData{ReportDimension: d}
		if err != nil {
			peakData.Error = err
		} else {
			peakData.Datapoints = datapoints
		}
		return peakData
	}

	dealPeakData := func(orderBy string, peakData PeakData) {

		var setDimension func(v *DimensionReport, datapoint Datapoint)

		if orderBy != orderByMax && orderBy != orderByAvg {
			return
		}
		if orderBy == orderByMax {
			setDimension = func(v *DimensionReport, datapoint Datapoint) {
				if v.Max > datapoint.Maximum {
					return
				}
				v.Max = datapoint.Maximum
			}
		}
		if orderBy == orderByAvg {
			setDimension = func(v *DimensionReport, datapoint Datapoint) {
				if v.Avg > datapoint.Average {
					return
				}
				v.Avg = datapoint.Average
			}
		}

		for _, datapoint := range peakData.Datapoints {
			v := resp[datapoint.InstanceID][peakData.Name]
			if peakData.Error != nil {
				v.Error = peakData.Error
				continue
			}
			setDimension(v, datapoint)
		}
		return
	}

	dealMaxCount, dealAvgCount := 0, 0
	endCount := len(DefaultReportDimensions[params.category])
	dealMaxChan, dealAvgChan := make(chan int), make(chan int)
	dealDimension := func(orderBy string, d ReportDimension) {
		peakData := getPeakData(d, orderBy)
		dealPeakData(orderBy, peakData)
		if orderBy == orderByMax {
			dealMaxCount++
			if dealMaxCount == endCount {
				dealMaxChan <- 0
			}
		}
		if orderBy == orderByAvg {
			dealAvgCount++
			if dealAvgCount == endCount {
				dealAvgChan <- 0
			}
		}
	}

	for _, d := range DefaultReportDimensions[params.category] {
		go dealDimension(orderByMax, d)
		go dealDimension(orderByAvg, d)
	}

	<-dealMaxChan
	<-dealAvgChan

	finish(nil)
	return

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
	report := make(chan MetricReport, len(resources))
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
		if category != "ECS" {
			continue
		}
		go aliyun.getMetricReport(getMetricReportParams{params, category, y, report})
		break
	}
	resp := &MetricReportResponse{
		Report: map[string][]MetricReport{},
	}

	for cursor := 0; cursor < len(resources); cursor++ {
		r := <-report
		if resp.Report[r.Category] == nil {
			resp.Report[r.Category] = []MetricReport{}
		}
		resp.Report[r.Category] = append(resp.Report[r.Category], r)
	}

	fmt.Println(resp)

	return
}
