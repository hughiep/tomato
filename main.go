package main

import (
	"os"
	"tomato/cmd"
	"tomato/internal/server"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			cmd.MigrationCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	server.App()
}
