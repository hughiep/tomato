package user

import (
	"errors"
	"fmt"
	"tomato/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("User not found")
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *UserRepository) UpdateUserRole(customerId string, role models.UserRole) {
	var user models.User
	r.DB.Where("stripe_customer_id = ?", customerId).First(&user)
	user.Role = role
	r.DB.Save(&user)
}

func (r *UserRepository) UpdateCustomerID(id uint, customerID string) {
	var user models.User
	r.DB.First(&user, id)
	user.StripeCustomerID = customerID
	r.DB.Save(&user)
}
