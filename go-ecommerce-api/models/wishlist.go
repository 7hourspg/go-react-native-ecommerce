package models

import "time"

type Wishlist struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null;uniqueIndex:idx_user_product"`
	User      *User     `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint      `json:"product_id" gorm:"not null;uniqueIndex:idx_user_product"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
