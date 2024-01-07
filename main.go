// basic https server
package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"gokanban/internal/db"
	handler "gokanban/internal/handlers"
)

func main() {
	godotenv.Load(".env")
	connStr := os.Getenv("NEON_LINK")
	myDB, err := db.NewDatabase(connStr)
	if err != nil {
		panic(err)
	}

	setupRoutes()

	myDB.GetVersion()
	// myDB.AddCard("test card", []string{"tag1", "tag2"})
	cards, err := myDB.GetAllCards()
	if err != nil {
		panic(err)
	}
	print(cards)

	http.ListenAndServe(":8080", nil)

}

func setupRoutes() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/dashboard", handler.DashboardHandler)
}
