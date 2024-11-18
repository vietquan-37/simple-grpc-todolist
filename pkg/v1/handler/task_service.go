package handler

import (
	"context"
	"errors"

	"github.com/bufbuild/protovalidate-go"
	"github.com/vietquan-37/todo-list/internal/enum"
	"github.com/vietquan-37/todo-list/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (server *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}
	if err := validator.Validate(req); err != nil {

		violation := ErrorResponses(err)
		return nil, invalidArgumentError(violation)
	}
	userId, _, err := GetFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	model := convertTask(req)
	model.UserID = uint(userId)
	task, err := server.TaskRepo.AddTask(model)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while creating task : %v", err)
	}
	return convertTaskResponse(*task), nil
}
func (server *Server) GetAllUserTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskListResponse, error) {
	userId, role, err := GetFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	if userId != int(req.UserId) && role != string(enum.Admin) {
		return nil, status.Errorf(codes.Unauthenticated, "not allow to view this task")
	}
	taskList, err := server.TaskRepo.GetUserTask(int(req.UserId), req.Request.GetPageNumber(), req.Request.GetPageSize())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while retrieving task: %v", err)
	}
	rsp := convertTaskListResponse(taskList)
	return rsp, nil
}
func (server *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	userId, _, err := GetFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	task, err := server.TaskRepo.GetTaskById(int(req.GetTaskId()))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no task with id %d was found", req.TaskId)
		}
		return nil, status.Errorf(codes.Internal, "error while retrieving  task: %v", err)
	}
	if userId != int(task.UserID) {
		return nil, status.Errorf(codes.Unauthenticated, "not allow to delete this task")
	}
	err = server.TaskRepo.DeleteTask(task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while deleting task: %v", err)
	}
	return &pb.DeleteTaskResponse{
		Message: "Delete task successfully",
	}, nil
}

func (server *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}
	if err = validator.Validate(req); err != nil {
		violation := ErrorResponses(err)
		invalidArgumentError(violation)
	}
	userId, _, err := GetFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	task, err := server.TaskRepo.GetTaskById(int(req.GetTaskId()))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found any task with id %d", req.TaskId)
		}
		return nil, status.Errorf(codes.Internal, "error while retrieving task: %v", err)
	}
	if userId != int(task.UserID) {
		return nil, status.Error(codes.Unauthenticated, "not allow to modify this task")
	}
	model := convertTaskUpdate(req)
	task, err = server.TaskRepo.UpdateTask(int(req.GetTaskId()), model)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while updating user: %v", err)
	}
	return convertTaskResponse(*task), nil
}
