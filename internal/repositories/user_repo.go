package repositories

import (
	// "errors"
	"go_ecommerce/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByResetToken(token string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("reset_token = ? AND reset_token_expiry > ?", token, time.Now()).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	if err := r.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Create(user *models.User) error {

	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {

	var user models.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
