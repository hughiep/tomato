package main

import (
	"os"
	"tomato/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	appClis := &cli.App{
		Commands: []*cli.Command{
			cmd.ServeCommand,
			cmd.MigrationCommand,
		},
	}

	appClis.Run(os.Args)
}
