package models

import (
	"time"
)

type AuditLog struct {
	ID    int       `gorm:"primaryKey;autoIncrement"`
	ActionType string    `gorm:"type:enum('CREATE', 'UPDATE', 'DELETE', 'TRANSFER', 'DEPOSIT', 'WITHDRAWAL')"`
	TableName  string    `gorm:"size:255"`
	RecordID   int       `gorm:"type:int"`
	OldValue   string    `gorm:"size:text"`
	NewValue   string    `gorm:"size:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	
}
