package handlers

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/services"
	"go-ecommerce-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderServices   *services.OrderServices
	CartServices    *services.CartServices
	PaymentServices *services.PaymentService
}

func NewOrderHandler(orderServ *services.OrderServices, cartServ *services.CartServices, paymentServ *services.PaymentService) *OrderHandler {
	return &OrderHandler{
		OrderServices:   orderServ,
		CartServices:    cartServ,
		PaymentServices: paymentServ,
	}
}

// GetOrderByID godoc
// @Summary      Get order by ID
// @Description  Retrieve a specific order by its ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  models.OrderResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order ID",
		})
		return
	}

	order, err := h.OrderServices.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// GetUserOrders godoc
// @Summary      Get user's orders
// @Description  Retrieve all orders for the authenticated user
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.OrdersResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /orders [get]
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	orders, err := h.OrderServices.GetOrdersByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GetAllOrders godoc
// @Summary      Get all orders (Admin)
// @Description  Retrieve all orders from all users (Admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.OrdersResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/orders [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.OrderServices.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CreateOrder godoc
// @Summary      Create order from cart
// @Description  Create a new order from the user's cart
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        request  body      models.CreateOrderRequest  true  "Order creation request"
// @Success      201  {object}  models.OrderResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	var req models.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		req.Shipping = 0 // Default shipping
	}

	// Get user's cart
	cart, err := h.CartServices.GetCartByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(cart.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cart is empty",
		})
		return
	}

	// Calculate total
	var total float64
	var orderItems []models.OrderItem

	for _, item := range cart.Items {
		itemTotal := item.Product.Price * float64(item.Quantity)
		total += itemTotal
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	total += req.Shipping

	// Create order
	order := models.Order{
		UserID:   userID.(uint),
		Status:   "pending",
		Total:    total,
		Shipping: req.Shipping,
		Items:    orderItems,
	}

	if err := h.OrderServices.Create(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	// Clear cart after order creation
	h.CartServices.ClearCart(cart.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"data":    order,
	})
}

// UpdateOrderStatus godoc
// @Summary      Update order status (Admin)
// @Description  Update the status of an order (Admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Param        request  body      models.UpdateOrderStatusRequest  true  "Status update"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order ID",
		})
		return
	}

	var req models.UpdateOrderStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	if err := h.OrderServices.UpdateStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update order status",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order status updated successfully",
	})
}

// Checkout godoc
// @Summary      Checkout and create payment
// @Description  Create order and payment intent in one step. Cart is automatically cleared after successful checkout.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.CheckoutResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /orders/checkout [post]
func (h *OrderHandler) Checkout(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	// Get user's cart
	cart, err := h.CartServices.GetCartByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve cart",
		})
		return
	}

	if len(cart.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cart is empty",
		})
		return
	}

	// Calculate total using cart utility
	cartSummary := utils.CalculateCartTotals(cart)

	// Create order items
	var orderItems []models.OrderItem
	for _, item := range cart.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	// Create order
	order := models.Order{
		UserID:        userID.(uint),
		Status:        "pending",
		Total:         cartSummary.Total,
		Shipping:      cartSummary.Shipping,
		Items:         orderItems,
		PaymentStatus: "pending",
	}

	createdOrder, err := h.OrderServices.CreateOrderWithPayment(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	// Create payment intent
	pi, payment, err := h.PaymentServices.CreatePaymentIntent(createdOrder.ID, cartSummary.Total, "usd")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Order created but payment initialization failed",
			"error":   err.Error(),
			"order":   createdOrder,
		})
		return
	}

	// Clear cart after successful order and payment intent creation
	h.CartServices.ClearCart(cart.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Checkout successful",
		"client_secret": pi.ClientSecret,
		"order": gin.H{
			"id":             createdOrder.ID,
			"status":         createdOrder.Status,
			"total":          createdOrder.Total,
			"shipping":       createdOrder.Shipping,
			"payment_status": createdOrder.PaymentStatus,
		},
		"payment": gin.H{
			"payment_intent_id": pi.ID,
			"amount":            payment.Amount,
			"currency":          payment.Currency,
			"status":            payment.Status,
		},
	})
}
