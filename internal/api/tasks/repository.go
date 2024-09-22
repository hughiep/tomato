package tasks

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) GetTasks() []string {
	return []string{"task1", "task2", "task3"}
}

func (r *TaskRepository) GetTaskByID(id int) string {
	return "task1"
}

func (r *TaskRepository) CreateTask(task string) string {
	return "task1"
}

func (r *TaskRepository) UpdateTask(id int, task string) string {
	return "task1"
}

func (r *TaskRepository) DeleteTask(id int) string {
	return "task1"
}
