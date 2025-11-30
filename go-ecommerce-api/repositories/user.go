package repositories

import (
	"encoding/json"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"

	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositories interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User, id uint) error
	Delete(id uint) error
	GetUserByEmail(e string) (*models.User, error)
}

type userRepositories struct {
	DB    *gorm.DB
	Redis database.RedisClient
}

func NewUserRepository(db *gorm.DB, redis database.RedisClient) UserRepositories {
	return &userRepositories{
		DB:    db,
		Redis: redis,
	}
}

func (r *userRepositories) GetAll() ([]models.User, error) {

	ctx := context.Background()
	redisKey := "user:all"

	val, err := r.Redis.Get(ctx, redisKey)

	// returing cached data
	if err == nil && val != "" {
		var users []models.User
		if err := json.Unmarshal([]byte(val), &users); err == nil {
			return users, nil
		}
	}

	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	// Cache the result
	userJSON, _ := json.Marshal(users)
	r.Redis.Set(ctx, redisKey, userJSON)

	return users, nil
}

func (r *userRepositories) GetByID(id uint) (*models.User, error) {
	// Redis
	ctx := context.Background()
	redisKey := fmt.Sprintf("user:%d", id)

	val, err := r.Redis.Get(ctx, redisKey)

	if err == nil && val != "" {
		var user models.User
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}

	// if data is not cached then fetching db
	var user models.User
	err = r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Cache the result
	userJSON, _ := json.Marshal(user)
	r.Redis.Set(ctx, redisKey, userJSON)

	return &user, nil
}

func (r *userRepositories) GetUserByEmail(e string) (*models.User, error) {
	// Redis
	ctx := context.Background()
	redisKey := fmt.Sprintf("user:%s", e)

	val, err := r.Redis.Get(ctx, redisKey)

	if err == nil && val != "" {
		var user models.User
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}

	// if data is not cached then fetching db
	var user models.User
	err = r.DB.Where("email = ?", e).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Cache the result
	userJSON, _ := json.Marshal(user)
	r.Redis.Set(ctx, redisKey, userJSON)

	return &user, nil
}

func (r *userRepositories) Create(user *models.User) error {
	err := r.DB.Create(user).Error

	// Redis
	ctx := context.Background()
	redisKey := fmt.Sprintf("user:%d", user.ID)

	// Cache the result
	userJSON, _ := json.Marshal(user)
	r.Redis.Set(ctx, redisKey, userJSON)

	// invalidate all user
	r.Redis.Del(ctx, "user:all")
	return err
}

func (r *userRepositories) Update(user *models.User, id uint) error {
	err := r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error

	// Invalidate cache
	ctx := context.Background()
	r.Redis.Del(ctx, fmt.Sprintf("user:%d", id))
	r.Redis.Del(ctx, "user:all")
	return err

}

func (r *userRepositories) Delete(id uint) error {
	err := r.DB.Where("id = ?", id).Delete(&models.User{}).Error

	// Invalidate cache
	ctx := context.Background()
	r.Redis.Del(ctx, fmt.Sprintf("user:%d", id))
	r.Redis.Del(ctx, "user:all")
	return err
}
