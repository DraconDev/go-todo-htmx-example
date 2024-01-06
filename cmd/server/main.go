// basic https server
package main

import (
	"gokanban/internal/route"
	"net/http"
)

func main() {
	setupRoutes()

	http.ListenAndServe(":8080", nil)

}

func setupRoutes() {
	http.HandleFunc("/", route.HomeHandler) // Serve HTMX front-end
	http.HandleFunc("/login", route.LoginHandler)
	http.HandleFunc("/register", route.RegisterHandler)
	http.HandleFunc("/dashboard", route.DashboardHandler) // Make sure this is protected
	// Add other necessary routes
}
