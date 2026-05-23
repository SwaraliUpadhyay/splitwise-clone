package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Register a route:
	router.GET("/health", healthCheck)

	// Start the server on port 8080
	router.Run(":8080")
}

// This is a "handler function" — it runs when /health is called
// c is the context — it holds the request info and lets you send a response
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running!",
	})
}
