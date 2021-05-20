package router

import ( 
	"github.com/gin-gonic/gin"
	"net/http"

	"fmt"
	"../api"
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

	engine.GET("/host", func(c *gin.Context) { 
                c.HTML(http.StatusOK, "dashboad.tmpl", gin.H {
                        "title": "Host",
                })
        })

	engine.POST("/api/tty", func(c *gin.Context) {
		h := c.PostForm("Hostname")
                u := c.PostForm("Username")
                p := c.PostForm("Password")
		status := api.TTY(h, u, p)
		url := "http://"+h+":8080"
		fmt.Print(url)
		if status == 0{
			c.Redirect(http.StatusSeeOther, url)
		} else {
                        c.String(http.StatusOK, "error")
		}
	})
}
