package utils

import "go-ecommerce-api/models"

const (
	TaxRate               = 0.18  // 18% tax rate
	ShippingCost          = 5.99  // Shipping cost
	FreeShippingThreshold = 100.0 // Free shipping for orders over $100
)

// CalculateCartTotals calculates the subtotal, shipping, tax, and total for a cart
func CalculateCartTotals(cart *models.Cart) *models.CartSummary {
	var subtotal float64

	// Calculate subtotal from cart items
	for _, item := range cart.Items {
		if item.Product.ID != 0 { // Ensure product is loaded
			subtotal += item.Product.Price * float64(item.Quantity)
		}
	}

	// Calculate shipping
	var shipping float64
	if subtotal > 0 && subtotal <= FreeShippingThreshold {
		shipping = ShippingCost
	}

	// Calculate tax
	tax := subtotal * TaxRate

	// Calculate total
	total := subtotal + shipping + tax

	return &models.CartSummary{
		Cart:     *cart,
		Subtotal: subtotal,
		Shipping: shipping,
		Tax:      tax,
		TaxRate:  TaxRate,
		Total:    total,
	}
}
