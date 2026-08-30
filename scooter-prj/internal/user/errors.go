package user

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user with this data already exists")
	ErrInvalidToken      = errors.New("invalid or expired reset token")
)
