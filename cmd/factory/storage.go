package factory

import (
	"database/sql"
	"fmt"
	"github.com/TemaKut/messenger-auth/internal/app/config"
	"github.com/TemaKut/messenger-auth/internal/app/logger"
	userservice "github.com/TemaKut/messenger-auth/internal/service/user"
	userstorage "github.com/TemaKut/messenger-auth/internal/storage/user"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

var StorageSet = wire.NewSet(
	ProvideUserStorage,
	ProvidePostgresDb,
	wire.Bind(new(userservice.Storage), new(*userstorage.Storage)),
)

type PostgresDb struct {
	*sql.DB
}

func ProvidePostgresDb(cfg *config.Config, logger *logger.Logger) (PostgresDb, func(), error) {
	logger.Infof("connect to postgres db")

	pgCfg, err := pgx.ParseConfig(cfg.Storage.Postgres.ConnStr)
	if err != nil {
		return PostgresDb{}, nil, fmt.Errorf("error parse config. %w", err)
	}

	db := stdlib.OpenDB(*pgCfg)

	if err := db.Ping(); err != nil {
		return PostgresDb{}, nil, fmt.Errorf("error ping postgres db. %w", err)
	}

	return PostgresDb{DB: db}, func() {
		logger.Infof("close postgres db")

		if err := db.Close(); err != nil {
			logger.Errorf("error close postgres db. %s", err)
		}
	}, nil
}

func ProvideUserStorage(
	db PostgresDb,
	logger *logger.Logger,
) *userstorage.Storage {
	return userstorage.NewStorage(db.DB, logger)
}
