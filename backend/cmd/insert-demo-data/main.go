package main

import (
	"context"
	"log"

	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

func main() {
	db, err := storage.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ctx := context.Background()

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
