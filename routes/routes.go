package routes

import ( 
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(engine *gin.Engine) {
	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")

	engine.GET("/overview", func(c *gin.Context) {
                c.HTML(http.StatusOK, "dashboad.tmpl", gin.H {
                        "title": "Overview",
                })
        })
	engine.GET("/dashboad", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboad.tmpl", gin.H {
			"title": "Dashboad",
		})
	})
}
