package migration

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

type Command struct {
	db *sql.DB
	fs embed.FS
}

func NewCommand(db *sql.DB, fs embed.FS) *Command {
	return &Command{db: db, fs: fs}
}

func (c *Command) Run(ctx context.Context, direction Direction, version int64) (err error) {
	provider, err := goose.NewProvider(goose.DialectPostgres, c.db, c.fs)
	if err != nil {
		return fmt.Errorf("error make goose provider. %w", err)
	}

	var result []*goose.MigrationResult

	switch direction {
	case DirectionUp:
		result, err = provider.Up(ctx)
		if err != nil {
			return fmt.Errorf("error up migration. %w", err)
		}
	case DirectionDown:
		res, err := provider.Down(ctx)
		if err != nil {
			return fmt.Errorf("error up migration. %w", err)
		}

		result = []*goose.MigrationResult{res}
	case DirectionDownTo:
		if version == 0 {
			return fmt.Errorf("error migration version is zero")
		}

		result, err = provider.DownTo(ctx, version)
		if err != nil {
			return fmt.Errorf("error up migration. %w", err)
		}
	default:
		return fmt.Errorf("error unknown direction %s", direction)
	}

	for _, res := range result {
		fmt.Println(res.String())
	}

	return nil
}
