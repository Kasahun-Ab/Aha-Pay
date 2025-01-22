package services

import (
	"errors"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"go_ecommerce/pkg/dto"
	"go_ecommerce/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req dto.RegisterRequest) (dto.RegisterResponse, error)

	Login(req dto.LoginRequest) (dto.LoginResponse, error)

	// ValidateToken(token string) (map[string]interface{}, error)
}

type authService struct {
	db        *gorm.DB
	secretKey string
	Repo      repositories.UserRepository
}

func NewAuthService(db *gorm.DB, secretKey string, Repo repositories.UserRepository) AuthService {

	return &authService{db: db, secretKey: secretKey, Repo: Repo}

}

func (s *authService) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	
	_, err := s.Repo.FindByEmail(req.Email)

	if err != nil {
		return dto.RegisterResponse{}, errors.New("user already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err
	}
	user := models.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Username:     req.Username,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Status:       "active",
	}

	if err := s.Repo.Create(&user); err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{ID: user.ID}, nil

}

func (s *authService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {

	user, err := s.Repo.FindByEmail(req.Email)

	if err != nil {

		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(*user, s.secretKey)

	if err != nil {
		
		return dto.LoginResponse{}, errors.New("failed to generate token")
	}

	return dto.LoginResponse{Token: token}, nil

}
