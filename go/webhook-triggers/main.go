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

	// Create a webhook trigger
	trigger, err := client.WebhookTriggers.Create(ctx, &promptrails.CreateWebhookTriggerParams{
		Name:      "GitHub Push Webhook",
		AgentID:   agentID,
		HasSecret: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trigger created: %s\n", trigger.ID)
	fmt.Printf("Webhook URL: %s\n", trigger.URL)
	fmt.Printf("Secret: %s\n", trigger.Secret)

	// List triggers
	triggers, err := client.WebhookTriggers.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range triggers.Data {
		fmt.Printf("  %s — active: %v\n", t.Name, t.IsActive)
	}

	// Delete trigger
	_ = client.WebhookTriggers.Delete(ctx, trigger.ID)
	fmt.Println("Trigger deleted.")
}
