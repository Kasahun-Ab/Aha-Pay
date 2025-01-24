// repositories/transaction_repository.go
package repositories

import (
	"gorm.io/gorm"
	"go_ecommerce/internal/models"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, transaction *models.Transaction) error
	GetByID(tx *gorm.DB, id int) (*models.Transaction, error)
}

type transactionRepository struct {}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *models.Transaction) error {
	return tx.Create(transaction).Error
}

func (r *transactionRepository) GetByID(tx *gorm.DB, id int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := tx.Preload("Wallet").First(&transaction, id).Error
	return &transaction, err
}
