# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Track LLM usage costs."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Get workspace cost summary
summary = client.costs.get_summary()
print(f"Workspace costs:")
print(f"  Total cost: ${summary.total_cost:.2f}")
print(f"  Total tokens: {summary.total_tokens}")
print(f"  Total executions: {summary.total_executions}")

# Filter by date range
summary_30d = client.costs.get_summary(from_date="2026-02-01", to_date="2026-03-01")
print(f"February costs: ${summary_30d.total_cost:.2f}")

# Get costs for a specific agent
AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")
agent_costs = client.costs.get_agent_summary(AGENT_ID)
print(f"Agent costs: ${agent_costs.total_cost:.2f}")

client.close()
