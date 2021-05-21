package main

import (
	"github.com/gin-gonic/gin"

	"github.com/DevelopNaoki/cvc/router"
)

func main() {

	engine:= gin.Default()
	engine.LoadHTMLGlob("src/templates/*.tmpl")

	router.Router(engine)

	engine.Run(":3000")
}
