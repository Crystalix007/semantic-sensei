package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
	"github.com/go-chi/chi/v5"
)

// API represents an API with a handler.
type API struct {
	handler http.Handler
	db      *storage.Database
}

// Ensure [API] implements [openapi.ServerInterface].
var _ openapi.StrictServerInterface = (*API)(nil)

// Ensure [API] implements [http.Handler].
var _ http.Handler = (*API)(nil)

// New creates a new instance of the API.
// It returns a pointer to the API and an error, if any.
func New(ctx context.Context, opts ...Option) (*API, error) {
	chi := chi.NewMux()

	a := &API{
		handler: chi,
	}

	for _, opt := range opts {
		opt(a)
	}

	if err := a.setDefaults(); err != nil {
		return nil, err
	}

	// TODO: the DB needs to be more persistent here, but for now, we'll just
	// force the schema to be reloaded on every startup.
	if err := a.db.Migrate(ctx); err != nil {
		return nil, fmt.Errorf(
			"api: error migrating database: %w",
			err,
		)
	}

	strictServerHandler := openapi.NewStrictHandler(a, nil)
	openapi.HandlerFromMuxWithBaseURL(strictServerHandler, chi, "/api")

	return a, nil
}

// ServeHTTP handles the HTTP requests for the API.
// It delegates the request to the underlying handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}
