package tasks

type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PomodoroEst uint   `json:"pomodoro_est"`
	Status      string `json:"status"`
	Note        string `json:"note"`
	ProjectID   uint   `json:"project_id"`
}
