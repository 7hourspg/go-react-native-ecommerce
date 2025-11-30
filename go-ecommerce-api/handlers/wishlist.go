package handlers

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	WishlistServices *services.WishlistServices
}

func NewWishlistHandler(s *services.WishlistServices) *WishlistHandler {
	return &WishlistHandler{
		WishlistServices: s,
	}
}

// GetWishlist godoc
// @Summary      Get user's wishlist
// @Description  Retrieve all items in the authenticated user's wishlist
// @Tags         wishlist
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.WishlistResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /wishlist [get]
func (h *WishlistHandler) GetWishlist(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	wishlist, err := h.WishlistServices.GetWishlistByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wishlist})
}

// AddToWishlist godoc
// @Summary      Add product to wishlist
// @Description  Add a product to the user's wishlist
// @Tags         wishlist
// @Accept       json
// @Produce      json
// @Param        request  body      models.AddToWishlistRequest  true  "Add to wishlist request"
// @Success      201  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /wishlist [post]
func (h *WishlistHandler) AddToWishlist(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	var req models.AddToWishlistRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	if err := h.WishlistServices.AddToWishlist(userID.(uint), req.ProductID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add to wishlist",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added to wishlist",
	})
}

// RemoveFromWishlist godoc
// @Summary      Remove product from wishlist
// @Description  Remove a product from the user's wishlist
// @Tags         wishlist
// @Accept       json
// @Produce      json
// @Param        product_id   path      int  true  "Product ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /wishlist/{product_id} [delete]
func (h *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	if err := h.WishlistServices.RemoveFromWishlist(userID.(uint), uint(productID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove from wishlist",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product removed from wishlist",
	})
}

// CheckWishlist godoc
// @Summary      Check if product is in wishlist
// @Description  Check if a product is in the user's wishlist
// @Tags         wishlist
// @Accept       json
// @Produce      json
// @Param        product_id   path      int  true  "Product ID"
// @Success      200  {object}  models.WishlistCheckResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /wishlist/{product_id} [get]
func (h *WishlistHandler) CheckWishlist(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	isInWishlist, err := h.WishlistServices.IsInWishlist(userID.(uint), uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"in_wishlist": isInWishlist,
	})
}

