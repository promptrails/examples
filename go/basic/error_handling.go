package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	_, err := client.Agents.Get(ctx, "nonexistent-id")
	if err != nil {
		var notFound *promptrails.NotFoundError
		var validationErr *promptrails.ValidationError
		var rateLimitErr *promptrails.RateLimitError
		var unauthorizedErr *promptrails.UnauthorizedError

		switch {
		case errors.As(err, &notFound):
			fmt.Printf("Agent not found: %s\n", notFound.Message)
		case errors.As(err, &validationErr):
			fmt.Printf("Invalid request: %s\n", validationErr.Message)
		case errors.As(err, &rateLimitErr):
			fmt.Println("Rate limited — retry after a delay")
		case errors.As(err, &unauthorizedErr):
			fmt.Println("Invalid API key")
		default:
			log.Fatalf("API error: %v", err)
		}
	}
}
