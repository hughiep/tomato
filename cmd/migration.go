package cmd

import (
	"fmt"
	"net/url"
	"strings"
	"tomato/internal/config"
	"tomato/internal/db/migrations"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	cli "github.com/urfave/cli/v2" // imports as package "cli"
)

var MigrationCommand = &cli.Command{
	Name:  "migrate",
	Usage: "Run the database migrations",
	Action: func(c *cli.Context) error {
		// Run the database migrations
		migrations.Up(getMigrationDatabaseURL())
		return nil
	},
	Subcommands: []*cli.Command{
		{
			Name:  "down",
			Usage: "Rollback the last database migration",
			Action: func(c *cli.Context) error {
				// Rollback the last database migration
				migrations.Down(getMigrationDatabaseURL())
				return nil

			},
		},
		{
			Name:  "up",
			Usage: "Run the database migrations",
			Action: func(c *cli.Context) error {
				// Run the database migrations
				migrations.Up(getMigrationDatabaseURL())
				return nil
			},
		},
	},
}

func getMigrationDatabaseURL() string {
	configs := config.Load()
	host := "127.0.0.1"
	port := configs.Database.MysqlDbPort
	database := configs.Database.MysqlDbName
	username := "root"
	password := configs.Database.MySqlDbRootPassword

	return fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s",
		strings.TrimSpace(username),
		strings.TrimSpace(password),
		strings.TrimSpace(host),
		strings.TrimSpace(port),
		strings.TrimSpace(database),
	) + fmt.Sprintf("?%s", url.Values{
		"charset": []string{"utf8mb4"},
		"loc":     []string{"UTC"},
	}.Encode())
}
