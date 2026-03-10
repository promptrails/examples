# Go Examples

## Prerequisites

```bash
export PROMPTRAILS_API_KEY="pr_key_..."
```

## Running

Each folder is a standalone `main` package:

```bash
cd go/basic && go run main.go
cd go/agents && go run main.go
cd go/chat && go run main.go
# ... etc
```

## Examples

| Folder | Description |
|--------|-------------|
| [basic/](basic/) | Client setup, configuration, error handling |
| [agents/](agents/) | Create, execute, version, and manage agents |
| [prompts/](prompts/) | Prompt templates, versioning, and execution |
| [chat/](chat/) | Multi-turn chat sessions |
| [traces/](traces/) | Execution tracing and span trees |
| [executions/](executions/) | Monitor execution history |
| [costs/](costs/) | Track LLM usage costs |
| [mcp-tools/](mcp-tools/) | MCP tool management |
| [webhook-triggers/](webhook-triggers/) | Event-driven agent execution |
| [a2a/](a2a/) | Agent-to-Agent communication |
