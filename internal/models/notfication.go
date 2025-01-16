package models

import (
	"time"
)
type Notification struct {
    ID   int       `gorm:"primaryKey;autoIncrement"`
    Message          string    `gorm:"type:text"`
    NotificationType string    `gorm:"type:enum('TRANSACTION', 'SECURITY', 'PROMOTIONAL')"`
    IsRead           bool      `gorm:"default:false"`
    CreatedAt        time.Time `gorm:"autoCreateTime"`
}