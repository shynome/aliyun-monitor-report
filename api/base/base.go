package base

import (
	"github.com/shynome/aliyun-monitor-report/api"
)

func init() {
	g := api.Register("/base")
	g.Any("/projects", Projects)
}
