package models

// Request DTOs for Swagger documentation

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required" example:"1"`
	Quantity  int  `json:"quantity" binding:"required,min=1" example:"2"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1" example:"3"`
}

type CreateOrderRequest struct {
	Shipping float64 `json:"shipping" binding:"min=0" example:"10"`
}

type CheckoutRequest struct {
	Shipping float64 `json:"shipping" binding:"min=0" example:"10"`
	Currency string  `json:"currency" binding:"required" example:"usd"`
}

type AddToWishlistRequest struct {
	ProductID uint `json:"product_id" binding:"required" example:"1"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required" example:"shipped"`
}

type BulkCreateProductsRequest struct {
	Products []Product `json:"products" binding:"required,min=1,dive" example:"[{\"name\":\"Product 1\",\"description\":\"Description\",\"price\":99.99,\"category\":\"Electronics\",\"stock\":10}]"`
}
