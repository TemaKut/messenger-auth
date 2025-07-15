package userservice

import (
	"context"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
	userstorage "github.com/TemaKut/messenger-auth/internal/storage/user"
)

type Storage interface {
	UserCreate(ctx context.Context, params userstorage.UserCreateParams) (*usermodels.User, error)
}
