package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

// The Go SDK's Prompts.Create accepts the prompt body inline —
// system_prompt and user_prompt live on CreatePromptParams, so you
// don't need a separate CreateVersion call to make a runnable prompt.
// Prompts.Run then executes it with live input.

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

	// Run the prompt — UserPrompt overrides the stored body for this
	// single invocation; Input supplies variables for Jinja templating.
	response, err := client.Prompts.Run(ctx, prompt.ID, &promptrails.RunPromptParams{
		UserPrompt: "Classify this ticket: {{ message }}",
		LLMModelID: "gpt-4o",
		Input:      map[string]any{"message": "I want a refund for my order"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s\n", response.Output)

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
