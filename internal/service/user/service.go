package userservice

import (
	"context"
	"fmt"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
)

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Register(ctx context.Context, params userdto.RegisterParams) (userdto.User, error) {
	userModel := usermodels.NewUser(
		params.Name,
		params.LastName,
		params.Email,
		params.Password,
	)

	if err := s.storage.UserCreate(ctx, userModel); err != nil {
		return userdto.User{}, encodeError(fmt.Errorf("error register user. %w", err))
	}

	return encodeUser(userModel), nil
}

func (s *Service) Authorize(
	ctx context.Context,
	params userdto.UserAuthorizeParams,
) (userdto.UserAuthorizeResult, error) {
	authenticator, err := s.chooseUserCredentialsAuthenticator(params.Credentials)
	if err != nil {
		return userdto.UserAuthorizeResult{}, fmt.Errorf("error choose authentificator. %w", err)
	}

	fmt.Println(authenticator.authentify(ctx))

	return userdto.UserAuthorizeResult{}, nil
}

func (s *Service) chooseUserCredentialsAuthenticator(
	credentials userdto.UserAuthorizeCredentials,
) (userCredentialsAuthenticator, error) {
	switch {
	case credentials.Email != nil:
		return newUserCredentialsEmailAuthenticator(
			s.storage,
			credentials.Email.Email,
			credentials.Email.Password,
		), nil
	default:
		return nil, fmt.Errorf("error unsupported credentials")
	}
}
