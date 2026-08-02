package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

// In v2 a prompt version is content only: system/user text (+ optional input
// schema). Model, sampling and output schema live on the agent version, not
// the prompt. Prompts are no longer runnable directly — attach one to an agent
// and call Agents.Execute (see the agents example).

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// Create a prompt with its initial body
	prompt, err := client.Prompts.Create(ctx, &promptrails.CreatePromptParams{
		Name:         "Support Classifier",
		Description:  "Classifies customer support tickets",
		SystemPrompt: "You are a support ticket classifier.",
		UserPrompt:   "Classify this ticket: {{ message }}",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created prompt: %s\n", prompt.ID)

	// Add a content-only version — no model/sampling fields here.
	version, err := client.Prompts.CreateVersion(ctx, prompt.ID, &promptrails.CreatePromptVersionParams{
		Version:      "1.0.0",
		SystemPrompt: "You are a support ticket classifier.",
		UserPrompt:   "Classify this ticket: {{ message }}",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"message": map[string]any{"type": "string"}},
		},
		SetCurrent: true,
		Message:    "Initial classifier version",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created version: %s\n", version.Version)

	// List prompts
	prompts, err := client.Prompts.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range prompts.Data {
		fmt.Printf("  - %s (%s)\n", p.Name, p.ID)
	}

	// Clean up
	_ = client.Prompts.Delete(ctx, prompt.ID)
}
