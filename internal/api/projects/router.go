package projects

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(r *echo.Group, db *gorm.DB) {
	handler := NewProjectHandler(db)

	// Routes
	r.GET("/projects", handler.GetProjects)
	r.GET("/projects/:id", handler.GetProjectByID)
	r.POST("/projects", handler.CreateProject)
	r.PUT("/projects/:id", handler.UpdateProject)
	r.DELETE("/projects/:id", handler.DeleteProject)
}
