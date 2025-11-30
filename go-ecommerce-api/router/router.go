package router

import (
	"go-ecommerce-api/database"
	"go-ecommerce-api/handlers"
	"go-ecommerce-api/middleware"
	"go-ecommerce-api/repositories"
	"go-ecommerce-api/services"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter(db database.Database) *gin.Engine {

	redis := database.NewRedisClient()

	// Get Stripe keys from environment (TODO: Move to config)
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	// User & Auth
	userRepo := repositories.NewUserRepository(db.GetDB(), redis)
	userServ := services.NewUserServices(userRepo)
	userHandle := handlers.NewUserHandlers(userServ)

	// Product
	productRepo := repositories.NewProductRepository(db.GetDB(), redis)
	productServ := services.NewProductServices(productRepo)
	productHandle := handlers.NewProductHandler(productServ)

	// Cart
	cartRepo := repositories.NewCartRepository(db.GetDB(), redis)
	cartServ := services.NewCartServices(cartRepo)
	cartHandle := handlers.NewCartHandler(cartServ)

	// Order
	orderRepo := repositories.NewOrderRepository(db.GetDB(), redis)
	orderServ := services.NewOrderServices(orderRepo)

	// Payment
	paymentRepo := repositories.NewPaymentRepository(db.GetDB())
	paymentServ := services.NewPaymentService(paymentRepo, orderRepo, stripeKey, webhookSecret)
	paymentHandle := handlers.NewPaymentHandler(paymentServ)

	// Order handler needs payment service for checkout
	orderHandle := handlers.NewOrderHandler(orderServ, cartServ, paymentServ)

	// Wishlist
	wishlistRepo := repositories.NewWishlistRepository(db.GetDB(), redis)
	wishlistServ := services.NewWishlistServices(wishlistRepo)
	wishlistHandle := handlers.NewWishlistHandler(wishlistServ)

	// ROUTER
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://127.0.0.1:3000", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(middleware.Logger())

	// SWAGGER DOCUMENTATION
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//PING
	router.GET("/ping", Health)

	// AUTH ROUTES
	authRoute := router.Group("/auth")
	authRoute.Use(middleware.UserServContext(userServ))
	authRoute.POST("/login", handlers.Login)
	authRoute.POST("/register", handlers.Register)
	authRoute.GET("/refresh", handlers.RefreshToken)

	// PUBLIC PRODUCT ROUTES
	productRoute := router.Group("/products")
	productRoute.GET("", productHandle.GetAllProducts)
	productRoute.GET("/featured", productHandle.GetFeaturedProducts)
	productRoute.GET("/category/:category", productHandle.GetProductsByCategory)
	productRoute.GET("/search", productHandle.SearchProducts)
	productRoute.GET("/:id", productHandle.GetProductByID)

	// PROTECTED ROUTES (require authentication)
	base := router.Group("/")
	base.Use(middleware.RequireAuth())

	// USER ROUTE
	userRoute := base.Group("user")
	userRoute.GET("", userHandle.GetByID)
	userRoute.DELETE("", userHandle.Delete)
	userRoute.PUT("", userHandle.Update)

	// CART ROUTES
	cartRoute := base.Group("cart")
	cartRoute.GET("", cartHandle.GetCart)
	cartRoute.DELETE("", cartHandle.ClearCart)
	cartRoute.POST("/items", cartHandle.AddToCart)
	cartRoute.PUT("/items/:id", cartHandle.UpdateCartItem)
	cartRoute.DELETE("/items/:id", cartHandle.RemoveFromCart)

	// ORDER ROUTES
	orderRoute := base.Group("orders")
	orderRoute.GET("", orderHandle.GetUserOrders)
	orderRoute.POST("", orderHandle.CreateOrder)
	orderRoute.POST("/checkout", orderHandle.Checkout) // Integrated checkout
	orderRoute.GET("/:id", orderHandle.GetOrderByID)

	// WISHLIST ROUTES
	wishlistRoute := base.Group("wishlist")
	wishlistRoute.GET("", wishlistHandle.GetWishlist)
	wishlistRoute.POST("", wishlistHandle.AddToWishlist)
	wishlistRoute.GET("/:product_id", wishlistHandle.CheckWishlist)
	wishlistRoute.DELETE("/:product_id", wishlistHandle.RemoveFromWishlist)

	// PAYMENT ROUTES
	paymentRoute := base.Group("payments")
	paymentRoute.POST("/create-intent", paymentHandle.CreatePaymentIntent)
	paymentRoute.POST("/confirm/:id", paymentHandle.ConfirmPayment)
	paymentRoute.POST("/cancel/:id", paymentHandle.CancelPayment)
	paymentRoute.POST("/confirm-success/:orderId", paymentHandle.ConfirmPaymentSuccess)
	paymentRoute.GET("/status/:id", paymentHandle.GetPaymentStatus)
	paymentRoute.GET("/order/:orderId", paymentHandle.GetPaymentByOrderID)
	paymentRoute.GET("/history", paymentHandle.GetUserPayments)

	// Webhook endpoint (outside auth middleware)
	router.POST("/webhooks/stripe", paymentHandle.HandleWebhook)

	// ADMIN ROUTES
	adminRoute := router.Group("/admin")
	adminRoute.Use(middleware.RequireAuth(), middleware.RequireAdmin())

	adminRoute.GET("/users", userHandle.GetAllAdmin)

	// Admin Product Routes
	adminProductRoute := adminRoute.Group("/products")
	adminProductRoute.POST("", productHandle.CreateProduct)
	adminProductRoute.POST("/bulk", productHandle.BulkCreateProducts)
	adminProductRoute.PUT("/:id", productHandle.UpdateProduct)
	adminProductRoute.DELETE("/:id", productHandle.DeleteProduct)

	// Admin Order Routes
	adminOrderRoute := adminRoute.Group("/orders")
	adminOrderRoute.GET("", orderHandle.GetAllOrders)
	adminOrderRoute.PUT("/:id/status", orderHandle.UpdateOrderStatus)

	return router

}

// Health godoc
// @Summary      Health check
// @Description  Check if the API is running
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.HealthResponse
// @Router       /ping [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Pong",
	})
}
