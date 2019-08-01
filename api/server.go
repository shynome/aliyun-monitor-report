package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shynome/aliyun-monitor-report/api/aliyun"
)

// Server all
var Server = echo.New()

// API group
var API = Server.Group("/api", func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &aliyun.Context{Context: c}
		return h(cc)
	}
})

// Register group
func Register(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return API.Group(prefix, m...)
}
