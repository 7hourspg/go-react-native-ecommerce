package services

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/repositories"
)

type WishlistServices struct {
	Repo repositories.WishlistRepository
}

func NewWishlistServices(repo repositories.WishlistRepository) *WishlistServices {
	return &WishlistServices{
		Repo: repo,
	}
}

func (s *WishlistServices) GetWishlistByUserID(userID uint) ([]models.Wishlist, error) {
	return s.Repo.GetWishlistByUserID(userID)
}

func (s *WishlistServices) AddToWishlist(userID uint, productID uint) error {
	return s.Repo.AddToWishlist(userID, productID)
}

func (s *WishlistServices) RemoveFromWishlist(userID uint, productID uint) error {
	return s.Repo.RemoveFromWishlist(userID, productID)
}

func (s *WishlistServices) IsInWishlist(userID uint, productID uint) (bool, error) {
	return s.Repo.IsInWishlist(userID, productID)
}

