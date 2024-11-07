package interfaces

import "github.com/vietquan-37/todo-list/internal/model"

type UserRepo interface {
	CreateUser(user model.User) (*model.User, error)
	UpdateUser(user model.User) error
	GetUser(id int) (*model.User, error)
	DeleteUser(id int) error
	GetUserByEmail(email string) (*model.User, error)
}
