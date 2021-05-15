package router

import ( 
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.File("/", "web/html/index.html")
}
