package handlers

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/models"
	"go-ecommerce-api/services"
	"go-ecommerce-api/utils"
	"net/http"
	"strings"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email,min=3" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password"`
}

type RegisterPayload struct {
	Name     string `json:"name" binding:"required,min=2" example:"John Doe"`
	Email    string `json:"email" binding:"required,email,min=3" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password"`
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and return access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body LoginPayload true "Login credentials"
// @Success      200  {object}  models.LoginResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var LoginPayload LoginPayload

	if err := c.ShouldBindBodyWithJSON(&LoginPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	// Get user service from gin context (injected via middleware)
	userService := c.MustGet("userService").(services.UserServices)

	user, err := userService.GetUserByEmail(LoginPayload.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exista",
		})
		return
	}

	if err := utils.CheckPassword(LoginPayload.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateTokenPair(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to generate token",
		})
		return
	}

	// Set access token cookie
	c.SetCookie("acc_token", token.AccessToken, 3600, "/", "localhost", false, true)

	// Set refresh token cookie
	c.SetCookie("ref_token", token.RefreshToken, 86400, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"tokens": token,
	})
}

// Register godoc
// @Summary      Register new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register body RegisterPayload true "User registration data"
// @Success      201  {object}  models.RegisterResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /auth/register [post]
func Register(c *gin.Context) {
	var RegisterPayload RegisterPayload

	if err := c.ShouldBindBodyWithJSON(&RegisterPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong while binding json",
		})
		return
	}

	hashedPass, err := utils.HashPassword(RegisterPayload.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to hashed password",
		})
		return
	}

	// Create user object
	user := &models.User{
		Name:     RegisterPayload.Name,
		Email:    RegisterPayload.Email,
		Password: hashedPass,
		Role:     "user", // default role
	}

	// Get user service from gin context (injected via middleware)
	userService := c.MustGet("userService").(services.UserServices)

	if err := userService.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	// Generate tokens for the new user
	token, err := utils.GenerateTokenPair(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate tokens",
		})
		return
	}

	// Set access token cookie
	c.SetCookie("acc_token", token.AccessToken, 3600, "/", "localhost", false, true)

	// Set refresh token cookie
	c.SetCookie("ref_token", token.RefreshToken, 86400, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"tokens": token,
	})
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Generate new access and refresh tokens using refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.RefreshTokenResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /auth/refresh [get]
func RefreshToken(c *gin.Context) {
	// Get refresh token from cookie or header
	// refreshToken, err := c.Cookie("ref_token")
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token not found"})
	// 	return
	// }

	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token not found"})
		return
	}
	token := strings.Split(refreshToken, " ")[1]

	// Validate refresh token
	claims, err := utils.ValidateRefreshToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	// New TokenPair
	user := &models.User{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}

	newTokenPair, err := utils.GenerateTokenPair(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate tokens"})
		return
	}

	// Set NEW cookies
	c.SetCookie("acc_token", newTokenPair.AccessToken, 3600, "/", "localhost", false, true)
	c.SetCookie("ref_token", newTokenPair.RefreshToken, 86400, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tokens refreshed successfully",
		"tokens":  newTokenPair,
	})
}
