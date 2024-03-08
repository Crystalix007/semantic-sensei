package serve

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/Crystalix007/semantic-sensei/backend/api"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

// Command returns a new instance of the cobra.Command for serving the
// semantic-sensei API server.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the semantic-sensei API server",
		RunE:  serve,
	}

	cmd.Flags().String("dev", "", "enable development mode, proxying non-API requests to the specified URL")
	cmd.Flags().String("address", "", "sets the address that the API is hosted on")
	cmd.Flags().Bool("debug", false, "enables debug logging of requests")

	return cmd
}

func serve(cmd *cobra.Command, _ []string) error {
	apiOpts := []api.Option{}

	if debug, err := cmd.Flags().GetBool("debug"); err == nil && debug {
		apiOpts = append(apiOpts, api.WithLogging(true))
	}

	api, err := api.New(cmd.Context(), apiOpts...)
	if err != nil {
		return fmt.Errorf("failed to create API: %w", err)
	}

	mux := chi.NewRouter()

	mux.Handle("/api/*", api)

	if dev, err := cmd.Flags().GetString("dev"); err == nil && dev != "" {
		slog.SetDefault(getDevLogger())

		devURL, err := url.Parse(dev)
		if err != nil {
			return fmt.Errorf("failed to parse dev URL: %w", err)
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(devURL)

		mux.Handle("/*", reverseProxy)

		// Serve OpenAPI documentation.
		mux.Handle("/api", reverseProxy)
		mux.Get("/api/", http.RedirectHandler("/api", http.StatusMovedPermanently).ServeHTTP)
	}

	address, err := cmd.Flags().GetString("address")
	if err != nil {
		address = ":0"
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}

	defer listener.Close()

	log.Printf("listening on http://%s", listener.Addr())

	return http.Serve(listener, mux)
}

func getDevLogger() *slog.Logger {
	programLevel := new(slog.LevelVar)

	programLevel.Set(slog.LevelDebug)

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})

	return slog.New(handler)
}
