package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderByID(id uint) (*models.Order, error)
	GetOrdersByUserID(userID uint) ([]models.Order, error)
	GetAllOrders() ([]models.Order, error)
	Create(order *models.Order) error
	UpdateStatus(id uint, status string) error
	Update(id uint, updates map[string]interface{}) error
}

type orderRepository struct {
	DB    *gorm.DB
	Redis database.RedisClient
}

func NewOrderRepository(db *gorm.DB, redis database.RedisClient) OrderRepository {
	return &orderRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *orderRepository) GetOrderByID(id uint) (*models.Order, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("order:%d", id)

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var order models.Order
		if err := json.Unmarshal([]byte(val), &order); err == nil {
			return &order, nil
		}
	}

	var order models.Order
	err = r.DB.Preload("Items.Product").Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	orderJSON, _ := json.Marshal(order)
	r.Redis.Set(ctx, redisKey, orderJSON)

	return &order, nil
}

func (r *orderRepository) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("orders:user:%d", userID)

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var orders []models.Order
		if err := json.Unmarshal([]byte(val), &orders); err == nil {
			return orders, nil
		}
	}

	var orders []models.Order
	err = r.DB.Preload("Items.Product").Where("user_id = ?", userID).Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	orderJSON, _ := json.Marshal(orders)
	r.Redis.Set(ctx, redisKey, orderJSON)

	return orders, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	ctx := context.Background()
	redisKey := "orders:all"

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var orders []models.Order
		if err := json.Unmarshal([]byte(val), &orders); err == nil {
			return orders, nil
		}
	}

	var orders []models.Order
	err = r.DB.Preload("Items.Product").Preload("User").Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	orderJSON, _ := json.Marshal(orders)
	r.Redis.Set(ctx, redisKey, orderJSON)

	return orders, nil
}

func (r *orderRepository) Create(order *models.Order) error {
	err := r.DB.Create(order).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	redisKey := fmt.Sprintf("order:%d", order.ID)
	orderJSON, _ := json.Marshal(order)
	r.Redis.Set(ctx, redisKey, orderJSON)

	r.Redis.Del(ctx, fmt.Sprintf("orders:user:%d", order.UserID))
	r.Redis.Del(ctx, "orders:all")

	return nil
}

func (r *orderRepository) UpdateStatus(id uint, status string) error {
	err := r.DB.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	var order models.Order
	r.DB.Where("id = ?", id).First(&order)

	r.Redis.Del(ctx, fmt.Sprintf("order:%d", id))
	r.Redis.Del(ctx, fmt.Sprintf("orders:user:%d", order.UserID))
	r.Redis.Del(ctx, "orders:all")

	return nil
}

func (r *orderRepository) Update(id uint, updates map[string]interface{}) error {
	err := r.DB.Model(&models.Order{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	var order models.Order
	r.DB.Where("id = ?", id).First(&order)

	r.Redis.Del(ctx, fmt.Sprintf("order:%d", id))
	r.Redis.Del(ctx, fmt.Sprintf("orders:user:%d", order.UserID))
	r.Redis.Del(ctx, "orders:all")

	return nil
}
