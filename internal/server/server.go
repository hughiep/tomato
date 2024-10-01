package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"tomato/internal/api"
	"tomato/internal/config"
	"tomato/internal/db"
	"tomato/internal/middlewares"
	"tomato/pkg/logger"

	builtinMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Serve() {
	// Config
	c := config.Load()

	// Logger
	if err := logger.Init(); err != nil {
		panic(err)
	}

	// Database
	db := db.Connect(c)

	// Router
	router := api.SetUpRouter(db)

	// Middleware
	router.Use(builtinMiddleware.CORS())
	router.Use(middlewares.ZapLogger(zap.L()))
	router.Use(builtinMiddleware.Recover())

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := router.Start(fmt.Sprintf(":%s", c.App.Port)); err != nil {
			fmt.Println(err)
			zap.L().Info("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server forced to shutdown:", zap.Error(err))
	}
}
