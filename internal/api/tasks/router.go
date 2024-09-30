package tasks

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(r *echo.Group, db *gorm.DB) {
	handler := NewTaskHandler(db)

	// Routes
	r.GET("/tasks", handler.GetTasks)
	r.GET("/tasks/:id", handler.GetTaskByID)
	r.POST("/tasks", handler.CreateTask)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
}
