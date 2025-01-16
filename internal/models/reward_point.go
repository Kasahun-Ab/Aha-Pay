package models

import (
	"time"
)

type RewardPoint struct {
	ID   int       `gorm:"primaryKey;autoIncrement"`
	Points     int       `gorm:"type:int"`
	RewardType string    `gorm:"type:enum('CASHBACK', 'DISCOUNT', 'BONUS')"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
