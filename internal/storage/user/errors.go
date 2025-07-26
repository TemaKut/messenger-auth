package userstorage

import "errors"

var (
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrUnknown                = errors.New("unknown error")
)
