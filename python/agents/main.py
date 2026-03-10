# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Agent management and execution."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Create an agent
agent = client.agents.create(
    name="Customer Support Bot",
    type="simple",
    description="Handles customer inquiries",
)
print(f"Created agent: {agent.id}")

# Create a version with configuration
version = client.agents.create_version(
    agent.id,
    version="1.0.0",
    config={
        "llm_model_id": "gpt-4o",
        "system_prompt": "You are a helpful customer support assistant.",
        "temperature": 0.7,
    },
    set_current=True,
    message="Initial version",
)
print(f"Created version: {version.version}")

# Execute the agent
result = client.agents.execute(
    agent.id,
    input={"message": "What is your refund policy?"},
)
print(f"Status: {result.status}")
print(f"Output: {result.output}")
print(f"Cost: ${result.cost:.4f}")
print(f"Duration: {result.duration_ms}ms")
print(f"Trace ID: {result.trace_id}")

# List agents with filtering
agents = client.agents.list(type="simple", page=1, limit=10)
for a in agents.data:
    print(f"  - {a.name} ({a.type}) — {a.status}")

# Update agent
updated = client.agents.update(agent.id, description="Updated description")

# List versions
versions = client.agents.list_versions(agent.id)
for v in versions:
    print(f"  Version {v.version} — {v.id}")

# Promote a specific version
client.agents.promote_version(agent.id, version.id)

# Delete agent
client.agents.delete(agent.id)
print("Agent deleted.")

client.close()
