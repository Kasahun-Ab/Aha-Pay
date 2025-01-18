package services

import (
	"errors"
	"go_ecommerce/internal/models"
	"go_ecommerce/pkg/dto"
	"go_ecommerce/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req dto.RegisterRequest) (dto.RegisterResponse, error)

	Login(req dto.LoginRequest) (dto.LoginResponse, error)

	ValidateToken(token string) (map[string]interface{}, error)
}

type authService struct {
	db        *gorm.DB
	secretKey string
}

func NewAuthService(db *gorm.DB, secretKey string) AuthService {

	return &authService{db: db, secretKey: secretKey}

}

func (s *authService) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	var existingUser models.User

	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {

		return dto.RegisterResponse{}, errors.New("user already exists")

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	user := models.User{
		Username:     req.Username,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}
	if err := s.db.Create(&user).Error; err != nil {

		return dto.RegisterResponse{}, err
	}

	token, err := utils.GenerateJWT(user, s.secretKey)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     token,
	}, nil
}

func (s *authService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {

	var user models.User
	if err := s.db.Where("email = ?", req.Email).First(&user).Error; err != nil {

		return dto.LoginResponse{}, errors.New("User Not Found")

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {

		return dto.LoginResponse{}, errors.New("Incorrect Password")

	}

	token, err := utils.GenerateJWT(user, s.secretKey)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{

		Token: token}, nil
}

func (s *authService) ValidateToken(token string) (map[string]interface{}, error) {

	claims, err := utils.ParseToken(token, s.secretKey)

	if err != nil {

		return nil, errors.New("invalid or expired token")

	}

	return claims, nil
}
