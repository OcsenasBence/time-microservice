package main

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"current_time": time.Now().UTC().Format(time.RFC3339)})
	})

	r.POST("/process", func(c *gin.Context) {
		var input struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Could not parse/bind input message",
				"details": err.Error(),
			})
			return
		}

		if strings.TrimSpace(input.Message) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Validation Error: 'message' field cannot be empty",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"original":  input.Message,
			"processed": strings.ToUpper(input.Message) + " [FELDOLGOZVA]",
			"status":    "OK",
		})
	})

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	r.Run(":8080")
}