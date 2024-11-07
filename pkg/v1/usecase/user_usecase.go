package usecase

import (
	"errors"

	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"
	userCase "github.com/vietquan-37/todo-list/pkg/v1/usecase/interfaces"
	"gorm.io/gorm"
)

type userUseCase struct {
	Repo interfaces.UserRepo
}

func NewUserCase(repo interfaces.UserRepo) userCase.UserUseCase {
	return &userUseCase{
		Repo: repo,
	}
}
func (uc *userUseCase) CreateUser(user model.User) (*model.User, error) {
	if _, err := uc.Repo.GetUserByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("the email has been register before")
	}
	return uc.Repo.CreateUser(user)
}
