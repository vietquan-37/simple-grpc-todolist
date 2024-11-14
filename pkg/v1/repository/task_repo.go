package repository

import (
	"github.com/vietquan-37/todo-list/internal/model"
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
func (repo *TaskRepo) AddTask(task *model.Task) (*model.Task, error) {
	err := repo.DB.Create(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}
