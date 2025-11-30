package models

// Response models for Swagger documentation

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type ProductResponse struct {
	Data Product `json:"data"`
}

type ProductsResponse struct {
	Data []Product `json:"data"`
}

type CartResponse struct {
	Data Cart `json:"data"`
}

type CartSummary struct {
	Cart
	Subtotal float64 `json:"subtotal"`
	Shipping float64 `json:"shipping"`
	Tax      float64 `json:"tax"`
	TaxRate  float64 `json:"tax_rate"`
	Total    float64 `json:"total"`
}

type CartSummaryResponse struct {
	Data CartSummary `json:"data"`
}

type CartItemResponse struct {
	Message string   `json:"message"`
	Data    CartItem `json:"data"`
}

type OrderResponse struct {
	Message string `json:"message"`
	Data    Order  `json:"data"`
}

type OrdersResponse struct {
	Data []Order `json:"data"`
}

type WishlistResponse struct {
	Data []Wishlist `json:"data"`
}

type WishlistCheckResponse struct {
	InWishlist bool `json:"in_wishlist"`
}

type CheckoutOrderResponse struct {
	ID            uint    `json:"id" example:"1"`
	Status        string  `json:"status" example:"pending"`
	Total         float64 `json:"total" example:"292.64"`
	Shipping      float64 `json:"shipping" example:"5.99"`
	PaymentStatus string  `json:"payment_status" example:"pending"`
}

type CheckoutPaymentResponse struct {
	PaymentIntentID string  `json:"payment_intent_id" example:"pi_3SZA4XIH7jU9U4pR0VhsSo69"`
	Amount          float64 `json:"amount" example:"292.64"`
	Currency        string  `json:"currency" example:"usd"`
	Status          string  `json:"status" example:"requires_payment_method"`
}

type CheckoutResponse struct {
	Message      string                  `json:"message" example:"Checkout successful"`
	ClientSecret string                  `json:"client_secret" example:"pi_xxx_secret_yyy"`
	Order        CheckoutOrderResponse   `json:"order"`
	Payment      CheckoutPaymentResponse `json:"payment"`
}

type PaymentSuccessResponse struct {
	Message string `json:"message" example:"Payment confirmed"`
}

type UserResponse struct {
	Data User `json:"data"`
}

type UsersResponse struct {
	Data []User `json:"data"`
}

type LoginResponse struct {
	Message string                 `json:"message"`
	User    map[string]interface{} `json:"user"`
	Tokens  map[string]interface{} `json:"tokens"`
}

type RegisterResponse struct {
	Message string                 `json:"message"`
	User    map[string]interface{} `json:"user"`
	Tokens  map[string]interface{} `json:"tokens"`
}

type RefreshTokenResponse struct {
	Message string                 `json:"message"`
	Tokens  map[string]interface{} `json:"tokens"`
}

type HealthResponse struct {
	Message string `json:"Message"`
}
