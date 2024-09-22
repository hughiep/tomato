package cmd

import (
	"fmt"

	"github.com/hughiep/tomato-payment-service/internal/api"
	"github.com/hughiep/tomato-payment-service/internal/config"
)

func App() {

	// Config
	c := config.Load()
	// Logger

	// Router
	router := api.Init()

	// Middleware

	fmt.Printf(":%s", c.App.Port)
	router.Start(fmt.Sprintf(":%s", c.App.Port))
}
