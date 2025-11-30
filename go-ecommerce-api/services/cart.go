package services

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/repositories"
)

type CartServices struct {
	Repo repositories.CartRepository
}

func NewCartServices(repo repositories.CartRepository) *CartServices {
	return &CartServices{
		Repo: repo,
	}
}

func (s *CartServices) GetCartByUserID(userID uint) (*models.Cart, error) {
	return s.Repo.GetCartByUserID(userID)
}

func (s *CartServices) AddItem(cartID uint, productID uint, quantity int) (*models.CartItem, error) {
	return s.Repo.AddItem(cartID, productID, quantity)
}

func (s *CartServices) UpdateItemQuantity(id uint, quantity int) error {
	return s.Repo.UpdateItemQuantity(id, quantity)
}

func (s *CartServices) RemoveItem(id uint) error {
	return s.Repo.RemoveItem(id)
}

func (s *CartServices) ClearCart(cartID uint) error {
	return s.Repo.ClearCart(cartID)
}
