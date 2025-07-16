package factory

import (
	mainfactory "github.com/TemaKut/messenger-auth/cmd/factory"
	"github.com/TemaKut/messenger-auth/internal/app/commands/migration"
	"github.com/TemaKut/messenger-auth/internal/app/config"
	"github.com/TemaKut/messenger-auth/internal/app/migrations/postgres"
	"github.com/google/wire"
)

var CommandSet = wire.NewSet(
	config.NewConfig,
	mainfactory.ProvideLogger,
	mainfactory.ProvidePostgresDb,
	ProvideCommand,
)

func ProvideCommand(db mainfactory.PostgresDb) *migration.Command {
	return migration.NewCommand(db.DB, postgres.Embed)
}
