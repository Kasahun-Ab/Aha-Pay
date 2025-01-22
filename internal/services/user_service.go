package services

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
)

type UserAccountService struct {
	Repo repositories.UserRepository
}

func NewUserAccountService(repo repositories.UserRepository) *UserAccountService {

	return &UserAccountService{Repo: repo}

}

func (s *UserAccountService) FindByEmail(email string) (*models.User, error) {

	return s.Repo.FindByEmail(email)

}

func (s *UserAccountService) FindByID(id int) (*models.User, error) {

	return s.Repo.FindByID(id)

}

func (s *UserAccountService) Create(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserAccountService) Update(user *models.User) error {
	return s.Repo.Update(user)
}

func (s *UserAccountService) DeleteUser(user *models.User) error {
	return s.Repo.Delete(user)
}
