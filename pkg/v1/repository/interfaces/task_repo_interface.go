package interfaces

import "github.com/vietquan-37/todo-list/internal/model"

type TaskRepo interface {
	AddTask(*model.Task) (*model.Task, error)
}
