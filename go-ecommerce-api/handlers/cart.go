package handlers

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/services"
	"go-ecommerce-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartServices *services.CartServices
}

func NewCartHandler(s *services.CartServices) *CartHandler {
	return &CartHandler{
		CartServices: s,
	}
}

// GetCart godoc
// @Summary      Get user's cart
// @Description  Retrieve the authenticated user's shopping cart with calculated totals
// @Tags         cart
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.CartSummaryResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /cart [get]
func (h *CartHandler) GetCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	cart, err := h.CartServices.GetCartByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	summary := utils.CalculateCartTotals(cart)
	c.JSON(http.StatusOK, gin.H{"data": summary})
}

// AddToCart godoc
// @Summary      Add item to cart
// @Description  Add a product to the user's cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        request  body      models.AddToCartRequest  true  "Add to cart request"
// @Success      201  {object}  models.CartItemResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /cart/items [post]
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	var req models.AddToCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	cart, err := h.CartServices.GetCartByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, err := h.CartServices.AddItem(cart.ID, req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add item to cart",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item added to cart",
		"data":    item,
	})
}

// UpdateCartItem godoc
// @Summary      Update cart item quantity
// @Description  Update the quantity of an item in the cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Cart Item ID"
// @Param        request  body      models.UpdateCartItemRequest  true  "Update quantity"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /cart/items/{id} [put]
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid item ID",
		})
		return
	}

	var req models.UpdateCartItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	if err := h.CartServices.UpdateItemQuantity(uint(id), req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update cart item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cart item updated successfully",
	})
}

// RemoveFromCart godoc
// @Summary      Remove item from cart
// @Description  Remove an item from the user's cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Cart Item ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /cart/items/{id} [delete]
func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid item ID",
		})
		return
	}

	if err := h.CartServices.RemoveItem(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove item from cart",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item removed from cart",
	})
}

// ClearCart godoc
// @Summary      Clear cart
// @Description  Remove all items from the user's cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.SuccessResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /cart [delete]
func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	cart, err := h.CartServices.GetCartByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.CartServices.ClearCart(cart.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to clear cart",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cart cleared successfully",
	})
}
