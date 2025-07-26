package user

import (
	"context"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
)

type Service interface {
	Register(ctx context.Context, params userdto.RegisterParams) (userdto.User, error)
	Authorize(ctx context.Context, params userdto.UserAuthorizeParams) (userdto.UserAuthorizeResult, error)
}
