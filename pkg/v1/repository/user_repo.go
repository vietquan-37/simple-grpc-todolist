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
func (repo *UserRepo) UpdateUser(id int, user *model.User) (*model.User, error) {

	var existingUser model.User
	if err := repo.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}
	if err := repo.DB.Model(&existingUser).Updates(user).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}
func (repo *UserRepo) GetUser(id int) (*model.User, error) {
	var dbUser model.User

	err := repo.DB.Where("id = ?", id).First(&dbUser).Error

	return &dbUser, err
}
func (repo *UserRepo) DeleteUser(id int) error {
	var existingUser model.User
	if err := repo.DB.First(&existingUser, id).Error; err != nil {
		return err
	}
	err := repo.DB.Delete(&existingUser).Error
	if err != nil {
		return err
	}
	return nil

}

func (repo *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var dbUser model.User
	err := repo.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
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
