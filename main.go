package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Server.Addr = ":8080"

	e.StartServer(e.Server)
}
