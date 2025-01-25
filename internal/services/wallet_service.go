
package services

import (
	"go_ecommerce/internal/models"
	// "go_ecommerce/internal/repositories"
	"go_ecommerce/internal/repositories"
	"errors"
)


type WalletService struct {
	repo repositories.WalletRepository
}


func NewWalletService(repo repositories.WalletRepository) *WalletService {
	return &WalletService{repo}
}


func (s *WalletService) CreateWallet(wallet *models.Wallet,userId int) (*models.Wallet, error) {
	

	if wallet.Currency == "" || wallet.Status == "" {
		return nil, errors.New("currency and status are required")
	}

   
	return s.repo.Create(wallet , userId)
}



func (s *WalletService) GetWalletByID(id int) (*models.Wallet, error) {

	return s.repo.GetByID(id)
}


func (s *WalletService) UpdateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	
	return s.repo.Update(wallet)
}


func (s *WalletService) DeleteWallet(id int) error {
	return s.repo.Delete(id)
}

func (s *WalletService) GetAllWallet(id int) ([]models.Wallet, error) {

	return s.repo.GetAllWalletsByUserID(id)
}
