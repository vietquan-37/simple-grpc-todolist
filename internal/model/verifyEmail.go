package model

import (
	"time"

	"gorm.io/gorm"
)

type VerifyEmail struct {
	gorm.Model
	UserID     uint
	SecretCode string
	CreateAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ExpiredAt  time.Time `gorm:"default:CURRENT_TIMESTAMP + interval '15 minutes'"`
}
