package models

import (
	"time"
)

type RecurringTransaction struct {
	RecurringID int       `gorm:"primaryKey;autoIncrement"`
	UserID      int       `gorm:"index"`
	Amount      float64   `gorm:"type:decimal(15,2)"`
	Frequency   string    `gorm:"type:enum('DAILY', 'WEEKLY', 'MONTHLY')"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	User        User      `gorm:"foreignKey:UserID"`
}
