package interfaces

import (
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/internal/pagination"
)

type TaskRepo interface {
	AddTask(task *model.Task) (*model.Task, error)
	GetUserTask(ID int, pageNumber, pageSize int64) (*pagination.Result[model.Task], error)
	UpdateTask(id int, task *model.Task) (*model.Task, error)
	DeleteTask(*model.Task) error
	GetTaskById(id int) (*model.Task, error)
}
