package cmd

import (
	"tomato/internal/server"

	cli "github.com/urfave/cli/v2"
)

// Command to run the server
var ServeCommand = &cli.Command{
	Name:  "serve",
	Usage: "Run the server",
	Action: func(c *cli.Context) error {
		// Run the server
		server.Serve()
		return nil
	},
}
