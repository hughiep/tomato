package tasks

import (
	"tomato/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) GetTasks() []models.Task {
	var tasks []models.Task
	r.DB.Find(&tasks)

	return tasks
}

func (r *TaskRepository) GetTaskByID(id string) models.Task {
	var task models.Task
	r.DB.First(&task, id)

	return task
}

func (r *TaskRepository) CreateTask(task models.Task) uint {
	r.DB.Create(&task)

	return task.ID
}

func (r *TaskRepository) UpdateTask(id string, task models.Task) {
	r.DB.Save(&task)
}

func (r *TaskRepository) DeleteTask(id string) {
	var task models.Task
	r.DB.Delete(&task, id)
}
