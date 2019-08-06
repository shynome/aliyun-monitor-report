package aliyun

import (
	"testing"
)

func TestGetMetricReport(t *testing.T) {

	groups, _ := aliyun.GetGroupList(&GetGroupListParams{})

	html, err := aliyun.GetMetricReport(&GetMetricReportParams{
		GroupID:   groups[0].GroupID,
		StartTime: "2019-07-21 00:00:00",
		EndTime:   "2019-08-01 00:00:00",
	})

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(html)

	return
}
