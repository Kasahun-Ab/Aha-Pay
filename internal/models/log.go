package models

import (
	"time"
)

type Log struct {
	ID      int       `gorm:"primaryKey;autoIncrement"`

	LogType    string    `gorm:"type:enum('LOGIN', 'PASSWORD_CHANGE', 'SECURITY_ALERT')"`
	LogDetails string    `gorm:"size:255"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	
}
