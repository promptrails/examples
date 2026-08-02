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
	agentID := os.Getenv("AGENT_ID")

	// Get agent card (A2A discovery)
	card, err := client.A2A.GetAgentCard(ctx, agentID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Agent: %s\n", card.Name)
	fmt.Printf("Description: %s\n", card.Description)

	// Send a message. The A2A message is a raw map with a role and parts.
	task, err := client.A2A.SendMessage(ctx, &promptrails.A2ASendMessageParams{
		AgentID: agentID,
		Message: map[string]any{
			"role": "user",
			"parts": []map[string]any{
				{"type": "text", "text": "Analyze the sales data for Q1 2026"},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task: %s\n", task.ID)
	fmt.Printf("Status: %s\n", task.Status)

	// Get task status
	updated, err := client.A2A.GetTask(ctx, task.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task status: %s\n", updated.Status)

	// List tasks
	tasks, err := client.A2A.ListTasks(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tasks.Data {
		fmt.Printf("  %s — %s\n", t.ID, t.Status)
	}

	// Cancel a running task
	_ = client.A2A.CancelTask(ctx, task.ID)
}
