package server

import (
	"fmt"

	"tomato/internal/api"
	"tomato/internal/config"
	"tomato/internal/db"
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
