package models

import (
	"time"
)

type RecurringTransaction struct {
	ID int       `gorm:"primaryKey;autoIncrement"`
	Amount      float64   `gorm:"type:decimal(15,2)"`
	Frequency   string    `gorm:"type:enum('DAILY', 'WEEKLY', 'MONTHLY')"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
