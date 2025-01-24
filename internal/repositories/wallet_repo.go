// repository/wallet_repository.go
package repositories

import (
	"go_ecommerce/internal/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {

	return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(wallet *models.Wallet, userId int) (*models.Wallet, error) {

	wallet.UserID = userId

	if err := r.db.Create(wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *WalletRepository) GetByID(id int) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.First(&wallet, id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) Update(wallet *models.Wallet) (*models.Wallet, error) {
	if err := r.db.Save(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *WalletRepository) Delete(id int) error {
	if err := r.db.Delete(&models.Wallet{}, id).Error; err != nil {
		return err
	}
	return nil
}
