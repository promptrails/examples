# Python Examples

All examples use [uv](https://docs.astral.sh/uv/) inline script metadata — no virtual environment setup needed.

## Prerequisites

```bash
pip install uv  # or: brew install uv
export PROMPTRAILS_API_KEY="pr_key_..."
```

## Running

```bash
uv run python/basic/main.py
uv run python/agents/main.py
uv run python/chat/main.py
# ... etc
```

## Examples

| Folder | Description |
|--------|-------------|
| [basic/](basic/) | Client setup, async usage, error handling |
| [agents/](agents/) | Create, execute, version, and manage agents |
| [prompts/](prompts/) | Prompt templates, versioning, and execution |
| [chat/](chat/) | Multi-turn chat sessions |
| [traces/](traces/) | Execution tracing and span trees |
| [executions/](executions/) | Monitor execution history |
| [costs/](costs/) | Track LLM usage costs |
| [mcp-tools/](mcp-tools/) | MCP tool management |
| [webhook-triggers/](webhook-triggers/) | Event-driven agent execution |
| [a2a/](a2a/) | Agent-to-Agent communication |
