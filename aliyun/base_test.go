package aliyun

import (
	"testing"
)

func TestGetBaseProjects(t *testing.T) {
	res, err := aliyun.GetBaseProjects()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
	return
}
