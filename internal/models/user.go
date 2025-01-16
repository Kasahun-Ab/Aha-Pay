package models

import (
	"time"
	// "gorm.io/gorm"
)

type User struct {
	UserID                int                    `gorm:"primaryKey;autoIncrement"`
	Username              string                 `gorm:"size:255"`
	Email                 string                 `gorm:"size:255"`
	PasswordHash          string                 `gorm:"size:255"`
	FirstName             string                 `gorm:"size:255"`
	LastName              string                 `gorm:"size:255"`
	Status                string                 `gorm:"type:enum('ACTIVE', 'INACTIVE', 'SUSPENDED');default:'ACTIVE'"`
	CreatedAt             time.Time              `gorm:"autoCreateTime"`
	Wallets               []Wallet               `gorm:"foreignKey:UserID"`
	PaymentMethods        []PaymentMethod        `gorm:"foreignKey:UserID"`
	Logs                  []Log                  `gorm:"foreignKey:UserID"`
	UserVerifications     []UserVerification     `gorm:"foreignKey:UserID"`
	CardDetails           []CardDetail           `gorm:"foreignKey:UserID"`
	RewardPoints          []RewardPoint          `gorm:"foreignKey:UserID"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:UserID"`
	SecurityLogs          []SecurityLog          `gorm:"foreignKey:UserID"`
	Notifications         []Notification         `gorm:"foreignKey:UserID;references:UserID"`
	UserSessions          []UserSession          `gorm:"foreignKey:UserID"`
	AuditLogs             []AuditLog             `gorm:"foreignKey:UserID"`
}
