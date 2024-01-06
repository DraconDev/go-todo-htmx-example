// basic https server
package main

import (
	handler "gokanban/internal/handlers"
	"net/http"
)

func main() {
	setupRoutes()

	http.ListenAndServe(":8080", nil)

}

func setupRoutes() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/dashboard", handler.DashboardHandler)
}
