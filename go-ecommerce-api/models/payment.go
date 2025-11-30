package models

import "time"

type Payment struct {
	ID              uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID         uint       `json:"order_id" gorm:"not null;index"`
	Order           Order      `json:"order" gorm:"foreignKey:OrderID"`
	PaymentIntentID string     `json:"payment_intent_id" gorm:"uniqueIndex"`
	Amount          float64    `json:"amount" gorm:"not null"`
	Currency        string     `json:"currency" gorm:"default:'usd'"`
	Status          string     `json:"status" gorm:"default:'pending'"` // pending, processing, succeeded, failed, cancelled
	PaymentMethod   string     `json:"payment_method,omitempty"`        // card, cash, wallet, etc.
	TransactionID   string     `json:"transaction_id,omitempty"`
	FailureReason   string     `json:"failure_reason,omitempty"`
	Metadata        *string    `json:"metadata,omitempty" gorm:"type:jsonb"` // Additional metadata as JSON
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
