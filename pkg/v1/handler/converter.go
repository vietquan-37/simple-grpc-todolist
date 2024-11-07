package handler

import (
	"github.com/vietquan-37/todo-list/internal/enum"
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/pb"
)

func convertUser(req *pb.CreateUserRequest) model.User {
	return model.User{
		Email:       req.GetEmail(),
		Password:    req.GetPassword(),
		PhoneNumber: req.GetPhoneNumber(),
		FullName:    req.GetFullName(),
		Role:        enum.User,
	}
}
func convertUserResponse(user model.User) *pb.User {
	return &pb.User{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FullName:    user.FullName,
		Role:        string(user.Role),
	}

}
