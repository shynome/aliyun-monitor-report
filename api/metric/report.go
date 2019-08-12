package metric

import (
	"net/http"

	"github.com/labstack/echo/v4"
	AliyunType "github.com/shynome/aliyun-monitor-report/aliyun"
	aliyun "github.com/shynome/aliyun-monitor-report/aliyun/echo"
)

// Report metric
func Report(c echo.Context) (err error) {
	params := &AliyunType.GetMetricReportParams{}
	if err = c.Bind(params); err != nil {
		return
	}
	aliyunInstance := c.(*aliyun.Context).GetAliyunInstance()
	res, err := aliyunInstance.GetMetricReport(params)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
