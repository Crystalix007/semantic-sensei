// Package redirect provides utilities for redirecting requests depending upon
// the origin.
package redirect

import (
	"context"

	"github.com/Crystalix007/semantic-sensei/backend/api/origin"
	"github.com/Crystalix007/semantic-sensei/backend/openapi"
)

// Should determines whether the request should be redirected based on the
// origin.
func Should(ctx context.Context) bool {
	return origin.Get(ctx) == origin.Frontend
}

// To redirects the user to the specified location.
// It returns a RedirectResponse containing the redirect headers.
func To(location string) (openapi.RedirectResponse, error) {
	return openapi.RedirectResponse{
		Headers: openapi.RedirectResponseHeaders{
			Location: location,
		},
	}, nil
}
