package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/schema/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Migrate applies database migrations using the provided migration source and
// driver.
// It returns an error if there was a problem creating the migration driver or
// source, or if there was an error running the migrations.
func (d Database) Migrate(ctx context.Context) error {
	driver, err := postgres.WithInstance(d.db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf(
			"storage: error creating migration driver: %w",
			err,
		)
	}

	source, err := iofs.New(migrations.Migrations, ".")
	if err != nil {
		return fmt.Errorf(
			"storage: error creating migration source: %w",
			err,
		)
	}

	migrator, err := migrate.NewWithInstance("migrations", source, "postgres", driver)
	if err != nil {
		return fmt.Errorf(
			"storage: error creating migrator: %w",
			err,
		)
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf(
			"storage: error running migrations: %w",
			err,
		)
	}

	return nil
}
