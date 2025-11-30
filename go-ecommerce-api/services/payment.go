package services

import (
	"encoding/json"
	"go-ecommerce-api/models"
	"go-ecommerce-api/repositories"
	"strconv"

	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/paymentintent"
	"github.com/stripe/stripe-go/v84/webhook"
)

type PaymentService struct {
	repo          *repositories.PaymentRepository
	orderRepo     repositories.OrderRepository
	stripeKey     string
	webhookSecret string
}

func NewPaymentService(
	repo *repositories.PaymentRepository,
	orderRepo repositories.OrderRepository,
	stripeKey string,
	webhookSecret string,
) *PaymentService {
	stripe.Key = stripeKey
	return &PaymentService{
		repo:          repo,
		orderRepo:     orderRepo,
		stripeKey:     stripeKey,
		webhookSecret: webhookSecret,
	}
}

// Create payment intent
func (s *PaymentService) CreatePaymentIntent(orderID uint, amount float64, currency string) (*stripe.PaymentIntent, *models.Payment, error) {
	// Convert amount to cents for Stripe
	amountInCents := int64(amount * 100)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCents),
		Currency: stripe.String(currency),
		Metadata: map[string]string{
			"order_id": strconv.FormatUint(uint64(orderID), 10),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, nil, err
	}

	// Save to database
	payment, err := s.repo.SavePaymentIntent(orderID, pi.ID, amount, currency, string(pi.Status))
	if err != nil {
		return nil, nil, err
	}

	return pi, payment, nil
}

// Confirm payment intent
func (s *PaymentService) ConfirmPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentConfirmParams{}
	return paymentintent.Confirm(paymentIntentID, params)
}

// Cancel payment intent
func (s *PaymentService) CancelPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentCancelParams{}
	return paymentintent.Cancel(paymentIntentID, params)
}

// Handle webhook events
func (s *PaymentService) HandleWebhook(payload []byte, signature string) error {
	event, err := webhook.ConstructEvent(payload, signature, s.webhookSecret)
	if err != nil {
		return err
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var pi stripe.PaymentIntent
		jsonData, err := json.Marshal(event.Data.Object)
		if err != nil {
			return err
		}
		err = json.Unmarshal(jsonData, &pi)
		if err != nil {
			return err
		}

		// Update payment status
		updates := map[string]interface{}{
			"status": string(pi.Status),
		}
		err = s.repo.UpdatePaymentDetails(pi.ID, updates)
		if err != nil {
			return err
		}

		// Update order status and payment info
		orderIDStr := pi.Metadata["order_id"]
		if orderIDStr != "" {
			orderID, _ := strconv.ParseUint(orderIDStr, 10, 32)
			orderUpdates := map[string]interface{}{
				"payment_status":    "succeeded",
				"payment_intent_id": pi.ID,
				"status":            "processing",
			}
			s.orderRepo.Update(uint(orderID), orderUpdates)
		}

	case "payment_intent.payment_failed":
		var pi stripe.PaymentIntent
		jsonData, err := json.Marshal(event.Data.Object)
		if err != nil {
			return err
		}
		err = json.Unmarshal(jsonData, &pi)
		if err != nil {
			return err
		}

		updates := map[string]interface{}{
			"status": "failed",
		}
		if pi.LastPaymentError != nil {
			updates["failure_reason"] = pi.LastPaymentError.Code
		}
		err = s.repo.UpdatePaymentDetails(pi.ID, updates)
		if err != nil {
			return err
		}

		// Update order payment status
		orderIDStr := pi.Metadata["order_id"]
		if orderIDStr != "" {
			orderID, _ := strconv.ParseUint(orderIDStr, 10, 32)
			s.orderRepo.Update(uint(orderID), map[string]interface{}{
				"payment_status": "failed",
			})
		}
	}

	return nil
}

// Get payment details
func (s *PaymentService) GetPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	return paymentintent.Get(paymentIntentID, nil)
}

// Get payment by order ID
func (s *PaymentService) GetPaymentByOrderID(orderIDStr string) (*models.Payment, error) {
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.repo.GetPaymentByOrderID(uint(orderID))
}

// Get user payments
func (s *PaymentService) GetUserPayments(userID uint) ([]models.Payment, error) {
	return s.repo.GetPaymentsByUserID(userID)
}

// Confirm payment success (for testing without webhooks)
func (s *PaymentService) ConfirmPaymentSuccess(orderID uint) error {
	// Update payment status
	payment, err := s.repo.GetPaymentByOrderID(orderID)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"status": "succeeded",
	}
	err = s.repo.UpdatePaymentDetails(payment.PaymentIntentID, updates)
	if err != nil {
		return err
	}

	// Update order status
	orderUpdates := map[string]interface{}{
		"payment_status": "succeeded",
		"status":         "processing",
	}
	return s.orderRepo.Update(orderID, orderUpdates)
}
