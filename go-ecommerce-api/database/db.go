package database

import (
	"go-ecommerce-api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	GetDB() *gorm.DB
	Migrate() error
	Close() error
}

type database struct {
	Db *gorm.DB
}

func NewDatabase() (Database, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://7hours:@localhost:5432/ecommerce_app?sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Printf("[error] failed to initialize database, got error %v", err)
		return nil, err
	}

	database := &database{Db: db}

	// Auto migrate tables
	if err := database.Migrate(); err != nil {
		log.Printf("[error] failed to migrate database, got error %v", err)
		return nil, err
	}

	return database, nil
}

// METHODS

func (d *database) GetDB() *gorm.DB {
	return d.Db
}

func (d *database) Migrate() error {
	err := d.Db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Wishlist{},
		&models.Payment{},
	)
	return err
}

func (d *database) Close() error {
	sqlDb, err := d.Db.DB()

	if err != nil {
		return err
	}

	return sqlDb.Close()
}
