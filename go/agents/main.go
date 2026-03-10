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

	// Create an agent
	agent, err := client.Agents.Create(ctx, &promptrails.CreateAgentParams{
		Name:        "Customer Support Bot",
		Type:        "simple",
		Description: "Handles customer inquiries",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created agent: %s\n", agent.ID)

	// Create a version
	version, err := client.Agents.CreateVersion(ctx, agent.ID, &promptrails.CreateVersionParams{
		Message: "Initial version",
		Config: map[string]any{
			"llm_model_id":  "gpt-4o",
			"system_prompt": "You are a helpful customer support assistant.",
			"temperature":   0.7,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created version: %s\n", version.ID)

	// Execute the agent
	result, err := client.Agents.Execute(ctx, agent.ID, &promptrails.ExecuteAgentParams{
		Input: map[string]any{"message": "What is your refund policy?"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %s\n", result.Status)
	fmt.Printf("Output: %v\n", result.Output)
	fmt.Printf("Cost: $%.4f\n", result.Cost)
	fmt.Printf("Duration: %dms\n", result.DurationMs)

	// List agents with filtering
	agents, err := client.Agents.List(ctx, &promptrails.ListAgentsParams{
		Type: "simple",
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range agents.Data {
		fmt.Printf("  - %s (%s) — %s\n", a.Name, a.Type, a.Status)
	}

	// Delete agent
	if err := client.Agents.Delete(ctx, agent.ID); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Agent deleted.")
}
