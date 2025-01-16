package models

import (
	"time"
)

type RewardPoint struct {
	RewardID   int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"index"`
	Points     int       `gorm:"type:int"`
	RewardType string    `gorm:"type:enum('CASHBACK', 'DISCOUNT', 'BONUS')"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}
