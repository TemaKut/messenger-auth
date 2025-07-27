package userservice

import (
	"context"
	"fmt"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
	"time"
)

type Service struct {
	storage Storage

	authTokenSecret              string
	accessTokenLifetimeDuration  time.Duration
	refreshTokenLifetimeDuration time.Duration
}

func NewService(
	storage Storage,
	authTokenSecret string,
	accessTokenLifetimeDuration time.Duration,
	refreshTokenLifetimeDuration time.Duration,
) *Service {
	return &Service{
		storage:                      storage,
		authTokenSecret:              authTokenSecret,
		accessTokenLifetimeDuration:  accessTokenLifetimeDuration,
		refreshTokenLifetimeDuration: refreshTokenLifetimeDuration,
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

	user, err := authenticator.authentify(ctx)
	if err != nil {
		return userdto.UserAuthorizeResult{}, fmt.Errorf("error authentify user. %w", err)
	}

	accessTokenExpiredAt := time.Now().Add(s.accessTokenLifetimeDuration)
	refreshTokenExpiredAt := time.Now().Add(s.refreshTokenLifetimeDuration)

	accessTokenStr, err := newAuthToken().
		setSubject(user.Id()).
		setExpiredAt(accessTokenExpiredAt).
		setType(authTokenTypeAccess).
		build(s.authTokenSecret)
	if err != nil {
		return userdto.UserAuthorizeResult{}, fmt.Errorf("error building access token. %w", err)
	}

	refreshTokenStr, err := newAuthToken().
		setSubject(user.Id()).
		setExpiredAt(refreshTokenExpiredAt).
		setType(authTokenTypeRefresh).
		build(s.authTokenSecret)
	if err != nil {
		return userdto.UserAuthorizeResult{}, fmt.Errorf("error building refresh token. %w", err)
	}

	return userdto.UserAuthorizeResult{
		User: encodeUser(user),
		AuthParams: userdto.AuthParams{
			AccessToken:  encodeAuthToken(accessTokenStr, accessTokenExpiredAt),
			RefreshToken: encodeAuthToken(refreshTokenStr, refreshTokenExpiredAt),
		},
	}, nil
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
