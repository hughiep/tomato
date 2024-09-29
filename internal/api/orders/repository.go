package orders

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

// func (r *OrderRepository) GetTasks() []string {
// 	db := db.GetDB()

// 	return []string{"task1", "task2", "task3"}
// }

func (r *OrderRepository) GetTaskByID(id int) string {
	return "task1"
}

func (r *OrderRepository) CreateOrder(task string) string {
	// Create a new order

	return "task1"
}

func (r *OrderRepository) UpdateTask(id int, task string) string {
	return "task1"
}

func (r *OrderRepository) DeleteTask(id int) string {
	return "task1"
}
