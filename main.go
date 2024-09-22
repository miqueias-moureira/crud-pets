package main

import (
	"crud-pets/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Root(e)

	e.Logger.Fatal(e.Start(":3000"))
}
