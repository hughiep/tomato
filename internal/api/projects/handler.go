package projects

import (
	"net/http"
	"tomato/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProjectHandler struct {
	Repository *ProjectRepository
}

func NewProjectHandler(db *gorm.DB) *ProjectHandler {
	return &ProjectHandler{
		Repository: NewProjectRepository(db),
	}
}

func (h *ProjectHandler) GetProjects(c echo.Context) error {
	// Bind
	projects := h.Repository.GetProjects()

	return c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetProjectByID(c echo.Context) error {
	// Bind
	id := c.Param("id")
	project := h.Repository.GetProjectByID(id)

	return c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) CreateProject(c echo.Context) error {
	// Bind
	project := new(models.Project)
	if err := c.Bind(project); err != nil {
		return err
	}

	// Create
	id := h.Repository.CreateProject(*project)

	return c.JSON(http.StatusCreated, id)
}

func (h *ProjectHandler) UpdateProject(c echo.Context) error {
	// Bind
	id := c.Param("id")
	project := new(models.Project)
	if err := c.Bind(project); err != nil {
		return err
	}

	// Update
	h.Repository.UpdateProject(id, *project)

	return c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) DeleteProject(c echo.Context) error {
	// Bind
	id := c.Param("id")

	// Delete
	h.Repository.DeleteProject(id)

	return c.JSON(http.StatusOK, id)
}
