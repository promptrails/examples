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

	// Create a chat session
	session, err := client.Chat.CreateSession(ctx, &promptrails.CreateSessionParams{
		AgentID: agentID,
		Title:   "Support conversation",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Session created: %s\n", session.ID)

	// Send messages
	msg1, err := client.Chat.SendMessage(ctx, session.ID, &promptrails.SendMessageParams{
		Content: "What is PromptRails?",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Assistant: %s\n", msg1.Content)

	msg2, err := client.Chat.SendMessage(ctx, session.ID, &promptrails.SendMessageParams{
		Content: "How do I create an agent?",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Assistant: %s\n", msg2.Content)

	// List message history
	messages, err := client.Chat.ListMessages(ctx, session.ID, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range messages.Data {
		content := msg.Content
		if len(content) > 80 {
			content = content[:80] + "..."
		}
		fmt.Printf("  [%s] %s\n", msg.Role, content)
	}

	// Clean up
	_ = client.Chat.DeleteSession(ctx, session.ID)
}
