package tasks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	Repository *TaskRepository
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		Repository: NewTaskRepository(),
	}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	// Bind
	tasks := h.Repository.GetTasks()

	return c.JSON(http.StatusOK, tasks)
}
