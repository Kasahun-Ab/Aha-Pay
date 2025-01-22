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

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// GenerateRandomOTP generates a random 6-digit OTP
func GenerateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000)) // Generates a random 6-digit OTP
}

func (s *UserService) RequestPasswordReset(email string) error {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return errors.New("email not found")
	}

	// Generate a 6-digit OTP and set the expiration time
	otp := GenerateRandomOTP()
	user.ResetToken = otp
	user.ResetTokenExpiry = time.Now().Add(15 * time.Minute) // OTP expires in 15 minutes

	if err := s.Repo.Update(user); err != nil {
		return errors.New("failed to generate reset token")
	}

	// Send the OTP to the user's email (abstracted for simplicity)
	resetMessage := fmt.Sprintf("Your password reset OTP is: %s\nIt will expire in 15 minutes.", otp)
	go utils.SendEmail(user.Email, "Password Reset OTP", resetMessage)

	return nil
}

func (s *UserService) ValidateResetToken(token string) (*models.User, error) {
	user, err := s.Repo.FindByResetToken(token)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}
	return user, nil
}

func (s *UserService) ResetPassword(token, newPassword string) error {
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
	user.ResetTokenExpiry = time.Time{} // Clear the token and expiry

	if err := s.Repo.Update(user); err != nil {
		return errors.New("failed to reset password")
	}

	return nil
}
