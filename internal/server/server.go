package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"tomato/internal/api"
	"tomato/internal/config"
	"tomato/internal/db"
)

func App() {

	// Config
	c := config.Load()
	// Logger

	// Database
	db := db.Connect(c)

	// Router
	router := api.SetUpRouter(db)

	// Middleware

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := router.Start(fmt.Sprintf(":%s", c.App.Port)); err != nil {
			router.Logger.Info("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal(err)
	}
}
