// repositories/transaction_repository.go
package repositories

import (
	"go_ecommerce/pkg/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, transaction *dto.CreateTransactionRequest) error
	GetByID(tx *gorm.DB, id int) (*dto.TransactionResponse, error)
}

type transactionRepository struct {}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *dto.CreateTransactionRequest) error {
	return tx.Create(transaction).Error
}

func (r *transactionRepository) GetByID(tx *gorm.DB, id int) (*dto.TransactionResponse, error) {
	var transaction dto.TransactionResponse
	err := tx.Preload("Wallet").First(&transaction, id).Error
	return &transaction, err
}
