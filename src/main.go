package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := ":5000"

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
	}

	router.Run(PORT)
}
