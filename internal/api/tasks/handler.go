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
	tasks := h.Repository.GetTasks()
	tasksResponse := make([]TaskResponse, 0)
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
		})
	}

	return c.JSON(http.StatusOK, tasksResponse)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	task := h.Repository.GetTaskByID(id)

	taskReponse := TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
	}

	return c.JSON(http.StatusOK, taskReponse)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	task := new(TaskRequest)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Message: "Invalid request body",
			Errors:  err.Error(),
		})
	}

	// Create
	id := h.Repository.CreateTask(*task)
	return c.JSON(http.StatusCreated, id)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	task := new(TaskRequest)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Message: "Invalid request body",
			Errors:  err.Error(),
		})
	}

	// Update
	h.Repository.UpdateTask(id, *task)
	return c.JSON(http.StatusOK, id)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	h.Repository.DeleteTask(id)
	return c.NoContent(http.StatusNoContent)
}
