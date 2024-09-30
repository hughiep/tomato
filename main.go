package main

import (
	"os"
	"tomato/cmd"
	"tomato/internal/server"

	"github.com/urfave/cli/v2"
)

func main() {
	appClis := &cli.App{
		Commands: []*cli.Command{
			cmd.MigrationCommand,
		},
	}

	appClis.Run(os.Args)
	server.Serve()
}
