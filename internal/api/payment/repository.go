package payment

import "tomato/internal/db"

type PaymentRepository struct{}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

var _db = db.GetDB()

// func (r *PaymentRepository) GetTasks() []string {
// 	db := db.GetDB()

// 	return []string{"task1", "task2", "task3"}
// }

func (r *PaymentRepository) CreatePayment(user string, product string) {
	// Create a new payment in the database

	_db.Create(&Payment{
		User:    "user1",
		Product: "product1",
		Status:  PaymentStatusPending,
	})
}

func (r *PaymentRepository) UpdateTask(id int, task string) string {
	return "task1"
}

func (r *PaymentRepository) DeleteTask(id int) string {
	return "task1"
}
