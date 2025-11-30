package main

import (
	"fmt"
	"go-ecommerce-api/database"
	"go-ecommerce-api/router"

	_ "go-ecommerce-api/docs" // This is important for swagger to work
)

// @title           E-commerce API
// @version         1.0
// @description     A RESTful E-commerce API built with Go and Gin
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {

	// calling DB
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("DB is not working properly")
		panic(err)
	}
	defer db.Close()

	r := router.SetUpRouter(db)

	port := "8080"
	fmt.Printf("Server starting on port %s\n", port)
	r.Run(":" + port)
}
