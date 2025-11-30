package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(id uint) (*models.Product, error)
	GetAll() ([]models.Product, error)
	GetFeatured() ([]models.Product, error)
	GetByCategory(category string) ([]models.Product, error)
	Search(query string) ([]models.Product, error)
	Create(product *models.Product) error
	BulkCreate(products []models.Product) ([]models.Product, error)
	Update(id uint, product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	DB    *gorm.DB
	Redis database.RedisClient
}

func NewProductRepository(db *gorm.DB, redis database.RedisClient) ProductRepository {
	return &productRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("product:%d", id)

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var product models.Product
		if err := json.Unmarshal([]byte(val), &product); err == nil {
			return &product, nil
		}
	}

	var product models.Product
	err = r.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	productJSON, _ := json.Marshal(product)
	r.Redis.Set(ctx, redisKey, productJSON)

	return &product, nil
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	ctx := context.Background()
	redisKey := "products:all"

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var products []models.Product
		if err := json.Unmarshal([]byte(val), &products); err == nil {
			return products, nil
		}
	}

	var products []models.Product
	err = r.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	productJSON, _ := json.Marshal(products)
	r.Redis.Set(ctx, redisKey, productJSON)

	return products, nil
}

func (r *productRepository) GetFeatured() ([]models.Product, error) {
	ctx := context.Background()
	redisKey := "products:featured"

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var products []models.Product
		if err := json.Unmarshal([]byte(val), &products); err == nil {
			return products, nil
		}
	}

	var products []models.Product
	err = r.DB.Where("featured = ?", true).Find(&products).Error
	if err != nil {
		return nil, err
	}

	productJSON, _ := json.Marshal(products)
	r.Redis.Set(ctx, redisKey, productJSON)

	return products, nil
}

func (r *productRepository) GetByCategory(category string) ([]models.Product, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("products:category:%s", category)

	val, err := r.Redis.Get(ctx, redisKey)
	if err == nil && val != "" {
		var products []models.Product
		if err := json.Unmarshal([]byte(val), &products); err == nil {
			return products, nil
		}
	}

	var products []models.Product
	err = r.DB.Where("category = ?", category).Find(&products).Error
	if err != nil {
		return nil, err
	}

	productJSON, _ := json.Marshal(products)
	r.Redis.Set(ctx, redisKey, productJSON)

	return products, nil
}

func (r *productRepository) Search(query string) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	return products, err
}

func (r *productRepository) Create(product *models.Product) error {
	err := r.DB.Create(product).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	redisKey := fmt.Sprintf("product:%d", product.ID)
	productJSON, _ := json.Marshal(product)
	r.Redis.Set(ctx, redisKey, productJSON)

	r.Redis.Del(ctx, "products:all")
	r.Redis.Del(ctx, "products:featured")
	if product.Category != "" {
		r.Redis.Del(ctx, fmt.Sprintf("products:category:%s", product.Category))
	}

	return nil
}

func (r *productRepository) BulkCreate(products []models.Product) ([]models.Product, error) {
	// Validate all products before creating
	for i := range products {
		if err := products[i].Validate(); err != nil {
			return nil, fmt.Errorf("product at index %d is invalid: %w", i, err)
		}
	}

	err := r.DB.Create(&products).Error
	if err != nil {
		return nil, err
	}

	// Invalidate all relevant Redis caches
	ctx := context.Background()
	categories := make(map[string]bool)
	for _, product := range products {
		redisKey := fmt.Sprintf("product:%d", product.ID)
		productJSON, _ := json.Marshal(product)
		r.Redis.Set(ctx, redisKey, productJSON)

		if product.Category != "" {
			categories[product.Category] = true
		}
	}

	// Clear cache for all affected categories and main lists
	r.Redis.Del(ctx, "products:all")
	r.Redis.Del(ctx, "products:featured")
	for category := range categories {
		r.Redis.Del(ctx, fmt.Sprintf("products:category:%s", category))
	}

	return products, nil
}

func (r *productRepository) Update(id uint, product *models.Product) error {
	err := r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	r.Redis.Del(ctx, fmt.Sprintf("product:%d", id))
	r.Redis.Del(ctx, "products:all")
	r.Redis.Del(ctx, "products:featured")
	if product.Category != "" {
		r.Redis.Del(ctx, fmt.Sprintf("products:category:%s", product.Category))
	}

	return nil
}

func (r *productRepository) Delete(id uint) error {
	var product models.Product
	r.DB.Where("id = ?", id).First(&product)

	err := r.DB.Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}

	ctx := context.Background()
	r.Redis.Del(ctx, fmt.Sprintf("product:%d", id))
	r.Redis.Del(ctx, "products:all")
	r.Redis.Del(ctx, "products:featured")
	if product.Category != "" {
		r.Redis.Del(ctx, fmt.Sprintf("products:category:%s", product.Category))
	}

	return nil
}
