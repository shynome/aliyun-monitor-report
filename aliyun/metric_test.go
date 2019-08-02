package aliyun

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGetMetricList(t *testing.T) {

	groups, _ := aliyun.GetGroupList(&GetGroupListParams{})

	resources, _ := aliyun.GetGroupResources(&GetGroupResourcesParams{GroupID: groups[0].GroupID})

	resource := resources[0]

	dimensions := []Dimension{
		Dimension{
			InstanceID: resource.InstanceID,
		},
	}
	dimensionsBytes, err := json.Marshal(dimensions)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(dimensionsBytes)
	res, err := aliyun.GetMetricList(&GetMetricListParams{
		Dimensions: string(dimensionsBytes),
		Namespace:  "acs_" + strings.ToLower(resource.Category),
		MetricName: "CPUUtilization",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	return
}
