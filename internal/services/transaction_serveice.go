// services/transaction_service.go
package services

import (
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"go_ecommerce/pkg/dto"

	"gorm.io/gorm"
)

type TransactionService interface {
	CreateWithTransaction(transaction *dto.CreateTransactionRequest) error
}

type transactionService struct {
	db         *gorm.DB
	repository repositories.TransactionRepository
}

func NewTransactionService(db *gorm.DB, repo repositories.TransactionRepository) TransactionService {
	return &transactionService{db: db, repository: repo}
}

// CreateWithTransaction handles creating a transaction within a GORM transaction
func (s *transactionService) CreateWithTransaction(transaction *dto.CreateTransactionRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Example: Check the wallet balance
		var wallet models.Wallet
		if err := tx.First(&wallet, transaction.WalletID).Error; err != nil {
			return err
		}
		// Simulate insufficient balance check for withdrawal
		if transaction.TransactionType == "WITHDRAWAL" && wallet.Balance < transaction.Amount {
			return gorm.ErrInvalidData // Insufficient funds
		}

		// Deduct balance from wallet (if applicable)
		if transaction.TransactionType == "WITHDRAWAL" {
			wallet.Balance -= transaction.Amount
			if err := tx.Save(&wallet).Error; err != nil {
				return err
			}
		}

		// Create the transaction record
		if err := s.repository.Create(tx, transaction); err != nil {
			return err
		}

		// Commit the transaction if no errors
		return nil
	})
}
