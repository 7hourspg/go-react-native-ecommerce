package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"

	"gorm.io/gorm"
)

type WishlistRepository interface {
	GetWishlistByUserID(userID uint) ([]models.Wishlist, error)
	AddToWishlist(userID uint, productID uint) error
	RemoveFromWishlist(userID uint, productID uint) error
	IsInWishlist(userID uint, productID uint) (bool, error)
}

type wishlistRepository struct {
	DB    *gorm.DB
	Redis database.RedisClient
}

func NewWishlistRepository(db *gorm.DB, redis database.RedisClient) WishlistRepository {
	return &wishlistRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *wishlistRepository) GetWishlistByUserID(userID uint) ([]models.Wishlist, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("wishlist:user:%d", userID)

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var wishlist []models.Wishlist
		if err := json.Unmarshal([]byte(val), &wishlist); err == nil {
			return wishlist, nil
		}
	}

	var wishlist []models.Wishlist
	err = r.DB.Preload("Product").Where("user_id = ?", userID).Find(&wishlist).Error
	if err != nil {
		return nil, err
	}

	wishlistJSON, _ := json.Marshal(wishlist)
	r.Redis.Set(ctx, redisKey, wishlistJSON)

	return wishlist, nil
}

func (r *wishlistRepository) AddToWishlist(userID uint, productID uint) error {
	// Check if already exists
	var existing models.Wishlist
	err := r.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&existing).Error
	if err == nil {
		return nil // Already in wishlist
	}

	wishlist := models.Wishlist{
		UserID:    userID,
		ProductID: productID,
	}

	if err := r.DB.Create(&wishlist).Error; err != nil {
		return err
	}

	ctx := context.Background()
	r.Redis.Del(ctx, fmt.Sprintf("wishlist:user:%d", userID))

	return nil
}

func (r *wishlistRepository) RemoveFromWishlist(userID uint, productID uint) error {
	err := r.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Wishlist{}).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	r.Redis.Del(ctx, fmt.Sprintf("wishlist:user:%d", userID))

	return nil
}

func (r *wishlistRepository) IsInWishlist(userID uint, productID uint) (bool, error) {
	var wishlist models.Wishlist
	err := r.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&wishlist).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
