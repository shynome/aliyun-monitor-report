package group

import (
	"net/http"

	"github.com/labstack/echo/v4"
	AliyunType "github.com/shynome/aliyun-monitor-report/aliyun"
	"github.com/shynome/aliyun-monitor-report/api/aliyun"
)

// Resources group
func Resources(c echo.Context) (err error) {
	params := &AliyunType.GetGroupResourcesParams{}
	if err = c.Bind(params); err != nil {
		return
	}
	aliyunInstance := c.(*aliyun.Context).GetAliyunInstance()
	res, err := aliyunInstance.GetGroupResources(params)
	c.JSON(http.StatusOK, res)
	return
}
