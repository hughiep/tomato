package payment

import (
	"log"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

func init() {
	// This is your test secret API key.
	stripe.Key = "sk_test_51Q4Kk4P4V1eL6Gp1l8EKyg3LTg61lsxY0ZKPO6j0eGTwQtdNxQmbhBWeBUwv3oY29tttl5SpiWNKhi56dzi0uheh00R8EGYYxI"
}

func createCheckoutSession() string {
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1Q4KltP4V1eL6Gp1sHJ7jNoS"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(domain + "/success"),
		CancelURL:  stripe.String(domain + "/cancel"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	return s.URL
}
