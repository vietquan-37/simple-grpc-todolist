package handler

import (
	"context"
	"strconv"

	"github.com/vietquan-37/todo-list/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetFromCtx(ctx context.Context) (int, string, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return 0, "", status.Error(codes.Unauthenticated, "user ID not found in context")
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return 0, "", status.Errorf(codes.InvalidArgument, "invalid user ID format: %s", userID)
	}
	role, ok := ctx.Value(middleware.RoleKey).(string)
	if !ok {
		return 0, "", status.Error(codes.Unauthenticated, "role not found in context")
	}
	return userIDInt, role, nil
}
