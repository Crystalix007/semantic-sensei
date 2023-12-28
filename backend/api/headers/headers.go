// Package headers provides utilities for storing and retrieving HTTP headers
// from a context.
package headers

import (
	"context"
	"net/http"
)

type headersKey struct{}

// Store is a middleware function that stores the incoming request in the
// context.
//
// This allows accessing the raw request headers, even in the strict server.
func Store(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, headersKey{}, r.Header.Clone())

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get retrieves the HTTP headers from the given context.
//
// Will return an empty header if the headers are not present in the context.
func Get(ctx context.Context) http.Header {
	header, ok := ctx.Value(headersKey{}).(http.Header)
	if !ok {
		return make(http.Header)
	}

	return header
}
