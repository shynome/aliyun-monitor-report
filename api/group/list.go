package group

import (
	"net/http"

	"github.com/labstack/echo/v4"
	AliyunType "github.com/shynome/aliyun-monitor-report/aliyun"
	"github.com/shynome/aliyun-monitor-report/api/aliyun"
)

// List group
func List(c echo.Context) (err error) {
	params := &AliyunType.GetGroupListParams{}
	if err = c.Bind(params); err != nil {
		return
	}
	res, err := aliyun.Instance.GetGroupList(params)
	c.JSON(http.StatusOK, res)
	return
}
