package handler

import (
	"net/http"
	"text/template"

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

	// // Set appropriate response headers
	// c.Header("Content-Type", "application/json")

	// // Return JSON response using Gin's JSON function
	// c.JSON(http.StatusOK, cards)

	tmpl, err := template.ParseFiles("templates/block/cards.html") // Assuming cards.html contains your template
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tmpl.Execute(c.Writer, cards) // Render directly to the response writer
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func IndexHandler(c *gin.Context) {
	data := map[string]string{"Title": "Kanban Board - Index"}
	c.HTML(http.StatusOK, "layout.html", data)
}

func DeleteCardHandler(c *gin.Context) {
	// Handle card deletion
	db := db.AccessDB()
	db.DeleteCard(c.Param("id"))

}
