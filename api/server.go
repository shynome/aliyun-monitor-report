package api

import (
	"github.com/labstack/echo/v4"
)

// Server all
var Server = echo.New()

// API group
var API = Server.Group("/api")

// Register group
func Register(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return API.Group(prefix, m...)
}
