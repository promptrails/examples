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

	// Get workspace cost summary
	summary, err := client.Costs.GetSummary(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Workspace costs:\n")
	fmt.Printf("  Total cost: $%.2f\n", summary.TotalCost)
	fmt.Printf("  Total tokens: %d\n", summary.TotalTokens)
	fmt.Printf("  Total executions: %d\n", summary.TotalExecutions)

	// Filter by date range
	feb, err := client.Costs.GetSummary(ctx, &promptrails.CostParams{
		From: "2026-02-01",
		To:   "2026-03-01",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFebruary costs: $%.2f\n", feb.TotalCost)

	// Get costs for a specific agent
	agentID := os.Getenv("AGENT_ID")
	if agentID != "" {
		agentCosts, err := client.Costs.GetAgentSummary(ctx, agentID, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Agent costs: $%.2f\n", agentCosts.TotalCost)
	}
}
