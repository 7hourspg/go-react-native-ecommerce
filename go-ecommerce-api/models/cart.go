package models

import "time"

type Cart struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint       `json:"user_id" gorm:"not null"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CartID    uint      `json:"cart_id" gorm:"not null"`
	Cart      *Cart     `json:"cart,omitempty" gorm:"foreignKey:CartID"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" gorm:"default:1;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
