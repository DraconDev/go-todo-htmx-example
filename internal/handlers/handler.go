package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	// Serve the HTMX index page
	// return json
	jsonData := map[string]string{"message": "Hello, World!"}
	c.JSON(http.StatusOK, jsonData)

}

func Index(c *gin.Context) {

	c.HTML(200, "index.tmpl", gin.H{
		"title": "Home Page",
	})
}

func CardsHandler(c *gin.Context) {
	// Serve the HTMX cards page
	// return index html
	// return json

	// c.HTML(200, "index.gohtml", gin.H{
	// 	"title": "Home Page",
	// })

}

func LoginHandler(c *gin.Context) {
	// Handle user login
}

func RegisterHandler(c *gin.Context) {
	// Handle new user registration
}

func DashboardHandler(c *gin.Context) {
	// Serve the dashboard only if the user is authenticated
}
