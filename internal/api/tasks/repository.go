package tasks

import (
	"tomato/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

type TaskRepositoryInterface interface {
	GetTasks() []models.Task
	GetTaskByID(id string) models.Task
	CreateTask(task TaskRequest) uint
	UpdateTask(id string, task TaskRequest)
	DeleteTask(id string)
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) GetTasks() []models.Task {
	var tasks []models.Task
	// khi thao tác với db nên handle error
	result := r.DB.Find(&tasks)
	if result.Error != nil {
		// return nil, result.Error
	}

	return tasks
}

func (r *TaskRepository) GetTaskByID(id string) models.Task {
	var task models.Task
	r.DB.First(&task, id)

	return task
}

func (r *TaskRepository) CreateTask(task TaskRequest) uint {
	tx := r.DB.Create(&models.Task{
		Title:       task.Title,
		Description: task.Description,
		PomodoroEst: task.PomodoroEst,
		Status:      models.TaskStatus(task.Status),
		Note:        task.Note,
		ProjectID:   task.ProjectID,
	})

	return uint(tx.RowsAffected)
}

func (r *TaskRepository) UpdateTask(id string, task TaskRequest) {
	var taskModel models.Task
	r.DB.First(&taskModel, id)

	taskModel.Title = task.Title
	taskModel.Description = task.Description
	taskModel.PomodoroEst = task.PomodoroEst
	taskModel.Status = models.TaskStatus(task.Status)
	taskModel.Note = task.Note
	taskModel.ProjectID = task.ProjectID

	r.DB.Save(&taskModel)
}

func (r *TaskRepository) DeleteTask(id string) {
	var task models.Task
	r.DB.Delete(&task, id)
}
