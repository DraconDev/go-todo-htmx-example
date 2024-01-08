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

	// myDB.AddCard("test card", []string{"tag1", "tag2"})
	// cards, err := myDB.GetAllCards()
	// if err != nil {
	// 	panic(err)
	// }
	// print(cards)

	r.Run(":8080")

}

func setupRoutes(r *gin.Engine) {

	r.GET("/", handler.IndexHandler)
	r.GET("/api/card/add", handler.AddCardHandler)
	r.GET("/hello", handler.HomeHandler)
	r.GET("/cards", handler.CardsHandler)
	r.GET("/login", handler.LoginHandler)
	r.GET("/register", handler.RegisterHandler)
	r.GET("/dashboard", handler.DashboardHandler)
}
