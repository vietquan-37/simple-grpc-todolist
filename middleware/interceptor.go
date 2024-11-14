package middleware

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	UserIDKey = contextKey("user_id")
	roleKey   = contextKey("role")
)

// Validator interface to validate JWT token
type Validator interface {
	ValidateToken(ctx context.Context, token string) (string, string, error) // returns userID and role
}

// AuthInterceptor struct to hold the validator
type authInterceptor struct {
	validator Validator
}

var publicMethods = map[string]struct{}{
	"/pb.TodoList/Login":    {},
	"/pb.TodoList/Register": {},
}

var methodRoles = map[string][]string{
	"/pb.TodoList/DeleteUser": {"ADMIN"},
	"/pb.TodoList/UpdateUser": {"USER", "ADMIN"},
	"/pb.TodoList/GetAllUser": {"ADMIN"},
	"/pb.TodoList/CreateTask": {"USER"},
}

func NewAuthInterceptor(validator Validator) (*authInterceptor, error) {
	if validator == nil {
		return nil, errors.New("validator cannot be nil")
	}
	return &authInterceptor{validator: validator}, nil
}

// Check if the method is public (does not require authentication)
func (i *authInterceptor) isPublicMethod(method string) bool {
	_, exists := publicMethods[method]
	return exists
}

// Check if the method is allowed for the user's role
func (i *authInterceptor) isAuthorized(method string, role string) bool {
	allowedRoles, exists := methodRoles[method]
	if !exists {
		return false
	}

	// Check if the user's role is in the allowed roles for this method
	for _, r := range allowedRoles {
		if r == role {
			return true
		}
	}
	return false
}

// The middleware for checking authentication and RBAC
func (i *authInterceptor) UnaryAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Skip authentication for public methods like login and register
	if i.isPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	// Get metadata object from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	// Extract token from the authorization header
	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	if !strings.HasPrefix(token[0], "Bearer ") {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not of type Bearer")
	}

	// Validate the token
	actualToken := strings.TrimPrefix(token[0], "Bearer ")
	userID, role, err := i.validator.ValidateToken(ctx, actualToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	ctx = context.WithValue(ctx, UserIDKey, userID)

	if !i.isAuthorized(info.FullMethod, role) {
		return nil, status.Error(codes.PermissionDenied, "access denied for the user role")
	}

	return handler(ctx, req)
}
