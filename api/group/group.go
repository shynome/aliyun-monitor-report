package group

import (
	"github.com/shynome/aliyun-monitor-report/api"
)

func init() {
	g := api.Register("/group")
	g.Any("/list", List)
	g.Any("/resources", Resources)
}
