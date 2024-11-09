package repository

import (
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/internal/pagination"
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{
		DB: db,
	}
}
func (repo *UserRepo) CreateUser(user *model.User) (*model.User, error) {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (repo *UserRepo) UpdateUser(user *model.User) error {
	err := repo.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
	return err
}

func (repo *UserRepo) GetUser(id int) (*model.User, error) {
	var dbUser model.User

	err := repo.DB.Where("id = ?", id).First(&dbUser).Error

	return &dbUser, err
}
func (repo *UserRepo) DeleteUser(id int) error {
	err := repo.DB.Where("id = ?", id).Delete(&model.User{}).Error
	return err
}
func (repo *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var dbUser model.User
	err := repo.DB.Where("email = ?", email).First(&dbUser).Error
	return &dbUser, err
}
func (repo *UserRepo) GetAllUser(name string, pageNumber, pageSize int64) (*pagination.Result[model.User], error) {

	query := repo.DB.Model(&model.User{})
	if name != "" {
		query = query.Where("full_name LIKE ?", "%"+name+"%")
	}
	result, err := pagination.Paginate[model.User](query, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	return result, nil
}
