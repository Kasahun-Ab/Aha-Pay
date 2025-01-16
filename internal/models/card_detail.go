package models

import (
	"time"
 
)

type CardDetail struct {
	ID     int       `gorm:"primaryKey;autoIncrement"`
	CardNumber string    `gorm:"size:255"`
	ExpiryDate time.Time `gorm:"size:date"`
	CVV        string    `gorm:"size:255"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	
}
