package model

import (
	"github.com/vietquan-37/todo-list/internal/enum"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	FullName     string
	PhoneNumber  string
	Password     string
	Role         enum.Role
	Verified     bool          `gorm:"default:false"`
	Tasks        []Task        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	VerifyEmails []VerifyEmail `gorm:"foreignKey:UserID"`
	Sessions     []Session     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}