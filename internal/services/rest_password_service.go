package services

import (
	"errors"
	"fmt"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"go_ecommerce/pkg/utils"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ResetService struct {
	Repo repositories.UserRepository
}

func NewResetService(repo repositories.UserRepository) *ResetService {
	return &ResetService{Repo: repo}
}


func GenerateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000)) 
}

func (s *ResetService) RequestPasswordReset(email string) error {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return errors.New("email not found")
	}


	otp := GenerateRandomOTP()
	user.ResetToken = otp
	user.ResetTokenExpiry = time.Now().Add(15 * time.Minute) // OTP expires in 15 minutes

	if err := s.Repo.Update(user); err != nil {
		return errors.New("failed to generate reset token")
	}


	resetMessage := fmt.Sprintf("Your password reset OTP is: %s\nIt will expire in 15 minutes.", otp)
	go utils.SendEmail(user.Email, "Password Reset OTP", resetMessage)

	return nil
}

func (s *ResetService) ValidateResetToken(token string) (*models.User, error) {
	user, err := s.Repo.FindByResetToken(token)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}
	return user, nil
}

func (s *ResetService) ResetPassword(token, newPassword string) error {
	user, err := s.ValidateResetToken(token)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.PasswordHash = string(hashedPassword)
	user.ResetToken = ""
	user.ResetTokenExpiry = time.Time{} 
	if err := s.Repo.Update(user); err != nil {
		return errors.New("failed to reset password")
	}

	return nil
}
