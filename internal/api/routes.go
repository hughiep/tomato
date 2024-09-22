package api

import "github.com/gin-gonic/gin"

func New() {
	// Router
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
}
