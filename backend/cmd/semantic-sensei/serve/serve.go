package serve

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Crystalix007/semantic-sensei/backend/api"
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

	return cmd
}

func serve(cmd *cobra.Command, _ []string) error {
	api, err := api.New(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to create API: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", api)

	if dev, err := cmd.Flags().GetString("dev"); err == nil && dev != "" {
		devURL, err := url.Parse(dev)
		if err != nil {
			return fmt.Errorf("failed to parse dev URL: %w", err)
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(devURL)

		mux.Handle("/", reverseProxy)
	}

	return http.ListenAndServe(":8080", mux)
}
