package userservice

import (
	"errors"
	"fmt"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
	userstorage "github.com/TemaKut/messenger-auth/internal/storage/user"
)

func encodeUser(user *usermodels.User) userdto.User {
	return userdto.User{
		Id:       user.Id(),
		Name:     user.Name(),
		LastName: user.LastName(),
	}
}

func encodeError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, userstorage.ErrUserEmailAlreadyExists):
		return fmt.Errorf("%w, %w", userdto.ErrUserEmailAlreadyExists, err)
	default:
		return userdto.ErrUnknown
	}
}
