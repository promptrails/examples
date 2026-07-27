# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Agent-to-Agent (A2A) communication protocol."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")

# Get agent card (A2A discovery)
card = client.a2a.get_agent_card(AGENT_ID)
print(f"Agent: {card.name}")
print(f"Description: {card.description}")
print(f"Skills: {card.skills}")

# Send a message (blocking — waits for completion)
task = client.a2a.send_message(
    AGENT_ID,
    message="Analyze the sales data for Q1 2026",
    blocking=True,
)
print(f"Task: {task.id}")
print(f"State: {task.status.state if task.status else 'unknown'}")
# Results come back as artifacts, each a list of parts.
for artifact in task.artifacts:
    for part in artifact.parts:
        print(f"Result: {part.text}")

# Send a non-blocking message
task = client.a2a.send_message(
    AGENT_ID,
    message="Generate a monthly report",
    blocking=False,
)
state = task.status.state if task.status else "unknown"
print(f"Task created: {task.id} (state: {state})")

# Poll for completion
task = client.a2a.get_task(task.id)
print(f"Task state: {task.status.state if task.status else 'unknown'}")

# List tasks
tasks = client.a2a.list_tasks(limit=5)
for t in tasks.data:
    print(f"  {t.id} — {t.status}")

# Cancel a running task
client.a2a.cancel_task(task.id)

client.close()
