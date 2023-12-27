package main

import (
	"log"

	"github.com/Crystalix007/semantic-sensei/backend/cmd/semantic-sensei/serve"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "semantic-sensei",
		Short: "Semantic Sensei is a tool for building KNN classifiers for LLM output",
	}

	cmd.AddCommand(serve.Command())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
