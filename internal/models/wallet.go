package models

import (
	"time"
)

type Wallet struct {
    WalletID   int        `gorm:"primaryKey;autoIncrement"`
    UserID     int        `gorm:"index"`
    Currency   string     `gorm:"size:255"`
    Balance    float64    `gorm:"type:decimal(15,2);default:0"`
    Status     string     `gorm:"type:enum('ACTIVE', 'INACTIVE', 'LOCKED');default:'ACTIVE'"`
    CreatedAt  time.Time  `gorm:"autoCreateTime"`
    User       User       `gorm:"foreignKey:UserID;references:UserID"`
    Transactions []Transaction `gorm:"foreignKey:WalletID"`
    Transfers  []Transfer `gorm:"foreignKey:SenderWalletID"`
    TransfersReceived []Transfer `gorm:"foreignKey:ReceiverWalletID"`
}