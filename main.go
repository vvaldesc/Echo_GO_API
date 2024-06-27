package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Echo/routes"
)

func main() {
	e := echo.New()
	routes.Routes(e)
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${time_rfc3339} ${status} ${method} ${path}\n${body}\n",
		},
	))
	e.Start(":8080")
}
