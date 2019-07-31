package base

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shynome/aliyun-monitor-report/api/aliyun"
)

// Projects aliyun monitor support
func Projects(c echo.Context) (err error) {
	res, err := aliyun.Instance.GetBaseProjects()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
