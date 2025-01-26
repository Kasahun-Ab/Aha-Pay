package services

import (
	"errors"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"go_ecommerce/pkg/dto"
	"sync"

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

func (s *transactionService) CreateWithTransaction(transaction *dto.CreateTransactionRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Get the sender's wallet
		var wallet models.Wallet
		if err := tx.First(&wallet, "id = ?", transaction.WalletID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("wallet not found")
			}
			return err
		}

		// Validate transaction amount
		if transaction.Amount <= 0 {
			return errors.New("transaction amount must be greater than zero")
		}

		var wg sync.WaitGroup          // WaitGroup for concurrent operations
		errChan := make(chan error, 2) // Buffered channel for error handling

		// Handle transaction types
		switch transaction.TransactionType {
		case "WITHDRAWAL":
			if wallet.Balance < transaction.Amount {
				return errors.New("insufficient balance")
			}
			wallet.Balance -= transaction.Amount

		case "DEPOSIT":
			wallet.Balance += transaction.Amount

		case "TRANSFER":
			// Get the receiver's wallet
			var receiverWallet models.Wallet
			if err := tx.First(&receiverWallet, "id = ?", transaction.ReceiverWalletID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("receiver wallet not found")
				}
				return err
			}

			if wallet.Balance < transaction.Amount {
				return errors.New("insufficient balance for transfer")
			}

			// Update balances
			wallet.Balance -= transaction.Amount
			receiverWallet.Balance += transaction.Amount

			// Save the receiver's wallet and create the transfer record concurrently
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := tx.Save(&receiverWallet).Error; err != nil {
					errChan <- errors.New("failed to save receiver wallet")
					return
				}
				transfer := models.Transfer{
					SenderWalletID:   wallet.ID,
					ReceiverWalletID: receiverWallet.ID,
					Amount:           transaction.Amount,
					Status:           "SUCCESS",
				}
				if err := tx.Save(&transfer).Error; err != nil {
					errChan <- errors.New("failed to save transfer")
				}
			}()
		default:
			return errors.New("invalid transaction type")
		}

		// Save the sender's wallet and create the transaction record concurrently
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := tx.Save(&wallet).Error; err != nil {
				errChan <- errors.New("failed to save wallet")
				return
			}
			newTransaction := models.Transaction{
				WalletID:        wallet.ID,
				TransactionType: transaction.TransactionType,
				Amount:          transaction.Amount,
				Status:          "SUCCESS",
			}
			if err := s.repository.Create(tx, &newTransaction); err != nil {
				errChan <- errors.New("failed to create transaction")
			}
		}()

		// Wait for all goroutines to complete
		wg.Wait()
		close(errChan)

		// Check if any errors occurred
		for err := range errChan {
			if err != nil {
				return err
			}
		}

		return nil
	})
}
