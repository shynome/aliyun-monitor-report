package monitor

import (
	"net/http"

	"github.com/labstack/echo/v4"
	aliyun "github.com/shynome/aliyun-monitor-report/aliyun/echo"
)

// Namespaces of aliyun monitor
func Namespaces(c echo.Context) (err error) {
	aliyunInstance := c.(*aliyun.Context).GetAliyunInstance()
	res, err := aliyunInstance.GetMonitorNamespaces()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
