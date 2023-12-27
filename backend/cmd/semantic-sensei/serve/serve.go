package serve

import (
	"fmt"
	"net/http"

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

	return cmd
}

func serve(cmd *cobra.Command, _ []string) error {
	api, err := api.New(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to create API: %w", err)
	}

	return http.ListenAndServe(":8080", api)
}
