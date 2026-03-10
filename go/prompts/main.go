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

	// Create a prompt
	prompt, err := client.Prompts.Create(ctx, &promptrails.CreatePromptParams{
		Name:        "Support Classifier",
		Description: "Classifies customer support tickets",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created prompt: %s\n", prompt.ID)

	// Create a version
	_, err = client.Prompts.CreateVersion(ctx, prompt.ID, &promptrails.CreateVersionParams{
		Message: "Initial classifier version",
		Config: map[string]any{
			"system_prompt": "You are a support ticket classifier.",
			"user_prompt":   "Classify this ticket: {{ message }}",
			"llm_model_id":  "gpt-4o",
			"temperature":   0.3,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Run the prompt
	response, err := client.Prompts.Run(ctx, prompt.ID, &promptrails.RunPromptParams{
		UserPrompt: "Classify: I want a refund for my order",
		LLMModelID: "gpt-4o",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", response.Content)
	fmt.Printf("Cost: $%.4f\n", response.Cost)

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
