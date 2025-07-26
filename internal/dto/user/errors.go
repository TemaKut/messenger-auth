package userdto

import "errors"

var (
	ErrUnknown                = errors.New("error unknown")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrInvalidCredentials     = errors.New("invalid credentials")
)
