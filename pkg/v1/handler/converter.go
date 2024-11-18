package handler

import (
	"time"

	"github.com/vietquan-37/todo-list/internal/enum"
	"github.com/vietquan-37/todo-list/internal/model"
	"github.com/vietquan-37/todo-list/internal/pagination"
	"github.com/vietquan-37/todo-list/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
func convertTask(req *pb.CreateTaskRequest) *model.Task {
	return &model.Task{
		TaskName:     req.GetTaskName(),
		Description:  req.GetDescription(),
		CreateAt:     time.Now(),
		Status:       enum.Pending,
		TaskDeadline: req.GetTaskDeadline().AsTime(),
	}
}
func convertTaskUpdate(req *pb.UpdateTaskRequest) *model.Task {
	return &model.Task{
		TaskName:    req.GetTaskName(),
		Description: req.GetDescription(),
		Status:      enum.Status(req.GetStatus().String()),
	}
}
func convertTaskResponse(task model.Task) *pb.TaskResponse {
	return &pb.TaskResponse{
		TaskName:     task.TaskName,
		Desctiption:  task.Description,
		TaskStatus:   string(task.Status),
		CreatedAt:    timestamppb.New(task.CreateAt),
		TaskDeadline: timestamppb.New(task.TaskDeadline),
	}
}
func convertUserListResponse(res *pagination.Result[model.User]) *pb.UserListResponse {
	var userResponses []*pb.UserResponse
	for _, user := range res.Results {
		userResponse := convertUserResponse(user)
		userResponses = append(userResponses, userResponse)
	}
	pageResponse := &pb.PaginationResponse{
		TotalPage: res.TotalPage,

		PageNumber: res.PageNumber,
		PageSize:   res.PageSize,
	}
	return &pb.UserListResponse{
		Page:  pageResponse,
		Users: userResponses,
	}
}

func convertTaskListResponse(res *pagination.Result[model.Task]) *pb.TaskListResponse {
	var taskResponses []*pb.TaskResponse
	for _, task := range res.Results {
		taskResponse := convertTaskResponse(task)
		taskResponses = append(taskResponses, taskResponse)
	}
	pageResponse := &pb.PaginationResponse{
		TotalPage:  res.TotalPage,
		PageNumber: res.PageNumber,
		PageSize:   res.PageSize,
	}
	return &pb.TaskListResponse{
		Page:  pageResponse,
		Tasks: taskResponses,
	}
}
