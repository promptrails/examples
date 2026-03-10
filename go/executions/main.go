package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// List recent executions
	executions, err := client.Executions.List(ctx, &promptrails.ListExecutionsParams{
		ListParams: promptrails.ListParams{Limit: 10},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total executions: %d\n", executions.Meta.Total)

	for _, ex := range executions.Data {
		fmt.Printf("  %s — %s\n", ex.ID, ex.Status)
		fmt.Printf("    Agent: %s\n", ex.AgentID)
		fmt.Printf("    Duration: %dms, Cost: $%.4f\n", ex.DurationMs, ex.Cost)
	}

	// Filter by status
	failed, err := client.Executions.List(ctx, &promptrails.ListExecutionsParams{
		Status: "failed",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFailed executions: %d\n", failed.Meta.Total)
	for _, ex := range failed.Data {
		fmt.Printf("  %s: %s\n", ex.ID, ex.Error)
	}

	// Get execution details
	if len(executions.Data) > 0 {
		detail, err := client.Executions.Get(ctx, executions.Data[0].ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nExecution detail:\n")
		fmt.Printf("  Input: %v\n", detail.Input)
		fmt.Printf("  Output: %v\n", detail.Output)
	}
}
