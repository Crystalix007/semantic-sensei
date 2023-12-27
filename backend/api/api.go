package api

import (
	"net/http"

	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/go-chi/chi/v5"
)

// API represents an API with a handler.
type API struct {
	handler http.Handler
}

// Ensure [API] implements [openapi.ServerInterface].
var _ openapi.ServerInterface = (*API)(nil)

// Ensure [API] implements [http.Handler].
var _ http.Handler = (*API)(nil)

// New creates a new instance of the API.
// It returns a pointer to the API and an error, if any.
func New() (*API, error) {
	chi := chi.NewMux()

	a := &API{}
	a.handler = openapi.HandlerFromMuxWithBaseURL(a, chi, "/api")

	return a, nil
}

// ServeHTTP handles the HTTP requests for the API.
// It delegates the request to the underlying handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}
