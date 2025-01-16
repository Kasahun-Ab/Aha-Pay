package models

import (
	"time"
)

type UserVerification struct {
	ID int       `gorm:"primaryKey;autoIncrement"`
	VerificationType string   `gorm:"type:enum('EMAIL', 'PHONE', 'ID')"`
	VerificationStatus string `gorm:"type:enum('PENDING', 'VERIFIED', 'FAILED')"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}
