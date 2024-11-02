package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TopicInfo represents the structure for topic information
type TopicInfo struct {
	Topic       string `json:"topic"`
	Timing      string `json:"timing"`
	Description string `json:"description"`
}

// Middleware to log requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request path
		c.Next()
	}
}

func submitHandler(c *gin.Context) {
	var topicInfos []TopicInfo
	if err := c.ShouldBindJSON(&topicInfos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process or store the information here
	// For now, just print it
	for _, info := range topicInfos {
		c.String(http.StatusOK, "Topic: %s, Timing: %s, Description: %s\n", info.Topic, info.Timing, info.Description)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Information submitted successfully!"})
}

func main() {
	router := gin.Default()

	// Apply the middleware
	router.Use(Logger())

	// Define the route for submitting the form
	router.POST("/submit", submitHandler)

	// Serve static files (for the HTML page)
	router.StaticFile("/", "./index1.html")

	// Start the server
	router.Run(":1233")
}
