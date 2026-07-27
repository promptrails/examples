# JavaScript/TypeScript Examples

> **⚠️ API v2 (unreleased):** `package.json` pins `@promptrails/sdk@^0.9.0`, which is **not yet on
> npm**. Until it publishes, build and link the local sibling checkout (branch `feat/api-v2`) — the
> pinned spec then reconciles against the published package (and lockfile) at release, with no edits
> here:
>
> ```bash
> (cd ../javascript-sdk && npm install && npm run build && npm link)
> cd javascript && npm link @promptrails/sdk
> ```

## Prerequisites

```bash
cd javascript
npm install   # once @promptrails/sdk 0.9.0 is on npm; until then, see the note above
export PROMPTRAILS_API_KEY="pr_key_..."
```

## Running

```bash
node javascript/basic/main.mjs
node javascript/agents/main.mjs
node javascript/chat/main.mjs
# ... etc
```

## Examples

| Folder | Description |
|--------|-------------|
| [basic/](basic/) | Client setup, error handling |
| [agents/](agents/) | Create, execute, version `agent` / `workflow` agents |
| [prompts/](prompts/) | Content-only prompt templates and versioning |
| [chat/](chat/) | Multi-turn chat sessions |
| [streaming/](streaming/) | Live SSE events for chat turns and executions |
| [traces/](traces/) | Tracing, span trees, and usage summaries |
| [executions/](executions/) | History, execution tree, and the approval inbox |
| [costs/](costs/) | Usage & cost aggregates via `traces.getSummary` |
| [mcp-tools/](mcp-tools/) | MCP tool management |
| [agent-triggers/](agent-triggers/) | Event-driven agent execution |
| [a2a/](a2a/) | Agent-to-Agent communication |
| [assets/](assets/) | Manage stored assets |
