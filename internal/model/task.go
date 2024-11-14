package model

import (
	"time"

	"github.com/vietquan-37/todo-list/internal/enum"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName     string
	Description  string
	Status       enum.Status
	CreateAt     time.Time
	TaskDeadline time.Time
	UserID       uint
}
