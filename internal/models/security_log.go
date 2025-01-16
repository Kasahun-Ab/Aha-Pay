package models

import (
	"time"
)

type SecurityLog struct {
	LogID        int       `gorm:"primaryKey;autoIncrement"`
	UserID       int       `gorm:"index"`
	ActionType   string    `gorm:"type:enum('FAILED_LOGIN', 'MFA_CHANGE', 'SUSPICIOUS_ACTIVITY', 'PASSWORD_RESET')"`
	IPAddress    string    `gorm:"size:255"`
	DeviceInfo   string    `gorm:"size:text"`
	ActionDetails string   `gorm:"size:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	User         User      `gorm:"foreignKey:UserID"`
}
