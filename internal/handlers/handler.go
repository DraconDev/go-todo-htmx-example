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
	cardTitle := c.PostForm("cardTitle")
	db := db.AccessDB()
	db.AddCard(cardTitle, []string{"tag1", "tag2"})
	// db.AddCard(c.PostForm("text"), c.PostFormArray("tags"))

}

func GetBoard(c *gin.Context) {
	db := db.AccessDB()

	cards, err := db.GetAllCards()
	if err != nil {
		panic(err)
	}
	print(cards)

	// Set appropriate response headers
	c.Header("Content-Type", "application/json")

	// Return JSON response using Gin's JSON function
	c.JSON(http.StatusOK, cards)

}

func IndexHandler(c *gin.Context) {
	data := map[string]string{"Title": "Kanban Board - Index"}
	c.HTML(http.StatusOK, "layout.html", data)
}

func CardsHandler(c *gin.Context) {
	data := map[string]string{"Title": "Kanban Board - Index"}
	c.HTML(http.StatusOK, "index.html", data)
}
