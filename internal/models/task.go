package models

import "gorm.io/gorm"

type TaskStatus string

// Status enum
const (
	Open       TaskStatus = "open"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
)

type Task struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PomodoroEst uint       `json:"pomodoro_est"`
	PomodoroAct uint       `json:"pomodoro_act"`
	Status      TaskStatus `json:"status"`
	Note        string     `json:"note"`
	UserID      uint       `json:"user_id"`
	ProjectID   uint       `json:"project_id"`
}
