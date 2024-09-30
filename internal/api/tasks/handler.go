package tasks

import (
	"net/http"
	"tomato/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Repository *TaskRepository
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		Repository: NewTaskRepository(db),
	}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	// Bind
	tasks := h.Repository.GetTasks()

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	// Bind
	id := c.Param("id")
	task := h.Repository.GetTaskByID(id)

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	// Bind
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return err
	}

	// Create
	id := h.Repository.CreateTask(*task)

	return c.JSON(http.StatusCreated, id)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	// Bind
	id := c.Param("id")
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return err
	}

	// Update
	h.Repository.UpdateTask(id, *task)

	return c.JSON(http.StatusOK, "task1")
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	// Bind
	id := c.Param("id")

	// Delete
	h.Repository.DeleteTask(id)

	return c.JSON(http.StatusOK, "task1")
}
