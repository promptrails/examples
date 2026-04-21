// Server-Sent Events streaming — live chat turns and executions.
//
// Chat.SendMessageStream and Executions.Stream both return a
// *ChatStream. The stream iterates typed events
// (ExecutionEvent, ThinkingEvent, ToolStartEvent, ToolEndEvent,
// ContentEvent, DoneEvent, ErrorEvent) — dispatch with a type switch.
// Always defer stream.Close(); cancel mid-stream by cancelling ctx.

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
	if agentID == "" {
		agentID = "your-agent-id"
	}

	// ---- Stream a chat turn -----------------------------------------

	session, err := client.Chat.CreateSession(ctx, &promptrails.CreateSessionParams{
		AgentID: agentID,
		Title:   "Streaming demo",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Session: %s\n\n", session.ID)

	stream, err := client.Chat.SendMessageStream(ctx, session.ID, &promptrails.SendMessageParams{
		Content: "What is PromptRails?",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	var executionID string
	for stream.Next() {
		switch e := stream.Event().(type) {
		case *promptrails.ExecutionEvent:
			executionID = e.ExecutionID
			fmt.Printf("[execution] %s\n", executionID)
		case *promptrails.ThinkingEvent:
			fmt.Printf("[thinking] %s\n", e.Content)
		case *promptrails.ToolStartEvent:
			fmt.Printf("[tool_start] %s\n", e.Name)
		case *promptrails.ToolEndEvent:
			fmt.Printf("[tool_end] %s — %s\n", e.Name, e.Summary)
		case *promptrails.ContentEvent:
			fmt.Print(e.Content)
		case *promptrails.DoneEvent:
			fmt.Printf("\n[done] %d tokens\n", e.TokenUsage.TotalTokens)
		case *promptrails.ErrorEvent:
			log.Fatalf("[error] %s", e.Message)
		}
	}
	if err := stream.Err(); err != nil {
		log.Fatal(err)
	}

	// ---- Re-attach to the execution stream --------------------------
	//
	// Executions.Stream lets a second process follow an execution that
	// started elsewhere (e.g. Agents.Execute). The backend replays
	// in-flight events and then streams new ones.

	if executionID != "" {
		fmt.Printf("\nReplaying %s\n", executionID)
		tail, err := client.Executions.Stream(ctx, executionID)
		if err != nil {
			log.Fatal(err)
		}
		defer tail.Close()
		for tail.Next() {
			if e, ok := tail.Event().(*promptrails.ContentEvent); ok {
				fmt.Print(e.Content)
			}
			if _, ok := tail.Event().(*promptrails.DoneEvent); ok {
				break
			}
		}
	}

	_ = client.Chat.DeleteSession(ctx, session.ID)
}
