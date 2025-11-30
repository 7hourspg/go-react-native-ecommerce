package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartByUserID(userID uint) (*models.Cart, error)
	GetCartItemByID(id uint) (*models.CartItem, error)
	AddItem(cartID uint, productID uint, quantity int) (*models.CartItem, error)
	UpdateItemQuantity(id uint, quantity int) error
	RemoveItem(id uint) error
	ClearCart(cartID uint) error
	CreateCart(userID uint) (*models.Cart, error)
}

type cartRepository struct {
	DB    *gorm.DB
	Redis database.RedisClient
}

func NewCartRepository(db *gorm.DB, redis database.RedisClient) CartRepository {
	return &cartRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *cartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("cart:user:%d", userID)

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var cart models.Cart
		if err := json.Unmarshal([]byte(val), &cart); err == nil {
			return &cart, nil
		}
	}

	var cart models.Cart
	err = r.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at ASC")
	}).Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		// Create cart if it doesn't exist
		cart = models.Cart{UserID: userID}
		if err := r.DB.Create(&cart).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	cartJSON, _ := json.Marshal(cart)
	r.Redis.Set(ctx, redisKey, cartJSON)

	return &cart, nil
}

func (r *cartRepository) GetCartItemByID(id uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.Preload("Product").Where("id = ?", id).First(&item).Error
	return &item, err
}

func (r *cartRepository) AddItem(cartID uint, productID uint, quantity int) (*models.CartItem, error) {
	// Check if item already exists in cart
	var existingItem models.CartItem
	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&existingItem).Error

	if err == nil {
		// Update quantity
		existingItem.Quantity += quantity
		if err := r.DB.Save(&existingItem).Error; err != nil {
			return nil, err
		}
		r.invalidateCartCache(cartID)
		return &existingItem, nil
	}

	// Create new item
	item := models.CartItem{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}

	if err := r.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	r.DB.Preload("Product").First(&item, item.ID)
	r.invalidateCartCache(cartID)

	return &item, nil
}

func (r *cartRepository) UpdateItemQuantity(id uint, quantity int) error {
	var item models.CartItem
	if err := r.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return err
	}

	item.Quantity = quantity
	if err := r.DB.Save(&item).Error; err != nil {
		return err
	}

	r.invalidateCartCache(item.CartID)
	return nil
}

func (r *cartRepository) RemoveItem(id uint) error {
	var item models.CartItem
	if err := r.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return err
	}

	cartID := item.CartID
	if err := r.DB.Delete(&item).Error; err != nil {
		return err
	}

	r.invalidateCartCache(cartID)
	return nil
}

func (r *cartRepository) ClearCart(cartID uint) error {
	if err := r.DB.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error; err != nil {
		return err
	}

	r.invalidateCartCache(cartID)
	return nil
}

func (r *cartRepository) CreateCart(userID uint) (*models.Cart, error) {
	cart := models.Cart{UserID: userID}
	if err := r.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	r.invalidateCartCache(cart.ID)
	return &cart, nil
}

func (r *cartRepository) invalidateCartCache(cartID uint) {
	ctx := context.Background()
	var cart models.Cart
	r.DB.Where("id = ?", cartID).First(&cart)
	if cart.UserID != 0 {
		r.Redis.Del(ctx, fmt.Sprintf("cart:user:%d", cart.UserID))
	}
}
