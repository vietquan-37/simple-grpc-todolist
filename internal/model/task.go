package model

import (
	"github.com/vietquan-37/todo-list/internal/enum"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName string
	Status   enum.Status
	UserID   uint
}
