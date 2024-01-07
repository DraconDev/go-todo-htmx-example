// basic https server
package main

import (
	"net/http"

	"gokanban/internal/db"
	handler "gokanban/internal/handlers"
)

func main() {
	connStr := "postgresql://DraconDev:s8JrG9azEWXN@ep-summer-credit-29657807.us-east-2.aws.neon.tech/Main?sslmode=require"
	myDB, err := db.NewDatabase(connStr)
	if err != nil {
		panic(err)
	}

	setupRoutes()

	myDB.GetVersion()
	myDB.AddCard("test card", []string{"tag1", "tag2"})
	// cards, err := myDB.GetAllCards()
	// if err != nil {
	// 	panic(err)
	// }
	// print(cards)

	http.ListenAndServe(":8080", nil)

}

func setupRoutes() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/dashboard", handler.DashboardHandler)
}
