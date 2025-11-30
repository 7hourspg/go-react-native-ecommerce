package repositories

import (
	"go-ecommerce-api/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// Create a new payment
func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

// Save payment intent details to database
func (r *PaymentRepository) SavePaymentIntent(orderID uint, paymentIntentID string, amount float64, currency string, status string) (*models.Payment, error) {
	payment := &models.Payment{
		OrderID:         orderID,
		PaymentIntentID: paymentIntentID,
		Amount:          amount,
		Currency:        currency,
		Status:          status,
	}
	err := r.db.Create(payment).Error
	return payment, err
}

// Update payment status
func (r *PaymentRepository) UpdatePaymentStatus(paymentIntentID string, status string) error {
	return r.db.Model(&models.Payment{}).
		Where("payment_intent_id = ?", paymentIntentID).
		Update("status", status).Error
}

// Update payment with transaction details
func (r *PaymentRepository) UpdatePaymentDetails(paymentIntentID string, updates map[string]interface{}) error {
	return r.db.Model(&models.Payment{}).
		Where("payment_intent_id = ?", paymentIntentID).
		Updates(updates).Error
}

// Get payment by payment intent ID
func (r *PaymentRepository) GetPaymentByIntentID(paymentIntentID string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("payment_intent_id = ?", paymentIntentID).First(&payment).Error
	return &payment, err
}

// Get payment by order ID
func (r *PaymentRepository) GetPaymentByOrderID(orderID uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

// Get all payments for a user
func (r *PaymentRepository) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Joins("JOIN orders ON orders.id = payments.order_id").
		Where("orders.user_id = ?", userID).
		Find(&payments).Error
	return payments, err
}

// Get payment by ID
func (r *PaymentRepository) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, id).Error
	return &payment, err
}
