package models

import (
	"time"
)

type Log struct {
	LogID      int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"index"`
	LogType    string    `gorm:"type:enum('LOGIN', 'PASSWORD_CHANGE', 'SECURITY_ALERT')"`
	LogDetails string    `gorm:"size:255"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}
