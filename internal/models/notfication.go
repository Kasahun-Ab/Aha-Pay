package models

import (
	"time"
)
type Notification struct {
    NotificationID   int       `gorm:"primaryKey;autoIncrement"`
    UserID           int       `gorm:"index"`
    Message          string    `gorm:"type:text"`
    NotificationType string    `gorm:"type:enum('TRANSACTION', 'SECURITY', 'PROMOTIONAL')"`
    IsRead           bool      `gorm:"default:false"`
    CreatedAt        time.Time `gorm:"autoCreateTime"`
    User             User      `gorm:"foreignKey:UserID;references:UserID"`
}