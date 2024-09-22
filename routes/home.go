package routes

import (
	"crud-pets/handlers"

	"github.com/labstack/echo/v4"
)

func Root(root *echo.Echo) *echo.Group {

	common := root.Group("")

	common.GET("/", handlers.HomeHandler("name"))

	return common
}
