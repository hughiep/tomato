package payment

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(r *echo.Group, db *gorm.DB) {
	handler := NewPaymentHandler(db)

	// Routes
	r.GET("/checkout/:user", handler.CheckoutSession)
	r.POST("/webhook", handler.Webhook)
}
