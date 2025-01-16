package models

import (
	"time"
)

type PaymentMethod struct {
	ID int       `gorm:"primaryKey;autoIncrement"`
	MethodType      string    `gorm:"type:enum('CREDIT_CARD', 'DEBIT_CARD', 'BANK_TRANSFER', 'PAYPAL')"`
	MethodDetails   string    `gorm:"size:255"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
