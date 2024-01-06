// basic https server
package main

import (
	"fmt"
	"gokanban/route"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func setupRoutes() {
	http.HandleFunc("/", route.HomeHandler) // Serve HTMX front-end
	http.HandleFunc("/login", route.LoginHandler)
	http.HandleFunc("/register", route.RegisterHandler)
	http.HandleFunc("/dashboard", route.DashboardHandler) // Make sure this is protected
	// Add other necessary routes
}
