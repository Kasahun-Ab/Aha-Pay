package repositories

import (
	"go_ecommerce/internal/models"

	"gorm.io/gorm"
)

type UserSessionRepo struct {
	db *gorm.DB
}

func NewUserSessionRepository(db *gorm.DB) *UserSessionRepo {
	return &UserSessionRepo{db: db}
}

func (r *UserSessionRepo) Create(userSession *models.UserSession) error {
	return r.db.Create(userSession).Error
}

func (r *UserSessionRepo) FindByID(id int) (*models.UserSession, error) {
	var userSession models.UserSession
	if err := r.db.First(&userSession, id).Error; err != nil {
		return nil, err
	}
	return &userSession, nil
}

func (r *UserSessionRepo) FindAll() ([]models.UserSession, error) {
	var userSessions []models.UserSession
	if err := r.db.Find(&userSessions).Error; err != nil {
		return nil, err
	}
	return userSessions, nil
}

func (r *UserSessionRepo) Update(userSession *models.UserSession) error {
	return r.db.Save(userSession).Error
}

func (r *UserSessionRepo) Delete(id int) error {
	return r.db.Delete(&models.UserSession{}, id).Error
}
func (r *UserSessionRepo) FindByToken(token string) (*models.UserSession, error) {
	var userSession models.UserSession
	if err := r.db.Where("session_token = ?", token).First(&userSession).Error; err != nil {
		return nil, err
	}
	return &userSession, nil
}
