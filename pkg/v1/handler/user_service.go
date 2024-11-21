package handler

import (
	"context"
	"errors"

	"github.com/bufbuild/protovalidate-go"
	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (server *Server) GetAllUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserListResponse, error) {
	userList, err := server.Repo.GetAllUser(req.GetQuery(), req.Request.GetPageNumber(), req.Request.GetPageSize())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while retrieving users: %s", err)
	}
	return convertUserListResponse(userList), nil
}
func (server *Server) UpdateUser(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}
	if err := validator.Validate(req); err != nil {

		violation := ErrorResponses(err)
		return nil, invalidArgumentError(violation)
	}
	userID, _, err := GetFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if req.Password != nil {
		hashPasword, err := util.HashedPassword(*req.Password)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error while hashing password: %s", err)
		}
		req.Password = &hashPasword
	}
	model := convertUserUpdate(req)
	user, err := server.Repo.UpdateUser(userID, model)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user with id %d not found", userID)
		}
		return nil, status.Errorf(codes.Internal, "error while updating user: %s", err)
	}
	return convertUserResponse(*user), nil
}
func (server *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.CommonResponse, error) {
	userId := req.GetId()
	err := server.Repo.DeleteUser(int(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user with id %d not found", userId)
		}
		return nil, status.Errorf(codes.Internal, "error while deleting user: %s", err)
	}
	return &pb.CommonResponse{
		Message: "Delete user successfully",
	}, nil

}
