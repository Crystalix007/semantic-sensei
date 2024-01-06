package api

import (
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

// Option represents a function that can be used to configure an API instance.
type Option = func(a *API)

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
func (a *API) setDefaults() error {
	var err error

	if a.db == nil {
		a.db, err = storage.Open()
		if err != nil {
			return fmt.Errorf(
				"api: error opening database: %w",
				err,
			)
		}
	}

	return nil
}
