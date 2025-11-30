package services

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/repositories"
)

type ProductServices struct {
	Repo repositories.ProductRepository
}

func NewProductServices(repo repositories.ProductRepository) *ProductServices {
	return &ProductServices{
		Repo: repo,
	}
}

func (s *ProductServices) GetProductByID(id uint) (*models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductServices) GetAll() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductServices) GetFeatured() ([]models.Product, error) {
	return s.Repo.GetFeatured()
}

func (s *ProductServices) GetByCategory(category string) ([]models.Product, error) {
	return s.Repo.GetByCategory(category)
}

func (s *ProductServices) Search(query string) ([]models.Product, error) {
	return s.Repo.Search(query)
}

func (s *ProductServices) Create(product *models.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return s.Repo.Create(product)
}

func (s *ProductServices) BulkCreate(products []models.Product) ([]models.Product, error) {
	return s.Repo.BulkCreate(products)
}

func (s *ProductServices) Update(id uint, product *models.Product) error {
	return s.Repo.Update(id, product)
}

func (s *ProductServices) Delete(id uint) error {
	return s.Repo.Delete(id)
}

