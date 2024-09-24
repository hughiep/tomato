package cmd

import (
	"fmt"

	"github.com/hughiep/tomato-payment-service/internal/api"
	"github.com/hughiep/tomato-payment-service/internal/config"
	"github.com/hughiep/tomato-payment-service/internal/db"
)

func App() {

	// Config
	c := config.Load()
	// Logger

	// Router
	router := api.Init()

	// Database
	db.Connect(c)

	// Middleware

	fmt.Printf(":%s", c.App.Port)
	router.Start(fmt.Sprintf(":%s", c.App.Port))
}
