package aliyun

import (
	"testing"
)

func TestGetGroupList(t *testing.T) {
	keyword := "nnnnn"
	res, err := aliyun.GetGroupList(&GetGroupListParams{RegionID: "cn-hangzhou", Keyword: keyword})
	if err != nil {
		t.Error(err)
		return
	}
	if len(res.Resources.Resource) != 0 {
		t.Errorf("keyword %v should find 0 resource", keyword)
		return
	}
	t.Log(res)
	return
}
