package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/vietquan-37/todo-list/middleware"
	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/pkg/v1/val"
	"github.com/vietquan-37/todo-list/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (server *Server) GetAllUser(ctx context.Context, req *pb.Pagination) (*pb.UserListResponse, error) {
	userList, err := server.Repo.GetAllUser(req.Request.GetQuery(), req.Request.GetPageNumber(), req.Request.GetPageSize())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while retrieving users: %s", err)
	}
	return convertUserListResponse(userList), nil
}
func (server *Server) UpdateUser(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserResponse, error) {

	userID, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user ID not found in context")
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID format: %s", userID)
	}

	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "user id missing")
	}
	violation := validateUpdateUserRequest(req)
	if violation != nil {
		return nil, invalidArgumentError(violation)
	}
	if req.Password != nil {
		hashPasword, err := util.HashedPassword(*req.Password)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error while hashing password: %s", err)
		}
		req.Password = &hashPasword
	}
	model := convertUserUpdate(req)
	user, err := server.Repo.UpdateUser(userIDInt, model)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user with id %d not found", userIDInt)
		}
		return nil, status.Errorf(codes.Internal, "error while updating user: %s", err)
	}
	return convertUserResponse(*user), nil
}
func (server *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	userId := req.GetId()
	err := server.Repo.DeleteUser(int(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user with id %d not found", userId)
		}
		return nil, status.Errorf(codes.Internal, "error while deleting user: %s", err)
	}
	return &pb.DeleteUserResponse{
		Message: "Delete user successfully",
	}, nil

}
func validateUpdateUserRequest(req *pb.UserUpdateRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violation = append(violation, ErrorResponse("password", err))
		}
	}
	if req.PhoneNumber != nil {
		if err := val.ValidatePhoneNumber(req.GetPhoneNumber()); err != nil {
			violation = append(violation, ErrorResponse("phone_number", err))
		}
	}
	if req.FullName != nil {
		if err := val.ValidateFullname(req.GetFullName()); err != nil {
			violation = append(violation, ErrorResponse("full_name", err))
		}
	}

	return violation
}
