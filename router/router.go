package router

import ( 
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	e.File("/dashboad", "src/html/dashboad.html")
	e.File("/host", "src/html/host.html")
	e.File("/volume", "src/html/volume.html")
	e.File("/network", "src/html/network.html")
	e.File("/virtualmachine", "src/html/virtualmachine")
	e.File("/container", "src/html/container.html")
	e.File("/setting", "src/html/setting.html")
	e.File("/test", "src/html/test.html")
}
