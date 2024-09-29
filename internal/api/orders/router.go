package orders

import (
	"github.com/labstack/echo/v4"
)

func Router(r *echo.Group) {
	handler := NewTaskHandler()

	// Routes
	r.GET("/tasks", handler.GetTasks)
}
