package userservice

import (
	"context"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
)

type Storage interface {
	UserCreate(ctx context.Context, params *usermodels.User) error
	UserByEmail(ctx context.Context, email string) (*usermodels.User, error)
}

type userCredentialsAuthenticator interface {
	authentify(ctx context.Context) (*usermodels.User, error)
}
