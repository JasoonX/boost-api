package auth

import "github.com/pkg/errors"

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("expired token")
	ErrInsufficientRole = errors.New("insufficient role")
	ErrUserNotFound     = errors.New("user not found")
)
