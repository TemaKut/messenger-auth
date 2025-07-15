package factory

import (
	"fmt"
	"github.com/TemaKut/messenger-auth/internal/app/config"
	"github.com/TemaKut/messenger-auth/internal/app/handler/grpc/user"
	"github.com/TemaKut/messenger-auth/internal/app/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	ProvideApp,
	ProvideLogger,
	config.NewConfig,
	user.NewHandler,
)

type App struct{}

func ProvideApp(
	logger *logger.Logger,
	_ GrpcProvider,
) (App, func()) {
	logger.Infof("app inited")

	return App{}, func() {
		logger.Infof("app shutting down")
	}
}

func ProvideLogger(cfg *config.Config) (*logger.Logger, error) {
	var level logger.LogLevel

	switch cfg.Logger.Level {
	case config.LoggerLevelDebug:
		level = logger.LogLevelDebug
	case config.LoggerLevelInfo:
		level = logger.LogLevelInfo
	default:
		return nil, fmt.Errorf("error invalid log level: %d", cfg.Logger.Level)
	}

	return logger.NewLogger(level)
}
