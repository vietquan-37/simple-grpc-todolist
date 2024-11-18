package repository

import (
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/internal/pagination"
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
func (repo *TaskRepo) GetUserTask(ID int, pageNumber, pageSize int64) (*pagination.Result[model.Task], error) {
	query := repo.DB.Model(&model.Task{}).Where("user_id= ?", ID)
	result, err := pagination.Paginate[model.Task](query, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (repo *TaskRepo) UpdateTask(id int, task *model.Task) (*model.Task, error) {
	var existingTask model.Task
	if err := repo.DB.First(&existingTask, id).Error; err != nil {
		return nil, err
	}
	if err := repo.DB.Model(&existingTask).Updates(task).Error; err != nil {
		return nil, err
	}

	return &existingTask, nil
}
func (repo *TaskRepo) DeleteTask(task *model.Task) error {
	//this is hard delete
	if err := repo.DB.Unscoped().Delete(&task).Error; err != nil {
		return err
	}
	return nil
}
func (repo *TaskRepo) GetTaskById(id int) (*model.Task, error) {
	var task model.Task

	err := repo.DB.Where("id = ?", id).First(&task).Error

	return &task, err
}
