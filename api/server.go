package api

import (
	"github.com/labstack/echo/v4"
)

// Server api
var Server = echo.New()

func Register(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return Server.Group(prefix, m...)
}
