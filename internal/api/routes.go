package api

import (
	"github.com/hughiep/tomato-payment-service/internal/api/tasks"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// Router
	router := echo.New()

	// Routes
	r := router.Group("/v1") // API version
	tasks.Router(r)

	return router
}
