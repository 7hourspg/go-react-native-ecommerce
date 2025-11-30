package services

import (
	"go-ecommerce-api/models"
	"go-ecommerce-api/repositories"
)

type UserServices interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User, id uint) error
	Delete(id uint) error
	GetUserByEmail(e string) (*models.User, error)
}

type userServices struct {
	repo repositories.UserRepositories
}

func NewUserServices(s repositories.UserRepositories) UserServices {
	return &userServices{
		repo: s,
	}
}

func (s userServices) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s userServices) GetByID(id uint) (*models.User, error) {

	return s.repo.GetByID(id)
}

func (s userServices) GetUserByEmail(e string) (*models.User, error) {

	return s.repo.GetUserByEmail(e)
}

func (s userServices) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s userServices) Update(user *models.User, id uint) error {
	return s.repo.Update(user, id)
}

func (s userServices) Delete(id uint) error {
	return s.repo.Delete(id)
}
