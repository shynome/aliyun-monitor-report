package aliyun

import (
	"testing"
)

func TestGetGroupList(t *testing.T) {
	keyword := "nnnnn"
	res, err := aliyun.GetGroupList(&GetGroupListParams{CommonParams{"cn-hangzhou"}, keyword})
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

func TestGetGroupDetails(t *testing.T) {
	groupList, err := aliyun.GetGroupList(&GetGroupListParams{})
	if err != nil {
		t.Error(err)
		return
	}
	resources := groupList.Resources.Resource
	if len(resources) == 0 {
		t.Error("can't find group")
		return
	}
	resource := resources[0]
	groupDetails, err := aliyun.GetGroupDetails(&GetGroupDetailsParams{GroupID: resource.GroupId})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(groupDetails)
	return
}
