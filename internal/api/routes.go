package api

import (
	"tomato/internal/api/payment"
	"tomato/internal/api/tasks"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *echo.Echo {
	// Router
	router := echo.New()

	// Routes
	r := router.Group("/api/v1") // API version

	// Health check
	r.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	tasks.Router(r, db)
	payment.Router(r, db)

	return router
}
