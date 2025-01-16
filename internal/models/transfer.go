package models

import (
	"time"
)

type Transfer struct {
	ID      int       `gorm:"primaryKey;autoIncrement"`
	// SenderWalletID  int       `gorm:"index"`
	// ReceiverWalletID int      `gorm:"index"`
	Amount          float64   `gorm:"type:decimal(15,2)"`
	Status          string    `gorm:"type:enum('PENDING', 'COMPLETED', 'FAILED')"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	SenderWallet    Wallet    `gorm:"foreignKey:ID"`
	ReceiverWallet  Wallet    `gorm:"foreignKey:ID"`
}
