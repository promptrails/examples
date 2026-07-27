package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

// Approvals are execution-scoped in v2. A run that hits an approval-gated tool
// or sub-agent parks at status "waiting_approval"; resume it with
// Executions.Approve / Executions.Deny. Executions.Tree returns the full
// parent -> child tree; Executions.Cancel requests cooperative cancellation.

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// List recent executions
	executions, err := client.Executions.List(ctx, &promptrails.ListExecutionsParams{
		Limit: 10,
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
		fmt.Printf("  %s: %v\n", ex.ID, ex.Error)
	}

	// Get execution details, then walk its full child tree
	if len(executions.Data) > 0 {
		detail, err := client.Executions.Get(ctx, executions.Data[0].ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nExecution detail:\n")
		fmt.Printf("  Input: %v\n", detail.Input)
		fmt.Printf("  Output: %v\n", detail.Output)

		tree, err := client.Executions.Tree(ctx, detail.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  Child executions: %d\n", len(tree.Children))
	}

	// Approval inbox — runs parked at waiting_approval
	inbox, err := client.Executions.ApprovalInbox(ctx, &promptrails.ListParams{Limit: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAwaiting approval: %d\n", inbox.Meta.Total)
	for _, ex := range inbox.Data {
		fmt.Printf("  %s — %s\n", ex.ID, ex.Status)
		// Approve to resume, or deny to resume with a denial:
		// client.Executions.Approve(ctx, ex.ID, &promptrails.DecideParams{Reason: "Looks safe"})
		// client.Executions.Deny(ctx, ex.ID, &promptrails.DecideParams{Reason: "Not allowed"})
	}

	// Cancel a still-running execution:
	// client.Executions.Cancel(ctx, "exec-id")
}
