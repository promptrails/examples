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

	// 1. Generic webhook trigger (default source)
	trigger, err := client.AgentTriggers.Create(ctx, &promptrails.CreateAgentTriggerParams{
		Name:           "GitHub Push Webhook",
		AgentID:        agentID,
		GenerateSecret: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generic trigger created: %s\n", trigger.ID)
	fmt.Printf("  Token:  %s\n", trigger.Token)
	fmt.Printf("  Secret: %s\n", trigger.Secret)

	// 2. Slack trigger — point at workspace credentials for signing secret + bot token
	slackTrigger, err := client.AgentTriggers.Create(ctx, &promptrails.CreateAgentTriggerParams{
		Name:    "Slack #incidents",
		AgentID: agentID,
		Source:  promptrails.AgentTriggerSourceSlack,
		SourceConfig: map[string]interface{}{
			"signing_secret_credential_id": "<credential id with the Slack signing secret>",
			"bot_token_credential_id":      "<credential id with the bot token>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Slack trigger created: %s\n", slackTrigger.ID)

	// 3. Schedule trigger — fires on a cron expression
	scheduleTrigger, err := client.AgentTriggers.Create(ctx, &promptrails.CreateAgentTriggerParams{
		Name:    "Daily digest at 9am",
		AgentID: agentID,
		Source:  promptrails.AgentTriggerSourceSchedule,
		SourceConfig: map[string]interface{}{
			"cron_expression": "0 9 * * *",
			"next_run_at":     "2026-05-22T09:00:00Z",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Schedule trigger created: %s\n", scheduleTrigger.ID)

	// List + delete still work the same way
	triggers, err := client.AgentTriggers.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range triggers.Data {
		fmt.Printf("  %s (%s) — active: %v\n", t.Name, t.Source, t.IsActive)
	}

	_ = client.AgentTriggers.Delete(ctx, trigger.ID)
	_ = client.AgentTriggers.Delete(ctx, slackTrigger.ID)
	_ = client.AgentTriggers.Delete(ctx, scheduleTrigger.ID)
	fmt.Println("Triggers cleaned up.")
}
