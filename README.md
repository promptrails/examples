# PromptRails SDK Examples

Code examples for the [PromptRails](https://promptrails.ai) SDKs — Python, JavaScript/TypeScript, and Go.

## Prerequisites

- A PromptRails account and API key
- Set your API key as an environment variable:

```bash
export PROMPTRAILS_API_KEY="pr_key_..."
```

## Examples

| Example | Python | JavaScript | Go |
|---------|--------|------------|-----|
| **Basic** — Client setup and configuration | [View](python/basic/) | [View](javascript/basic/) | [View](go/basic/) |
| **Agents** — Create, execute, and manage agents | [View](python/agents/) | [View](javascript/agents/) | [View](go/agents/) |
| **Prompts** — Template management and execution | [View](python/prompts/) | [View](javascript/prompts/) | [View](go/prompts/) |
| **Chat** — Multi-turn conversations | [View](python/chat/) | [View](javascript/chat/) | [View](go/chat/) |
| **Streaming** — Live SSE events for chat turns and executions | [View](python/streaming/) | [View](javascript/streaming/) | [View](go/streaming/) |
| **Traces** — Execution tracing and observability | [View](python/traces/) | [View](javascript/traces/) | [View](go/traces/) |
| **Executions** — Monitor execution history | [View](python/executions/) | [View](javascript/executions/) | [View](go/executions/) |
| **Costs** — Track LLM usage costs | [View](python/costs/) | [View](javascript/costs/) | [View](go/costs/) |
| **MCP Tools** — External tool integrations | [View](python/mcp-tools/) | [View](javascript/mcp-tools/) | [View](go/mcp-tools/) |
| **Agent Triggers** — Event-driven agent execution | [View](python/agent-triggers/) | [View](javascript/agent-triggers/) | [View](go/agent-triggers/) |
| **A2A** — Agent-to-Agent communication | [View](python/a2a/) | [View](javascript/a2a/) | [View](go/a2a/) |
| **Media Studio** — Generate images, speech, and video | [View](python/media-studio/) | [View](javascript/media-studio/) | [View](go/media-studio/) |
| **Assets** — Manage generated media assets | [View](python/assets/) | [View](javascript/assets/) | [View](go/assets/) |
| **Media Models** — Browse available media models | [View](python/media-models/) | [View](javascript/media-models/) | [View](go/media-models/) |

## SDK Installation

**Python:**
```bash
pip install "promptrails>=0.3.0"
```

**JavaScript/TypeScript:**
```bash
npm install @promptrails/sdk@^0.3.1
```

**Go:**
```bash
go get github.com/promptrails/go-sdk@v0.3.1
```

## Links

- [Documentation](https://promptrails.ai/docs)
- [Python SDK](https://github.com/promptrails/python-sdk)
- [JavaScript SDK](https://github.com/promptrails/javascript-sdk)
- [Go SDK](https://github.com/promptrails/go-sdk)
