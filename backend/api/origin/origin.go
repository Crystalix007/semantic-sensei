// Package origin provides utilities for determining the origin of a request.
package origin

import (
	"context"

	"github.com/Crystalix007/semantic-sensei/backend/api/headers"
)

// Origin represents the origin of a request or response.
type Origin string

const (
	// API represents a request who's origin is the API.
	API Origin = "API"
	// Frontend represents a request who's origin is the frontend.
	Frontend Origin = "Frontend"
)

// Get returns the origin based on the provided HTTP request.
// If the "Origin" header is not empty, it returns the Frontend origin.
// Otherwise, it returns the API origin.
func Get(ctx context.Context) Origin {
	headers := headers.Get(ctx)

	if headers != nil && headers.Get("Origin") != "" {
		return Frontend
	}

	return API
}
