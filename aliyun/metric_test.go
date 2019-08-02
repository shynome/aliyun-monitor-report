package aliyun

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGetMetricList(t *testing.T) {

	groups, _ := aliyun.GetGroupList(&GetGroupListParams{})

	resources, _ := aliyun.GetGroupResources(&GetGroupResourcesParams{GroupID: groups[0].GroupID})

	dimensions := []Dimension{}

	for _, resource := range resources {
		if resource.Category != "ECS" {
			continue
		}
		dimensions = append(dimensions, Dimension{InstanceID: resource.InstanceID})
	}

	dimensionsBytes, err := json.Marshal(dimensions)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := aliyun.GetMetricList(&GetMetricListParams{
		Dimensions: string(dimensionsBytes),
		Namespace:  "acs_" + strings.ToLower("ECS"),
		MetricName: "CPUUtilization",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	return
}
