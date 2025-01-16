package models

import (
	"time"
)

type UserVerification struct {
	VerificationID int       `gorm:"primaryKey;autoIncrement"`
	UserID         int       `gorm:"index"`
	VerificationType string   `gorm:"type:enum('EMAIL', 'PHONE', 'ID')"`
	VerificationStatus string `gorm:"type:enum('PENDING', 'VERIFIED', 'FAILED')"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	User           User      `gorm:"foreignKey:UserID"`
}
