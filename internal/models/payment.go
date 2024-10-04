package models

import "gorm.io/gorm"

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
)

type Payment struct {
	gorm.Model
	// Unexported struct fields are invisible to the JSON package.
	user    string        `json:"user"`
	product string        `json:"product"`
	status  PaymentStatus `json:"status"`
}
