package payment

import (
	"github.com/labstack/echo/v4"
)

func Router(r *echo.Group) {
	handler := NewPaymentHandler()

	// Routes
	r.GET("/checkout", handler.CheckoutSession)
	r.POST("/webhook", handler.Webhook)
}
