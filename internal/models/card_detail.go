package models

import (
	"time"
)

type CardDetail struct {
	CardID     int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"index"`
	CardNumber string    `gorm:"size:255"`
	ExpiryDate time.Time `gorm:"size:date"`
	CVV        string    `gorm:"size:255"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}
