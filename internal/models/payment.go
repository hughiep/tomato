package models

import "gorm.io/gorm"

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
)

type Payment struct {
	gorm.Model
	user    string        `json:"user"`
	product string        `json:"product"`
	status  PaymentStatus `json:"status"`
}
