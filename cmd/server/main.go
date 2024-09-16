package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hughiep/tomato-payment-service/internal/config"
	"github.com/hughiep/tomato-payment-service/pkg/logger"
)

func main() {
	router := gin.Default()

	// Config
	appConfig := config.New()
	// Logger
	appLogger := logger.New()
	// Router
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// Middleware

	router.Run(":8080")
}
