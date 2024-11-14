package util

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vietquan-37/todo-list/internal/model"
)

var ErrInvalidToken = errors.New("invalid token")
var ErrExpiredToken = errors.New("expired token")

type JwtMaker struct {
	secret []byte
}

func NewService(secret string) (*JwtMaker, error) {
	if secret == "" {
		return nil, errors.New("cannot have an empty secret")
	}
	return &JwtMaker{secret: []byte(secret)}, nil
}
func (maker *JwtMaker) GenerateJWT(user *model.User, expiration time.Duration) (tokenString string, err error) {
	expirationTime := time.Now().Add(expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"iss":  time.Now().Unix(),
		"exp":  expirationTime.Unix(),
		"role": user.Role,
	})
	signed, err := token.SignedString(maker.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}
	return signed, nil
}

func (maker *JwtMaker) ValidateToken(_ context.Context, token string) (string, string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return maker.secret, nil
	})
	if err != nil {
		return "", "", errors.Join(ErrInvalidToken, err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		// Extract expiration time and check if token has expired
		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", "", fmt.Errorf("%w: missing or invalid expiration time", ErrInvalidToken)
		}
		if int64(exp) < time.Now().Unix() {
			return "", "", fmt.Errorf("%w: token is expired", ErrExpiredToken)
		}

		role, ok := claims["role"].(string)
		if !ok {
			return "", "", fmt.Errorf("%w: failed to extract role from claims", ErrInvalidToken)
		}

		var id string
		if sub, ok := claims["sub"].(string); ok {
			id = sub
		} else if sub, ok := claims["sub"].(float64); ok {

			id = fmt.Sprintf("%.0f", sub)
		} else {
			return "", "", fmt.Errorf("%w: failed to extract id from claims", ErrInvalidToken)
		}

		return id, role, nil
	}

	return "", "", ErrInvalidToken
}
