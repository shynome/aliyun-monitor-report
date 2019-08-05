package aliyun

import (
	"github.com/shynome/aliyun-monitor-report/aliyun"
)

// CommonParams aliyun request
type CommonParams struct {
	RegionID string `json:"regionID"`
}

// Instance aliyun
var Instance = aliyun.NewWithEnv()
