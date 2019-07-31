package monitor

import (
	"github.com/shynome/aliyun-monitor-report/api"
)

func init() {
	g := api.Register("/monitor")
	g.Any("/namespaces", Namespaces)
}
