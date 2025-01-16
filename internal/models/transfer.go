package models

import (
	"time"
)

type Transfer struct {
	TransferID      int       `gorm:"primaryKey;autoIncrement"`
	SenderWalletID  int       `gorm:"index"`
	ReceiverWalletID int      `gorm:"index"`
	Amount          float64   `gorm:"type:decimal(15,2)"`
	Status          string    `gorm:"type:enum('PENDING', 'COMPLETED', 'FAILED')"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	SenderWallet    Wallet    `gorm:"foreignKey:SenderWalletID"`
	ReceiverWallet  Wallet    `gorm:"foreignKey:ReceiverWalletID"`
}
