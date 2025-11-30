package services

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/repositories"
)

type OrderServices struct {
	Repo repositories.OrderRepository
}

func NewOrderServices(repo repositories.OrderRepository) *OrderServices {
	return &OrderServices{
		Repo: repo,
	}
}

func (s *OrderServices) GetOrderByID(id uint) (*models.Order, error) {
	return s.Repo.GetOrderByID(id)
}

func (s *OrderServices) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.Repo.GetOrdersByUserID(userID)
}

func (s *OrderServices) GetAllOrders() ([]models.Order, error) {
	return s.Repo.GetAllOrders()
}

func (s *OrderServices) Create(order *models.Order) error {
	return s.Repo.Create(order)
}

func (s *OrderServices) UpdateStatus(id uint, status string) error {
	return s.Repo.UpdateStatus(id, status)
}

func (s *OrderServices) CreateOrderWithPayment(order *models.Order) (*models.Order, error) {
	if err := s.Repo.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}
