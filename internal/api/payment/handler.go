package payment

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"tomato/internal/api/user"
	"tomato/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/webhook"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	UserRepo *user.UserRepository
}

func NewPaymentHandler(db *gorm.DB) *PaymentHandler {
	return &PaymentHandler{
		UserRepo: user.NewUserRepository(db),
	}
}

func (h *PaymentHandler) CheckoutSession(c echo.Context) error {
	// TODO: Retrieve from user session
	userId := c.Param("user")
	user, err := h.UserRepo.GetUserByID(userId)

	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if user.StripeCustomerID == "" {
		// Create a new customer if one doesn't exist
		customerID := createCustomer(user.Email)
		// Save the customer ID to the user
		h.UserRepo.UpdateCustomerID(user.ID, customerID)
	}

	// Create a new checkout session
	checkoutUrl := createCheckoutSession(user.StripeCustomerID)
	return c.JSON(http.StatusOK, checkoutUrl)
}

func (h *PaymentHandler) Webhook(c echo.Context) error {
	req := c.Request()
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		zap.S().Errorf("Error reading request body: %v\n", err)
		c.Response().WriteHeader(http.StatusServiceUnavailable)
		return err
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Webhook error while parsing basic request. %v\n", err.Error())
		c.Response().WriteHeader(http.StatusBadRequest)
		return err
	}

	// Replace this endpoint secret with your endpoint's unique secret
	endpointSecret := "whsec_7a73f948de68da76d261200008611c4c55a7f75f7ce57321d8c8d08ce11c7502"
	signatureHeader := req.Header.Get("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		zap.S().Errorf("⚠️  Webhook signature verification failed. %v\n", err.Error())
		c.Response().WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return err
	}
	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}

		// Update user role
		h.UserRepo.UpdateUserRole(paymentIntent.Customer.ID, models.Premium)

	default:
		zap.S().Infof("Unhandled event type: %s\n", event.Type)
	}

	// Handle webhook
	return c.JSON(http.StatusOK, "Webhook received")
}
