package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

// Agent versions carry a typed config. The go-sdk defines one concrete
// type per agent kind (SimpleAgentConfig, ChainAgentConfig,
// MultiAgentConfig, WorkflowAgentConfig, CompositeAgentConfig). Each
// implements the AgentConfig interface and injects the "type"
// discriminator on marshal.

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// Agent configs reference prompts by ID — create a prompt first.
	prompt, err := client.Prompts.Create(ctx, &promptrails.CreatePromptParams{
		Name:         "Support reply",
		Description:  "Answers a customer question politely",
		SystemPrompt: "You are a helpful customer support assistant.",
		UserPrompt:   "Customer asks: {{ message }}",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create an agent and attach a typed SimpleAgentConfig version.
	agent, err := client.Agents.Create(ctx, &promptrails.CreateAgentParams{
		Name:        "Customer Support Bot",
		Type:        "simple",
		Description: "Handles customer inquiries",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created agent: %s\n", agent.ID)

	temperature := 0.7
	version, err := client.Agents.CreateVersion(ctx, agent.ID, &promptrails.CreateVersionParams{
		Version: "1.0.0",
		Message: "Initial version",
		Config: promptrails.SimpleAgentConfig{
			PromptID:    prompt.ID,
			Temperature: &temperature,
		},
		SetCurrent: true,
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
	fmt.Printf("Trace ID: %s\n", result.TraceID)

	// Chain two prompts with ChainAgentConfig. PromptLink identifies
	// each step by role + sort_order.
	extract, err := client.Prompts.Create(ctx, &promptrails.CreatePromptParams{
		Name:         "Extract",
		Description:  "Extracts named entities",
		SystemPrompt: "Extract named entities as a JSON list.",
		UserPrompt:   "{{ text }}",
	})
	if err != nil {
		log.Fatal(err)
	}

	chainAgent, err := client.Agents.Create(ctx, &promptrails.CreateAgentParams{
		Name:        "Extract-then-Summarize",
		Type:        "chain",
		Description: "Two-step reasoning chain",
	})
	if err != nil {
		log.Fatal(err)
	}

	if _, err := client.Agents.CreateVersion(ctx, chainAgent.ID, &promptrails.CreateVersionParams{
		Version: "1.0.0",
		Message: "Initial chain",
		Config: promptrails.ChainAgentConfig{
			PromptIDs: []promptrails.PromptLink{
				{PromptID: extract.ID, Role: "extract", SortOrder: 0},
				{PromptID: prompt.ID, Role: "answer", SortOrder: 1},
			},
		},
		SetCurrent: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Housekeeping
	agents, err := client.Agents.List(ctx, &promptrails.ListAgentsParams{Type: "simple"})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range agents.Data {
		fmt.Printf("  - %s (%s) — %s\n", a.Name, a.Type, a.Status)
	}

	_ = client.Agents.Delete(ctx, chainAgent.ID)
	_ = client.Agents.Delete(ctx, agent.ID)
	_ = client.Prompts.Delete(ctx, extract.ID)
	_ = client.Prompts.Delete(ctx, prompt.ID)
	fmt.Println("Cleanup complete.")
}
