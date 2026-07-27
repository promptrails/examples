# Go Examples

> **⚠️ API v2 (unreleased):** `go.mod` requires a **pseudo-version** of the go-sdk `feat/api-v2`
> branch (e.g. `v0.7.1-0.YYYYMMDDhhmmss-<hash>`) — a real module reference resolved with
> `go get github.com/promptrails/go-sdk@feat/api-v2`, not a local path. Repin to the tagged release
> once published: `go get github.com/promptrails/go-sdk@v0.7.0`.

## Prerequisites

```bash
export PROMPTRAILS_API_KEY="pr_key_..."
```

## Running

All examples share one Go module (`go/go.mod`); each folder is its own `main` package:

```bash
cd go
go run ./basic
go run ./basic/error-handling
go run ./agents
go run ./chat
# ... etc
```

Build everything at once with `go build ./...`.

## Examples

| Folder | Description |
|--------|-------------|
| [basic/](basic/) | Client setup and configuration |
| [basic/error-handling/](basic/error-handling/) | Typed error handling |
| [agents/](agents/) | Create, execute, version `agent` / `workflow` agents |
| [prompts/](prompts/) | Content-only prompt templates and versioning |
| [chat/](chat/) | Multi-turn chat sessions |
| [streaming/](streaming/) | Live SSE events for chat turns and executions |
| [traces/](traces/) | Tracing, span trees, and usage summaries |
| [executions/](executions/) | History, execution tree, and the approval inbox |
| [costs/](costs/) | Usage & cost aggregates via `Traces.GetSummary` |
| [mcp-tools/](mcp-tools/) | MCP tool management |
| [agent-triggers/](agent-triggers/) | Event-driven agent execution |
| [a2a/](a2a/) | Agent-to-Agent communication |
| [assets/](assets/) | Manage stored assets |
