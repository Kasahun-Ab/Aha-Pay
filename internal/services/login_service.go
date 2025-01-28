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
	Register(req *dto.RegisterRequest) (dto.RegisterResponse, error)
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
}

type authService struct {
	db         *gorm.DB
	secretKey  string
	Repo       repositories.UserRepository
	WalletRepo repositories.WalletRepository
}

func NewAuthService(
	db *gorm.DB,
	secretKey string,
	userRepo repositories.UserRepository,
	walletRepo repositories.WalletRepository,

) AuthService {
	return &authService{
		db:         db,
		secretKey:  secretKey,
		Repo:       userRepo,
		WalletRepo: walletRepo,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) (dto.RegisterResponse, error) {
	// Check if the user already exists
	_, err := s.Repo.FindByEmail(req.Email)
	if err == nil {
		return dto.RegisterResponse{}, errors.New("user already exists")
	}

	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// Create user
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

	// Generate JWT token
	token, err := utils.GenerateJWT(user, s.secretKey)

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// Create wallet asynchronously
	walletChan := make(chan *models.Wallet)
	errorChan := make(chan error)

	go func() {
		wallet, err := s.WalletRepo.Create(&models.Wallet{Currency: "USD", Status: "active"}, user.ID)
		if err != nil {
			errorChan <- err
			return
		}
		walletChan <- wallet
	}()

	// Wait for wallet creation
	select {
	case wallet := <-walletChan:
		// Success, return response with wallet ID
		return dto.RegisterResponse{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Status:    user.Status,
			Token:     token,
			Wallet:    wallet.ID,
		}, nil
	case err := <-errorChan:
		// Failed to create wallet
		return dto.RegisterResponse{}, err
	}
}

func (s *authService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	// Find user by email
	user, err := s.Repo.FindByEmail(req.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(*user, s.secretKey)
	if err != nil {
		return dto.LoginResponse{}, errors.New("failed to generate token")
	}

	// Return token in response
	return dto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
		Token:    token}, nil
}
