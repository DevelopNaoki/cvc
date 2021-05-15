package router

import ( 
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.Static("/css", "web/css")

	e.File("/", "web/html/index.html")
}
