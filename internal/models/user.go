package models

type UserRole string

// Role enum
const (
	Free    UserRole = "free"
	Premium UserRole = "premium"
)

type User struct {
	ID               uint     `json:"id" gorm:"primaryKey"`
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Role             UserRole `json:"role"`
	StripeCustomerID string   `json:"stripe_customer_id"`
}
