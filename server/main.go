package main

import (
	"strings"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{"current_time": time.Now().UTC().Format(time.RFC3339)})
	})

	r.POST("/process", func(c *gin.Context) {
		var input struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"original":  input.Message,
			"processed": strings.ToUpper(input.Message) + " [FELDOLGOZVA]",
			"status":    "OK",
		})
	})

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	r.Run(":8080")
}
