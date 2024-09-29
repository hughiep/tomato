package migrations

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func migration(dsn string) (*migrate.Migrate, error) {
	fmt.Println("Migrating...")
	files, err := iofs.New(sqlFiles, "sql")
	if err != nil {
		fmt.Println("error" + err.Error())
		panic(err)
	}

	return migrate.NewWithSourceInstance(
		"iofs",
		files,
		dsn)
}

func Up(dsn string) error {

	m, err := migration(dsn)
	if err != nil {
		fmt.Println("error" + err.Error())
		return err
	}
	fmt.Println("Migrating up...", m)

	return m.Up()
}

func Down(dsn string) error {
	m, err := migration(dsn)
	if err != nil {
		return err
	}

	return m.Down()
}
