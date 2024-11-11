package handler

import (
	"github.com/vietquan-37/todo-list/internal/enum"
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/internal/pagination"
	"github.com/vietquan-37/todo-list/pb"
)

func convertUser(req *pb.CreateUserRequest) *model.User {
	return &model.User{
		Email:       req.GetEmail(),
		Password:    req.GetPassword(),
		PhoneNumber: req.GetPhoneNumber(),
		FullName:    req.GetFullName(),
		Role:        enum.User,
	}
}
func convertUserUpdate(req *pb.UserUpdateRequest) *model.User {
	return &model.User{
		Password:    req.GetPassword(),
		PhoneNumber: req.GetPhoneNumber(),
		FullName:    req.GetFullName(),
	}
}
func convertUserResponse(user model.User) *pb.UserResponse {
	return &pb.UserResponse{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FullName:    user.FullName,
		Role:        string(user.Role),
	}

}
func convertUserListResponse(res *pagination.Result[model.User]) *pb.UserListResponse {
	var userResponses []*pb.UserResponse
	for _, user := range res.Results {
		userResponse := convertUserResponse(user)
		userResponses = append(userResponses, userResponse)
	}
	pageResponse := &pb.PaginationResponse{
		TotalPage:  res.TotalPage,
		PageNumber: res.PageNumber,
		PageSize:   res.PageSize,
	}
	return &pb.UserListResponse{
		Page:  pageResponse,
		Users: userResponses,
	}
}
