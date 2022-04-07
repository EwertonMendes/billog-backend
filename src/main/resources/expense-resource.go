package resources

import (
	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup) {
	expense := v1.Group("/expenses")
	{
		expense.GET("/create", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
	}
}
