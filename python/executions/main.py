# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Monitor agent execution history."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# List recent executions
executions = client.executions.list(page=1, limit=10)
print(f"Total executions: {executions.meta.total}")

for ex in executions.data:
    print(f"  {ex.id} — {ex.status}")
    print(f"    Agent: {ex.agent_id}")
    print(f"    Duration: {ex.duration_ms}ms, Cost: ${ex.cost:.4f}")
    print()

# Filter by status
failed = client.executions.list(status="failed", limit=5)
print(f"Failed executions: {failed.meta.total}")
for ex in failed.data:
    print(f"  {ex.id}: {ex.error}")

# Filter by agent
AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")
agent_execs = client.executions.list(agent_id=AGENT_ID, limit=5)
print(f"Executions for agent: {agent_execs.meta.total}")

# Get execution details
if executions.data:
    detail = client.executions.get(executions.data[0].id)
    print(f"Execution detail:")
    print(f"  Input: {detail.input}")
    print(f"  Output: {detail.output}")
    print(f"  Token usage: {detail.token_usage}")

client.close()
