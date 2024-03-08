// Package URL allows for storing and retrieving request URL details in the
// strict server endpoints.
package url

import (
	"context"
	"net/http"
	"net/url"
)

type urlKey struct{}

// Absolute is a middleware function that resolves request URLs to absolute
// URLs using the Origin header.
func Absolute(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil {
			r.URL.Scheme = "http"
		} else {
			r.URL.Scheme = "https"
		}

		r.URL.Host = r.Host

		h.ServeHTTP(w, r)
	})
}

// Store is a middleware function that stores the incoming request URL in the
func Store(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, urlKey{}, r.URL)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get retrieves the URL from the given context.
func Get(ctx context.Context) *url.URL {
	url, ok := ctx.Value(urlKey{}).(*url.URL)
	if !ok {
		return nil
	}

	return url
}
