package api

import (
	"context"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/config"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

// Option represents a function that can be used to configure an API instance.
type Option = func(a *API)

// WithConfig sets the configuration on the API.
func WithConfig(cfg *config.Config) Option {
	return func(a *API) {
		a.config = cfg
	}
}

// WithDatabase sets the database on the API.
func WithDatabase(db *storage.Database) Option {
	return func(a *API) {
		a.db = db
	}
}

// WithLogging sets the logging enabled flag.
func WithLogging(logging bool) Option {
	return func(a *API) {
		a.loggingEnabled = logging
	}
}

// setDefaults sets the default values for the API instance.
func (a *API) setDefaults(ctx context.Context) error {
	var err error

	if a.config == nil {
		a.config, err = config.New()
		if err != nil {
			return fmt.Errorf(
				"api: error reading configuration: %w",
				err,
			)
		}
	}

	if a.db == nil {
		a.db, err = storage.Open(ctx, *a.config)
		if err != nil {
			return fmt.Errorf(
				"api: error opening database: %w",
				err,
			)
		}
	}

	return nil
}
