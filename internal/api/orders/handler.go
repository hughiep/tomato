package orders

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Repository *TaskRepository
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		Repository: NewTaskRepository(),
	}
}

func (h *OrderHandler) GetTasks(c echo.Context) error {
	// Bind
	tasks := h.Repository.GetTasks()

	return c.JSON(http.StatusOK, tasks)
}
