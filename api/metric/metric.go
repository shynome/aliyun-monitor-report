package metric

import (
	"github.com/shynome/aliyun-monitor-report/api"
)

func init() {
	g := api.Register("/metric")
	g.Any("/list", List)
	g.Any("/top", Top)
	g.Any("/report", Report)
}
