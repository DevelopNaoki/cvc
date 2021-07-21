package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DevelopNaoki/cvc/api"
)

func Router(engine *gin.Engine) {
	sidevar := [...]string{
		"Overview",
		"Dashboard",
		"Network",
		"Virtual Machie",
		"Container",
		"Setting",
	}

	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")

	engine.GET("/overview", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboad.tmpl", gin.H{
			"title":   "Overview",
			"sidevar": sidevar,
		})
	})
	engine.GET("/dashboad", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboad.tmpl", gin.H{
			"title": "Dashboad",
		})
	})

	engine.GET("/host", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboad.tmpl", gin.H{
			"title": "Host",
		})
	})

	engine.POST("/api/tty", func(c *gin.Context) {
		h := c.PostForm("Hostname")
		u := c.PostForm("Username")
		p := c.PostForm("Password")
		status := api.TTY(h, u, p)
		if status == 0 {
			c.Redirect(http.StatusSeeOther, "http://192.168.0.50:8080")
		} else {
			c.String(http.StatusOK, "error")
		}
	})
}
