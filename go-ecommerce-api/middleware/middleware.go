package middleware

import (
	"go-ecommerce-api/repositories"
	"go-ecommerce-api/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("latency: %d ms", latency.Milliseconds())

		// access the status we are sending
		// status := c.Writer.Status()
		// log.Println(status)
	}
}

// ErrorHandler captures errors and returns a consistent JSON error response
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// First check if user is authenticated
		role, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
			c.Abort()
			return
		}

		// is Admin
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header not found"})
			c.Abort()
			return
		}
		token := strings.Split(authorization, " ")[1]

		// fmt.Println(token)
		// token, err := c.Cookie("acc_token")
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Access token not found"})
		// 	c.Abort()
		// 	return
		// }

		// validate
		claims, err := utils.ValidateAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid access token"})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}

func UserServContext(s repositories.UserRepositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userService", s)
		c.Next()
	}
}
