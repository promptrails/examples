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

	// List installed MCP tools
	tools, err := client.MCPTools.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tools.Data {
		fmt.Printf("  %s (%s) — %s\n", t.Name, t.Type, t.Status)
	}

	// Create a new MCP tool
	tool, err := client.MCPTools.Create(ctx, &promptrails.CreateMCPToolParams{
		Name:      "Web Search",
		ServerURL: "https://mcp.example.com/brave-search",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created tool: %s\n", tool.ID)

	// Clean up
	_ = client.MCPTools.Delete(ctx, tool.ID)
	fmt.Println("Tool deleted.")
}
