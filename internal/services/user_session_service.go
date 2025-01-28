package services

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"go_ecommerce/pkg/dto"
)


type UserSessionService struct {
	repo *repositories.UserSessionRepo
}

func NewUserSessionService(repo repositories.UserSessionRepo) *UserSessionService {
	return &UserSessionService{repo: &repo}
}

func (s *UserSessionService) CreateSession(data dto.CreateUserSessionDTO) (*models.UserSession, error) {
	userSession := &models.UserSession{
		SessionToken: data.SessionToken,
		IPAddress:    data.IPAddress,
		DeviceInfo:   data.DeviceInfo,
	}

	if err := s.repo.Create(userSession); err != nil {
		return nil, err
	}

	return userSession, nil
}

func (s *UserSessionService) GetSessionByID(id int) (*models.UserSession, error) {
	return s.repo.FindByID(id)
}
 
func (s *UserSessionService) GetSessionByToken(token string) (*models.UserSession, error) {
	return s.repo.FindByToken(token)
}


func (s *UserSessionService) GetAllSessions() ([]models.UserSession, error) {
	return s.repo.FindAll()
}

func (s *UserSessionService) UpdateSession(id int, data dto.UpdateUserSessionDTO) error {
	userSession, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	userSession.LastActivity = data.LastActivity
	if data.IsActive != nil {
		userSession.IsActive = *data.IsActive
	}

	return s.repo.Update(userSession)
}

func (s *UserSessionService) DeleteSession(id int) error {
	return s.repo.Delete(id)
}
