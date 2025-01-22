package services

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) FindByEmail(email string) (*models.User, error) {
	return s.Repo.FindByEmail(email)
}

func (s *UserService) FindByID(id int) (*models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) Create(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) Update(user *models.User) error {
	return s.Repo.Update(user)
}
