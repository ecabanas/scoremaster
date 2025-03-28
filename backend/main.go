package main

import (
	"log"
	"net/http"
	"scoremaster/backend/models"

	"github.com/gin-gonic/gin"
)

func main() {

	dbPath := "./scoremaster.db"

	// Initialize the database
	db, err := models.OpenDatabase(dbPath)
	if err != nil {
		log.Fatal(err) // Log the error if OpenDatabase fails
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Define routes
	r.GET("/scores", func(c *gin.Context) {
		scores, err := models.GetScores(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, scores)
	})

	r.POST("/scores", func(c *gin.Context) {
		var json struct {
			Name  string `json:"name"`
			Score int    `json:"score"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := models.InsertScore(db, json.Name, json.Score)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "score added"})
	})

	// Run the server
	r.Run(":8080")
}
