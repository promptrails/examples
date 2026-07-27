package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

// The dedicated Costs service is gone in v2. Cost, token and latency
// aggregates now come from Traces.GetSummary, which accepts a TraceFilterParams
// (date range, status, level, model name, agent/session/execution id).

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// Workspace-wide summary
	summary, err := client.Traces.GetSummary(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Workspace usage:\n")
	fmt.Printf("  Total cost:    $%.2f\n", summary.TotalCost)
	fmt.Printf("  Total tokens:  %d\n", summary.TotalTokens)
	fmt.Printf("  Total traces:  %d\n", summary.TotalTraces)
	fmt.Printf("  Errors:        %d\n", summary.ErrorCount)
	fmt.Printf("  Avg duration:  %.0fms\n", summary.AvgDurationMs)
	fmt.Printf("  Unique models: %d\n", summary.UniqueModels)

	// Filter by date range
	feb, err := client.Traces.GetSummary(ctx, &promptrails.TraceFilterParams{
		DateFrom: "2026-02-01",
		DateTo:   "2026-03-01",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFebruary cost: $%.2f over %d traces\n", feb.TotalCost, feb.TotalTraces)

	// Scope to a single agent
	agentID := os.Getenv("AGENT_ID")
	if agentID != "" {
		agentSummary, err := client.Traces.GetSummary(ctx, &promptrails.TraceFilterParams{
			AgentID: agentID,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Agent cost: $%.2f\n", agentSummary.TotalCost)
	}
}
