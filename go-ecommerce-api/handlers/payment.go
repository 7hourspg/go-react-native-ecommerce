package handlers

import (
	"go-ecommerce-api/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *services.PaymentService
}

func NewPaymentHandler(service *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

type CreatePaymentIntentRequest struct {
	OrderID  uint    `json:"order_id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

// Create payment intent
func (h *PaymentHandler) CreatePaymentIntent(c *gin.Context) {
	var req CreatePaymentIntentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pi, payment, err := h.service.CreatePaymentIntent(req.OrderID, req.Amount, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_secret":     pi.ClientSecret,
		"payment_intent_id": pi.ID,
		"payment":           payment,
	})
}

// Confirm payment
func (h *PaymentHandler) ConfirmPayment(c *gin.Context) {
	paymentIntentID := c.Param("id")

	pi, err := h.service.ConfirmPaymentIntent(paymentIntentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         pi.Status,
		"payment_intent": pi,
	})
}

// Cancel payment
func (h *PaymentHandler) CancelPayment(c *gin.Context) {
	paymentIntentID := c.Param("id")

	pi, err := h.service.CancelPaymentIntent(paymentIntentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  pi.Status,
		"message": "Payment cancelled",
	})
}

// Webhook handler
func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	signature := c.GetHeader("Stripe-Signature")

	err = h.service.HandleWebhook(payload, signature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Get payment status
func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	paymentIntentID := c.Param("id")

	pi, err := h.service.GetPaymentIntent(paymentIntentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   pi.Status,
		"amount":   pi.Amount,
		"currency": pi.Currency,
	})
}

// Get payment by order ID
func (h *PaymentHandler) GetPaymentByOrderID(c *gin.Context) {
	orderID := c.Param("orderId")

	payment, err := h.service.GetPaymentByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// Get payment history for a user
func (h *PaymentHandler) GetUserPayments(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	payments, err := h.service.GetUserPayments(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// Confirm payment success godoc
// @Summary      Confirm payment success
// @Description  Manually confirm payment success and update order status (for testing without webhooks)
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        orderId  path  string  true  "Order ID"
// @Success      200  {object}  models.PaymentSuccessResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /payments/confirm-success/{orderId} [post]
func (h *PaymentHandler) ConfirmPaymentSuccess(c *gin.Context) {
	orderID := c.Param("orderId")

	payment, err := h.service.GetPaymentByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Update payment and order status
	err = h.service.ConfirmPaymentSuccess(payment.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment confirmed"})
}
