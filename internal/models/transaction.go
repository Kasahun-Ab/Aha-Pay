package models

import (
	"time"
)

type Transaction struct {
	TransactionID int       `gorm:"primaryKey;autoIncrement"`
	WalletID      int       `gorm:"index"`
	Amount        float64   `gorm:"type:decimal(15,2)"`
	TransactionType string   `gorm:"type:enum('DEPOSIT', 'WITHDRAWAL', 'TRANSFER')"`
	Status        string    `gorm:"type:enum('PENDING', 'COMPLETED', 'FAILED')"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	Wallet        Wallet    `gorm:"foreignKey:WalletID"`
}
