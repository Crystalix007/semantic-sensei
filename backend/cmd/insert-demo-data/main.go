package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Crystalix007/semantic-sensei/backend/config"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	db, err := storage.Open(ctx, *cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Migrate(ctx); err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateProject(ctx, storage.Project{
		Name:        "Test Project",
		Description: "This is a test project",
	})
	if err != nil {
		log.Fatal(err)
	}
}
