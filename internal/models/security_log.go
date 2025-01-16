package models

import (
	"time"
)

type SecurityLog struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	ActionType   string    `gorm:"type:enum('FAILED_LOGIN', 'MFA_CHANGE', 'SUSPICIOUS_ACTIVITY', 'PASSWORD_RESET')"`
	IPAddress    string    `gorm:"size:255"`
	DeviceInfo   string    `gorm:"size:text"`
	ActionDetails string   `gorm:"size:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
