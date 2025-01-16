package models

import (
	"time"
)

type UserSession struct {
	ID    int       `gorm:"primaryKey"`
	SessionToken string    `gorm:"size:255"`
	IPAddress    string    `gorm:"size:255"`
	DeviceInfo   string    `gorm:"size:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	LastActivity time.Time `gorm:"autoUpdateTime"`
	IsActive     bool      `gorm:"default:true"`
}
