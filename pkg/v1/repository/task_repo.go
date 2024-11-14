package repository

import (
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"
	"gorm.io/gorm"
)

type TaskRepo struct {
	DB *gorm.DB
}

func NewTaskRepo(db *gorm.DB) interfaces.TaskRepo {
	return &TaskRepo{
		DB: db,
	}
}
