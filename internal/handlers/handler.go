package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gokanban/internal/db"
)

func HomeHandler(c *gin.Context) {
	// Serve the HTMX index page
	// return json
	jsonData := map[string]string{"message": "Hello, World!"}
	c.JSON(http.StatusOK, jsonData)

}

func AddCardHandler(c *gin.Context) {
	// Handle card creation
	db := db.AccessDB()
	db.AddCard("test card", []string{"tag1", "tag2"})
	// db.AddCard(c.PostForm("text"), c.PostFormArray("tags"))

}

func GetBoard(c *gin.Context) {

}

func IndexHandler(c *gin.Context) {
	data := map[string]string{"Title": "Kanban Board - Index"}
	c.HTML(http.StatusOK, "layout.html", data)
}

func CardsHandler(c *gin.Context) {
	data := map[string]string{"Title": "Kanban Board - Index"}
	c.HTML(http.StatusOK, "index.html", data)
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
