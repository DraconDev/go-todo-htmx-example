// basic https server
package main

import (
	"github.com/gin-gonic/gin"

	handler "gokanban/internal/handlers"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Static("/styles", "./styles")
	r.LoadHTMLGlob("templates/**/*")

	r.Run(":8080")

}

func setupRoutes(r *gin.Engine) {

	r.GET("/", handler.IndexHandler)
	r.POST("/api/card/add", handler.AddCardHandler)
	r.GET("/api/card/all", handler.GetBoard)
	r.GET("/hello", handler.HomeHandler)

}
