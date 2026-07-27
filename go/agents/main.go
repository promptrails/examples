package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

// PromptRails v2 has two agent kinds: "agent" (a prompt + optional tools /
// sub-agents) and "workflow" (a deterministic DAG). The go-sdk models each as
// a concrete AgentConfig — PromptAgentConfig or WorkflowAgentConfig — and
// injects the "type" discriminator on marshal. Model + sampling (ModelConfig),
// the run budget, approval policy, tools and sub-agents are version-scoped
// siblings of Config, not part of it. Prompts are content-only.

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// An agent references a prompt by id — create the content-only prompt
	// first. Model/sampling live on the agent version, not the prompt.
	prompt, err := client.Prompts.Create(ctx, &promptrails.CreatePromptParams{
		Name:         "Support reply",
		Description:  "Answers a customer question politely",
		SystemPrompt: "You are a helpful customer support assistant.",
		UserPrompt:   "Customer asks: {{ message }}",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create an "agent" and attach a version. ModelConfig + RunBudget ride on
	// the version alongside Config.
	agent, err := client.Agents.Create(ctx, &promptrails.CreateAgentParams{
		Name:        "Customer Support Bot",
		Type:        "agent",
		Description: "Handles customer inquiries",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created agent: %s\n", agent.ID)

	temperature := 0.7
	maxCost := 0.50
	version, err := client.Agents.CreateVersion(ctx, agent.ID, &promptrails.CreateVersionParams{
		Version: "1.0.0",
		Message: "Initial version",
		Config:  promptrails.PromptAgentConfig{PromptID: prompt.ID},
		ModelConfig: &promptrails.ModelConfig{
			ModelID:     "gpt-4o",
			Temperature: &temperature,
		},
		RunBudget:  &promptrails.RunBudget{MaxCost: &maxCost},
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

	// A "workflow" agent runs a deterministic DAG. Each WorkflowNode pins a
	// prompt and declares its dependencies.
	extract, err := client.Prompts.Create(ctx, &promptrails.CreatePromptParams{
		Name:         "Extract",
		Description:  "Extracts named entities",
		SystemPrompt: "Extract named entities as a JSON list.",
		UserPrompt:   "{{ text }}",
	})
	if err != nil {
		log.Fatal(err)
	}

	workflowAgent, err := client.Agents.Create(ctx, &promptrails.CreateAgentParams{
		Name:        "Extract-then-Answer",
		Type:        "workflow",
		Description: "Two-step workflow",
	})
	if err != nil {
		log.Fatal(err)
	}

	if _, err := client.Agents.CreateVersion(ctx, workflowAgent.ID, &promptrails.CreateVersionParams{
		Version: "1.0.0",
		Message: "Initial workflow",
		Config: promptrails.WorkflowAgentConfig{
			Nodes: []promptrails.WorkflowNode{
				{ID: "extract", PromptID: extract.ID, DependsOn: []string{}},
				{ID: "answer", PromptID: prompt.ID, DependsOn: []string{"extract"}},
			},
		},
		ModelConfig: &promptrails.ModelConfig{ModelID: "gpt-4o-mini"},
		SetCurrent:  true,
	}); err != nil {
		log.Fatal(err)
	}

	// Housekeeping
	agents, err := client.Agents.List(ctx, &promptrails.ListAgentsParams{Type: "agent"})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range agents.Data {
		fmt.Printf("  - %s (%s) — %s\n", a.Name, a.Type, a.Status)
	}

	_ = client.Agents.Delete(ctx, workflowAgent.ID)
	_ = client.Agents.Delete(ctx, agent.ID)
	_ = client.Prompts.Delete(ctx, extract.ID)
	_ = client.Prompts.Delete(ctx, prompt.ID)
	fmt.Println("Cleanup complete.")
}
