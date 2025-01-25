package models

import (
	"time"
	// "gorm.io/gorm"
)

type User struct {
	ID                    int `gorm:"primaryKey"`
	Username              string
	Email                 string
	PasswordHash          string
	FirstName             string `gorm:""` // Ignore this field when creating or updating a user
	LastName              string
	Status                string
	CreatedAt             time.Time
	ResetToken            string
	ResetTokenExpiry      time.Time
	Wallets               [] Wallet               `gorm:"foreignKey:UserID"`
	PaymentMethods        []PaymentMethod        `gorm:"foreignKey:ID"`
	Logs                  []Log                  `gorm:"foreignKey:ID"`
	UserVerifications     []UserVerification     `gorm:"foreignKey:ID"`
	CardDetails           []CardDetail           `gorm:"foreignKey:ID"`
	RewardPoints          []RewardPoint          `gorm:"foreignKey:ID"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:ID"`
	SecurityLogs          []SecurityLog          `gorm:"foreignKey:ID"`
	Notifications         []Notification         `gorm:"foreignKey:ID"`
	UserSessions          []UserSession          `gorm:"foreignKey:ID"`
	AuditLogs             []AuditLog             `gorm:"foreignKey:ID"`
}
