# Python Examples

All examples use [uv](https://docs.astral.sh/uv/) inline script metadata — no virtual environment setup needed.

Every script pins `promptrails>=0.9.0` (API v2), which is published on PyPI — uv
resolves it automatically on first run.

## Prerequisites

```bash
pip install uv  # or: brew install uv
export PROMPTRAILS_API_KEY="pr_key_..."
```

## Running

```bash
cd python/basic && uv run main.py
cd python/agents && uv run main.py
cd python/chat && uv run main.py
# ... etc
```

## Examples

| Folder | Description |
|--------|-------------|
| [basic/](basic/) | Client setup, async usage, error handling |
| [agents/](agents/) | Create, execute, version `agent` / `workflow` agents |
| [prompts/](prompts/) | Content-only prompt templates and versioning |
| [chat/](chat/) | Multi-turn chat sessions |
| [streaming/](streaming/) | Live SSE events for chat turns and executions |
| [traces/](traces/) | Tracing, span trees, and usage summaries |
| [executions/](executions/) | History, execution tree, and the approval inbox |
| [costs/](costs/) | Usage & cost aggregates via `traces.get_summary` |
| [mcp-tools/](mcp-tools/) | MCP tool management |
| [agent-triggers/](agent-triggers/) | Event-driven agent execution |
| [a2a/](a2a/) | Agent-to-Agent communication |
| [assets/](assets/) | Manage stored assets |
