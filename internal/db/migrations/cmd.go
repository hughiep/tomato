package migrations

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func migration(dsn string) (*migrate.Migrate, error) {
	files, err := iofs.New(sqlFiles, "sql")
	if err != nil {
		zap.L().Error("Failed to read sql files", zap.Error(err))
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
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Migration up successful")
	return nil
}

func Down(dsn string) error {
	m, err := migration(dsn)
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Migration down successful")
	return nil
}
