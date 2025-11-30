package models

import (
	"errors"
	"time"
)

// Valid product categories
const (
	CategoryElectronics = "Electronics"
	CategoryAccessories = "Accessories"
	CategoryHome        = "Home"
	CategoryOffice      = "Office"
)

var ValidCategories = []string{
	CategoryElectronics,
	CategoryAccessories,
	CategoryHome,
	CategoryOffice,
}

// IsValidCategory checks if a category is valid
func IsValidCategory(category string) bool {
	for _, validCategory := range ValidCategories {
		if category == validCategory {
			return true
		}
	}
	return false
}

type Product struct {
	ID            uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string     `json:"name" gorm:"not null"`
	Description   string     `json:"description"`
	Price         float64    `json:"price" gorm:"not null"`
	OriginalPrice *float64   `json:"original_price,omitempty"`
	Rating        float64    `json:"rating" gorm:"default:0"`
	Image         string     `json:"image"`
	Category      string     `json:"category" gorm:"not null"`
	Badge         *string    `json:"badge,omitempty"`
	BadgeColor    *string    `json:"badge_color,omitempty"`
	Featured      bool       `json:"featured" gorm:"default:false"`
	Stock         int        `json:"stock" gorm:"default:0"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// Validate validates the product
func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if p.Category == "" {
		return errors.New("product category is required")
	}
	if !IsValidCategory(p.Category) {
		return errors.New("invalid category. Must be one of: Electronics, Accessories, Home, Office")
	}
	return nil
}
