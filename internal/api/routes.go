package api

import (
	"tomato/internal/api/payment"
	"tomato/internal/api/tasks"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// Router
	router := echo.New()

	// Routes
	r := router.Group("/api/v1") // API version

	// Health check
	r.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	tasks.Router(r)
	payment.Router(r)

	return router
}
