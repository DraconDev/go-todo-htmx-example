// basic https server
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gokanban/internal/db"
	handler "gokanban/internal/handlers"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.LoadHTMLGlob("templates/partial/*")
	r.LoadHTMLGlob("templates/base/*")
	r.Static("/styles", "./styles")

	godotenv.Load(".env")
	connStr := os.Getenv("NEON_LINK")

	myDB, err := db.NewDatabase(connStr)
	if err != nil {
		panic(err)
	}
	myDB.GetVersion()
	// myDB.AddCard("test card", []string{"tag1", "tag2"})
	// cards, err := myDB.GetAllCards()
	// if err != nil {
	// 	panic(err)
	// }
	// print(cards)

	r.Run(":8080")

}

func setupRoutes(r *gin.Engine) {

	r.GET("/", handler.Index)
	r.GET("/hello", handler.HomeHandler)
	r.GET("/cards", handler.CardsHandler)
	r.GET("/login", handler.LoginHandler)
	r.GET("/register", handler.RegisterHandler)
	r.GET("/dashboard", handler.DashboardHandler)
}
