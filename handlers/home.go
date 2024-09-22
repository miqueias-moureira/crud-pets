package handlers

import (
	"crud-pets/views"

	"github.com/labstack/echo/v4"
)

func HomeHandler(name string) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		return views.Home(name).Render(ctx.Request().Context(), ctx.Response())
	}
}
