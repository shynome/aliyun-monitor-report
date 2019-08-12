package group

import (
	"net/http"

	"github.com/labstack/echo/v4"
	AliyunType "github.com/shynome/aliyun-monitor-report/aliyun"
	aliyun "github.com/shynome/aliyun-monitor-report/aliyun/echo"
)

// List group
func List(c echo.Context) (err error) {
	params := &AliyunType.GetGroupListParams{}
	if err = c.Bind(params); err != nil {
		return
	}
	aliyunInstance := c.(*aliyun.Context).GetAliyunInstance()
	res, err := aliyunInstance.GetGroupList(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
