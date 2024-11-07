package model

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID       uint // Foreign key referring to User
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool `gorm:"default:false"`
	ExpiredAt    time.Time
}
