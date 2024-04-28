package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"

	"github.com/Crystalix007/semantic-sensei/backend/config"
)

// Register the default flavor for the SQL builder, so that the generated
// queries are valid.
func init() {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
}

// Database represents a connection to an SQL database.
type Database struct {
	db *sql.DB
}

// Open opens a connection to the database and returns a pointer to the
// Database struct. It uses the sqlite driver and connects to the
// "semantic-sensei.sqlite" file with foreign key support enabled.
// If the connection is successful, it returns a pointer to the Database struct
// and nil error.
// Otherwise, it returns nil and the error encountered during the connection.
func Open(ctx context.Context, cfg config.Config) (*Database, error) {
	db, err := sql.Open("postgres", cfg.Database.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf(
			"storage: error opening database connection: %w",
			err,
		)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf(
			"storage: error pinging database connection: %w",
			err,
		)
	}

	return &Database{db: db}, nil
}

// Close closes the database connection.
// It returns an error if there was a problem closing the connection.
func (d Database) Close() error {
	if err := d.db.Close(); err != nil {
		return fmt.Errorf(
			"storage: error closing database connection: %w",
			err,
		)
	}

	return nil
}
