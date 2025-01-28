package services

import (
	"errors"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"net/http"

	"go_ecommerce/pkg/dto"
	"go_ecommerce/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) (dto.RegisterResponse, error, *http.Cookie)
	Login(req dto.LoginRequest) (dto.LoginResponse, error, *http.Cookie)
}

type authService struct {
	db          *gorm.DB
	secretKey   string
	Repo        repositories.UserRepository
	WalletRepo  repositories.WalletRepository
	sessionRepo repositories.UserSessionRepo
}

func NewAuthService(db *gorm.DB, secretKey string, Repo repositories.UserRepository, walletRepo repositories.WalletRepository, session repositories.UserSessionRepo) AuthService {
	return &authService{db: db, secretKey: secretKey, Repo: Repo, WalletRepo: walletRepo, sessionRepo: session}
}

func (s *authService) Register(req *dto.RegisterRequest) (dto.RegisterResponse, error, *http.Cookie) {

	_, err := s.Repo.FindByEmail(req.Email)

	if err == nil {

		return dto.RegisterResponse{}, errors.New("user already exists"), nil
	}
	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err, nil
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

		return dto.RegisterResponse{}, err, nil
	}

	token, err := utils.GenerateJWT(user, s.secretKey)
	if err != nil {
		return dto.RegisterResponse{}, err, nil
	}

	cookie := utils.SetCookie("token", token, 3600)
  
	var wallet *models.Wallet
	var walletErr error

	done := make(chan struct{})
	go func() {
		defer close(done)
		wallet, walletErr = s.WalletRepo.Create(&models.Wallet{Currency: "USD", Status: "active"}, user.ID)
	}()

	<-done

	if walletErr != nil {
		return dto.RegisterResponse{}, walletErr, nil
	}

	return dto.RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		Wallet:    wallet.ID,
	}, nil, cookie

}

func (s *authService) Login(req dto.LoginRequest) (dto.LoginResponse, error, *http.Cookie) {

	user, err := s.Repo.FindByEmail(req.Email)

	if err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials"), nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {

		return dto.LoginResponse{}, errors.New("invalid credentials"), nil

	}

	token, err := utils.GenerateJWT(*user, s.secretKey)
	if err != nil {
		return dto.LoginResponse{}, err, nil
	}
	cookie := utils.SetCookie("token", token, 3600)

	if err != nil {

		return dto.LoginResponse{}, errors.New("failed to generate token"), nil

	}

	return dto.LoginResponse{Token: token}, nil, cookie
}
