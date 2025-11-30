package handlers

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	services services.UserServices
}

func NewUserHandlers(h services.UserServices) *userHandlers {
	return &userHandlers{
		services: h,
	}
}

// GetAllAdmin godoc
// @Summary      Get all users (Admin)
// @Description  Retrieve all users (Admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UsersResponse
// @Failure      400  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /admin/users [get]
func (u *userHandlers) GetAllAdmin(c *gin.Context) {
	users, err := u.services.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetByID godoc
// @Summary      Get current user
// @Description  Retrieve the authenticated user's information
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /user [get]
func (u *userHandlers) GetByID(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		c.Abort()
		return
	}

	id := userID.(uint)

	user, err := u.services.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *userHandlers) Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	if err := u.services.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Woo hoo! user has been created",
	})

}
// Update godoc
// @Summary      Update current user
// @Description  Update the authenticated user's information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Updated user data"
// @Success      202  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /user [put]
func (u *userHandlers) Update(c *gin.Context) {

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		c.Abort()
		return
	}

	id := userID.(uint)

	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	if err := u.services.Update(&user, uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User Updated successfully",
	})

}
// Delete godoc
// @Summary      Delete current user
// @Description  Delete the authenticated user's account
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      202  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /user [delete]
func (u *userHandlers) Delete(c *gin.Context) {

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		c.Abort()
		return
	}

	id := userID.(uint)

	if err := u.services.Delete(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User Deleted successfully",
	})
}
