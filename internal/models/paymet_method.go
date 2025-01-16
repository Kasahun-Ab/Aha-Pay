package models

import (
	"time"
)

type PaymentMethod struct {
	PaymentMethodID int       `gorm:"primaryKey;autoIncrement"`
	UserID          int       `gorm:"index"`
	MethodType      string    `gorm:"type:enum('CREDIT_CARD', 'DEBIT_CARD', 'BANK_TRANSFER', 'PAYPAL')"`
	MethodDetails   string    `gorm:"size:255"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	User            User      `gorm:"foreignKey:UserID"`
}
