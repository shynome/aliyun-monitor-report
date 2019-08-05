package aliyun

import (
	"github.com/labstack/echo/v4"
	"github.com/shynome/aliyun-monitor-report/aliyun"
)

// Context aliyun
type Context struct {
	echo.Context
}

// GetAliyunInstance func
func (c *Context) GetAliyunInstance() *aliyun.Aliyun {
	return Instance
}
