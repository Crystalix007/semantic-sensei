package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Crystalix007/semantic-sensei/backend/api/headers"
	"github.com/Crystalix007/semantic-sensei/backend/api/url"
	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// API represents an API with a handler.
type API struct {
	handler http.Handler
	db      *storage.Database

	loggingEnabled bool
}

// Ensure [API] implements [openapi.ServerInterface].
var _ openapi.StrictServerInterface = (*API)(nil)

// Ensure [API] implements [http.Handler].
var _ http.Handler = (*API)(nil)

// New creates a new instance of the API.
// It returns a pointer to the API and an error, if any.
func New(ctx context.Context, opts ...Option) (*API, error) {
	chi := chi.NewMux()

	// Resolve request URLs to their absolute representation.
	chi.Use(url.Absolute)

	// Store request headers in the context.
	chi.Use(headers.Store)

	// Store URL in the context.
	chi.Use(url.Store)

	a := &API{
		handler: chi,
	}

	for _, opt := range opts {
		opt(a)
	}

	if err := a.setDefaults(); err != nil {
		return nil, err
	}

	if a.loggingEnabled {
		chi.Use(middleware.Logger)
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

	chi.Mount("/api", openapi.Handler(strictServerHandler))

	return a, nil
}

// ServeHTTP handles the HTTP requests for the API.
// It delegates the request to the underlying handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "serving request", slog.String("method", r.Method), slog.String("path", r.URL.Path))

	a.handler.ServeHTTP(w, r)
}
