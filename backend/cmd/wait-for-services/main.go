package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Crystalix007/semantic-sensei/backend/config"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// LivenessCheck represents a function that performs a liveness check.
// It takes a context.Context as input and returns an error if the liveness
// check fails.
type LivenessCheck func(ctx context.Context) error

// Services represents a map of service names to their corresponding liveness
// checks.
type Services map[string]LivenessCheck

// ErrNonTransientServiceFailure represents a non-transient error.
// It is used to indicate an error that is not expected to be resolved by
// continuing to wait for the service to become healthy.
var ErrNonTransientServiceFailure = errors.New("non-transient service failure")

// retryInterval represents the interval at which to retry the liveness check.
const retryInterval = 3 * time.Second

// Flags represents the command line flags for the application.
type Flags struct {
	// Database flag indicates whether to wait for the database service.
	Database bool
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var flags Flags

	command := cobra.Command{
		Use:  "wait-for-services",
		RunE: flags.waitForServices,
	}

	command.Flags().BoolVar(&flags.Database, "database", false, "Wait for the database to be ready")

	var logLevel slog.LevelVar
	logLevel.Set(slog.LevelDebug)

	if err := command.ExecuteContext(ctx); err != nil {
		slog.ErrorContext(ctx, "error waiting for services", slog.Any("error", err))

		os.Exit(1)
	}
}

func (f Flags) waitForServices(cmd *cobra.Command, _ []string) error {
	services, err := f.getServiceHealthchecks()
	if err != nil {
		return err
	}

	errg, ctx := errgroup.WithContext(cmd.Context())

	for name, check := range services {
		name, check := name, check
		errg.Go(func() error {
			err := waitForService(ctx, check)

			if err != nil {
				return fmt.Errorf(
					"error waiting for service '%s': %w",
					name,
					err,
				)
			}

			return err
		})
	}

	return errg.Wait()
}

func (f Flags) getServiceHealthchecks() (Services, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("wait-for-services: error reading config: %w", err)
	}

	services := make(Services)

	if f.Database {
		services["database"] = func(ctx context.Context) error {
			_, err := storage.Open(ctx, *cfg)

			var netError *net.OpError

			if errors.As(err, &netError) && netError.Op == "dial" && netError.Err == syscall.EHOSTUNREACH {
				return fmt.Errorf("database: host unreachable: %w: %w", ErrNonTransientServiceFailure, err)
			}

			return err
		}
	}

	return services, nil
}

func waitForService(ctx context.Context, check LivenessCheck) error {
	var lastErr error

	for {
		select {
		case <-ctx.Done():
			return lastErr
		case <-time.After(retryInterval):
			err := check(ctx)

			if err == nil {
				return nil
			}

			if errors.Is(err, ErrNonTransientServiceFailure) {
				slog.ErrorContext(ctx, "error checking service", slog.Any("error", err))

				return err
			}

			slog.InfoContext(ctx, "service not ready", slog.Any("error", err))
			lastErr = err
		}
	}
}
