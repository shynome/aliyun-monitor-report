package monitor

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shynome/aliyun-monitor-report/api/aliyun"
)

// Namespaces of aliyun monitor
func Namespaces(c echo.Context) (err error) {
	aliyunInstance := c.(*aliyun.Context).GetAliyunInstance()
	res, err := aliyunInstance.GetMonitorNamespaces()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
