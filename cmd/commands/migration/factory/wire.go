//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/TemaKut/messenger-auth/internal/app/commands/migration"
	"github.com/google/wire"
)

func InitCommand() (*migration.Command, func(), error) {
	panic(wire.Build(CommandSet))
}
