package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./router"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano}: ${remote_ip}: \"${method} ${uri} ${status}\"\n",
	}))
	e.Use(middleware.Recover())

	router.Router(e)

	e.Start(":8080")
}
