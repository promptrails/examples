package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	promptrails "github.com/promptrails/go-sdk"
)

func main() {
	// Basic client setup
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))

	// List agents to verify connection
	ctx := context.Background()
	agents, err := client.Agents.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected! Found %d agents.\n", agents.Meta.Total)

	for _, agent := range agents.Data {
		fmt.Printf("  - %s (%s)\n", agent.Name, agent.Type)
	}

	// Custom configuration
	customClient := promptrails.NewClient(
		os.Getenv("PROMPTRAILS_API_KEY"),
		promptrails.WithBaseURL("https://api.promptrails.ai"),
		promptrails.WithTimeout(60*time.Second),
		promptrails.WithMaxRetries(5),
	)
	_ = customClient
}
