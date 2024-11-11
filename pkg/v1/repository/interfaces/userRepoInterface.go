package interfaces

import (
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/internal/pagination"
)

type UserRepo interface {
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(id int, user *model.User) (*model.User, error)
	GetUser(id int) (*model.User, error)
	DeleteUser(id int) error
	GetUserByEmail(email string) (*model.User, error)
	GetAllUser(name string, pageNumber, pageSize int64) (*pagination.Result[model.User], error)
}
