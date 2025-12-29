package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/time", func(c *gin.Context) {
		currentTime := time.Now().Format(time.RFC3339)
		c.JSON(http.StatusOK, gin.H{
			"current_time": currentTime,
		})
	})

	r.Run(":8080")
}
