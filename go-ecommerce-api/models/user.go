package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"uniqueIndex"`
	Password  string     `json:"-" gorm:"not null"`
	Role      string     `json:"role" gorm:"default:'user'"`
	Cart      *Cart      `json:"cart,omitempty" gorm:"foreignKey:UserID"`
	Orders    []Order    `json:"orders,omitempty" gorm:"foreignKey:UserID"`
	Wishlist  []Wishlist `json:"wishlist,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
