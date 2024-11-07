package interfaces

import "github.com/vietquan-37/todo-list/internal/model"

type UserUseCase interface {
	CreateUser(model.User) (*model.User, error)
}
