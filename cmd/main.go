package main

import (
	"context"
	"fmt"
	"github.com/TemaKut/messenger-auth/cmd/factory"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := cli.App{
		Name: "Auth",
		Action: func(cliCtx *cli.Context) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer cancel()

			_, shutdown, err := factory.InitApp()
			if err != nil {
				if shutdown != nil {
					shutdown()
				}

				return fmt.Errorf("error init app. %w", err)
			}

			<-ctx.Done()

			shutdown()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("fatal run app. %s", err)
	}
}
