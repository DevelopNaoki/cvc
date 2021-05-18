package main

import (
	"github.com/gin-gonic/gin"

	"./routes"
)

func main() {

	engine:= gin.Default()
	engine.LoadHTMLGlob("src/templates/*.tmpl")

	routes.Router(engine)

	engine.Run(":3000")
}
