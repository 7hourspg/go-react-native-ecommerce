package models

import "time"

type Order struct {
	ID              uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          uint        `json:"user_id" gorm:"not null"`
	User            User        `json:"user" gorm:"foreignKey:UserID"`
	Status          string      `json:"status" gorm:"default:'pending'"` // pending, processing, shipped, delivered, cancelled
	Total           float64     `json:"total" gorm:"not null"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Shipping        float64     `json:"shipping" gorm:"default:0"`
	PaymentIntentID string      `json:"payment_intent_id,omitempty"`
	PaymentStatus   string      `json:"payment_status" gorm:"default:'pending'"` // pending, processing, succeeded, failed, cancelled
	PaymentMethod   string      `json:"payment_method,omitempty"`                // card, cash, etc.
	TransactionID   string      `json:"transaction_id,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	DeletedAt       *time.Time  `json:"deleted_at,omitempty" gorm:"index"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   uint      `json:"order_id" gorm:"not null"`
	Order     Order     `json:"order" gorm:"foreignKey:OrderID"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"` // Price at time of order
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
