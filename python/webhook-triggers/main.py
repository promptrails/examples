# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Webhook trigger management for event-driven agent execution."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")

# Create a webhook trigger
trigger = client.webhook_triggers.create(
    name="GitHub Push Webhook",
    agent_id=AGENT_ID,
    generate_secret=True,
)
print(f"Trigger created: {trigger.id}")
print(f"Webhook URL: {trigger.url}")
print(f"Secret: {trigger.secret}")

# List triggers
triggers = client.webhook_triggers.list()
for t in triggers.data:
    print(f"  {t.name} — active: {t.is_active}")

# Update trigger
client.webhook_triggers.update(trigger.id, name="Updated Webhook")

# Delete trigger
client.webhook_triggers.delete(trigger.id)
print("Trigger deleted.")

client.close()
