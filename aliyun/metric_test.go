package aliyun

import (
	"encoding/json"
	"testing"
)

func getECSDimensions() string {

	groups, _ := aliyun.GetGroupList(&GetGroupListParams{})

	resources, _ := aliyun.GetGroupResources(&GetGroupResourcesParams{GroupID: groups[0].GroupID})

	dimensions := []Dimension{}

	for _, resource := range resources {
		if resource.Category != "ECS" {
			continue
		}
		dimensions = append(dimensions, Dimension{resource.InstanceID})
		break
	}

	dimensionsBytes, _ := json.Marshal(dimensions)

	return string(dimensionsBytes)

}

func TestGetMetricList(t *testing.T) {

	dimensions := getECSDimensions()

	res, err := aliyun.GetMetricList(&GetMetricListParams{
		Dimensions: dimensions,
		Namespace:  "acs_ecs",
		MetricName: "CPUUtilization",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	return
}

func TestGetMetricTop(t *testing.T) {

	dimensions := getECSDimensions()

	res, err := aliyun.GetMetricTop(&GetMetricTopParams{
		GetMetricListParams: GetMetricListParams{
			Dimensions: dimensions,
			Namespace:  "acs_ecs",
			MetricName: "CPUUtilization",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	return
}
