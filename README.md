# PromptRails SDK Examples

Code examples for the [PromptRails](https://promptrails.ai) SDKs — Python, JavaScript/TypeScript, and Go.

> **API v2** — these examples target the v2 SDKs (Python `0.9.0`, JavaScript `0.9.0`, Go `0.7.0`).
> Agents are now `agent` or `workflow` only; model/sampling live on the agent version; the
> `costs`, `scores`, `media`, `templates`, `sessions` and `dashboard` resources were removed
> (use `traces.get_summary` for cost/usage).
>
> **⚠️ Unreleased dependency pins:** the examples reference the **published** package specs
> (`promptrails>=0.9.0`, `@promptrails/sdk@^0.9.0`, and a `feat/api-v2` pseudo-version of the
> go-sdk). Those v2 SDKs are not on PyPI/npm yet, so until they publish you must resolve them from
> the local sibling checkouts to run the examples (uv `--with-editable`, `npm link`, `go get
> @feat/api-v2`) — see each per-language README. Nothing in the manifests points at a local path;
> the pins/lockfiles reconcile automatically once the SDKs publish and the Go pseudo-version is
> repinned to its tag.

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
| **Agents** — Create, execute, and manage `agent` / `workflow` agents | [View](python/agents/) | [View](javascript/agents/) | [View](go/agents/) |
| **Prompts** — Content-only prompt templates & versioning | [View](python/prompts/) | [View](javascript/prompts/) | [View](go/prompts/) |
| **Chat** — Multi-turn conversations | [View](python/chat/) | [View](javascript/chat/) | [View](go/chat/) |
| **Streaming** — Live SSE events for chat turns and executions | [View](python/streaming/) | [View](javascript/streaming/) | [View](go/streaming/) |
| **Traces** — Tracing, span trees, and usage summaries | [View](python/traces/) | [View](javascript/traces/) | [View](go/traces/) |
| **Executions** — History, execution tree, and the approval inbox | [View](python/executions/) | [View](javascript/executions/) | [View](go/executions/) |
| **Costs** — Usage & cost aggregates via `traces.get_summary` | [View](python/costs/) | [View](javascript/costs/) | [View](go/costs/) |
| **MCP Tools** — External tool integrations | [View](python/mcp-tools/) | [View](javascript/mcp-tools/) | [View](go/mcp-tools/) |
| **Agent Triggers** — Event-driven agent execution | [View](python/agent-triggers/) | [View](javascript/agent-triggers/) | [View](go/agent-triggers/) |
| **A2A** — Agent-to-Agent communication | [View](python/a2a/) | [View](javascript/a2a/) | [View](go/a2a/) |
| **Assets** — Manage stored assets | [View](python/assets/) | [View](javascript/assets/) | [View](go/assets/) |

## SDK Installation

> These are the target published versions. Until they ship, the examples resolve the
> unreleased SDKs locally (see the ⚠️ note above and the per-language READMEs).

**Python:**
```bash
pip install "promptrails>=0.9.0"
```

**JavaScript/TypeScript:**
```bash
npm install @promptrails/sdk@^0.9.0
```

**Go:**
```bash
go get github.com/promptrails/go-sdk@v0.7.0
```

## Links

- [Documentation](https://promptrails.ai/docs)
- [Python SDK](https://github.com/promptrails/python-sdk)
- [JavaScript SDK](https://github.com/promptrails/javascript-sdk)
- [Go SDK](https://github.com/promptrails/go-sdk)
