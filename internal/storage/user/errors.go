package userstorage

import "errors"

var (
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrUnknown                = errors.New("unknown error")
)
