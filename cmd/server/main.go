package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hughiep/tomato-payment-service/internal/config"
	"github.com/hughiep/tomato-payment-service/pkg/logger"
)

func App() {
	router := gin.Default()

	// Config
	config.New()
	// Logger
	logger.New()

	// Router
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// Middleware

	router.Run(":8080")
}
