package aliyun

import (
	"testing"
)

func TestGetGroupList(t *testing.T) {
	keyword := "nnnnn"
	groupList, err := aliyun.GetGroupList(&GetGroupListParams{CommonParams{"cn-hangzhou"}, keyword})
	if err != nil {
		t.Error(err)
		return
	}
	if len(groupList) != 0 {
		t.Errorf("keyword %v should find 0 resource", keyword)
		return
	}
	t.Log(groupList)
	return
}

func TestGetGroupResources(t *testing.T) {
	groupDetailsList, err := aliyun.GetGroupList(&GetGroupListParams{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(groupDetailsList) == 0 {
		t.Error("can't find group")
		return
	}
	resource := groupDetailsList[0]
	groupDetails, err := aliyun.GetGroupResources(&GetGroupResourcesParams{GroupID: resource.GroupID})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(groupDetails)
	return
}
