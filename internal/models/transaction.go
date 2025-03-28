package models

import "time"

type Transaction struct {
	ID int `gorm:"primaryKey;autoIncrement"`

	WalletID int `gorm:"not null"`

	Amount float64 `gorm:"type:decimal(15,2)"`

	TransactionType string `gorm:"type:enum('DEPOSIT', 'WITHDRAWAL', 'TRANSFER')"`

	Status string `gorm:"type:enum('PENDING', 'COMPLETED', 'FAILED')"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}
