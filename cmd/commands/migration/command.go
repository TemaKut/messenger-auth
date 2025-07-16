package migration

import (
	"fmt"
	"github.com/TemaKut/messenger-auth/cmd/commands/migration/factory"
	"github.com/TemaKut/messenger-auth/internal/app/commands/migration"
	"github.com/urfave/cli/v2"
)

var Command *cli.Command

func init() {
	Command = &cli.Command{
		Name:  "migration",
		Usage: "Manage migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "direction",
				Usage: fmt.Sprintf(
					"Direction to migrate. (%s, %s, %s)",
					migration.DirectionUp,
					migration.DirectionDown,
					migration.DirectionDownTo,
				),
				Required: true,
			},
			&cli.Int64Flag{
				Name:     "version",
				Required: false,
			},
		},
		Action: func(cliCtx *cli.Context) error {
			cmd, cleanup, err := factory.InitCommand()
			if err != nil {
				if cleanup != nil {
					cleanup()
				}

				return fmt.Errorf("error init command. %w", err)
			}

			defer cleanup()

			err = cmd.Run(cliCtx.Context, cliCtx.String("direction"), cliCtx.Int64("version"))
			if err != nil {
				return fmt.Errorf("error run command. %w", err)
			}

			return nil
		},
	}
}
