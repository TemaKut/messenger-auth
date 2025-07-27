package factory

import (
	"fmt"
	"github.com/TemaKut/messenger-auth/internal/app/config"
	usergrpchandler "github.com/TemaKut/messenger-auth/internal/app/handler/grpc/user"
	userservice "github.com/TemaKut/messenger-auth/internal/service/user"
	"github.com/google/wire"
	"time"
)

var ServiceSet = wire.NewSet(
	ProvideUserService,
	wire.Bind(new(usergrpchandler.Service), new(*userservice.Service)),
)

func ProvideUserService(cfg *config.Config, storage userservice.Storage) (*userservice.Service, error) {
	accessTokenLifetimeDuration, err := time.ParseDuration(cfg.Services.User.AccessTokenLifetimeDuration)
	if err != nil {
		return nil, fmt.Errorf("error parse access token lifetime duration. %w", err)
	}

	refreshTokenLifetimeDuration, err := time.ParseDuration(cfg.Services.User.RefreshTokenLifetimeDuration)
	if err != nil {
		return nil, fmt.Errorf("error parse refresh token lifetime duration. %w", err)
	}

	return userservice.NewService(
		storage,
		cfg.Services.User.AuthTokenSecret,
		accessTokenLifetimeDuration,
		refreshTokenLifetimeDuration,
	), nil
}
