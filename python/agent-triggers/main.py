# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.4"]
# ///
"""Agent trigger management — generic webhook, Slack, Telegram, Teams,
WhatsApp, or cron schedule. All six sources share one create() method.

Run with uv (recommended):
    uv run examples/python/agent-triggers/main.py
"""

import os

from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")

# 1. Generic webhook trigger (default source)
trigger = client.agent_triggers.create(
    name="GitHub Push Webhook",
    agent_id=AGENT_ID,
    generate_secret=True,
)
print(f"Generic trigger created: {trigger.id}")
print(f"  Token:  {trigger.token}")
print(f"  Secret: {trigger.secret}")

# 2. Slack trigger — point at workspace credentials for signing secret + bot token
slack_trigger = client.agent_triggers.create(
    name="Slack #incidents",
    agent_id=AGENT_ID,
    source="slack",
    source_config={
        "signing_secret_credential_id": "<credential id with the Slack signing secret>",
        "bot_token_credential_id": "<credential id with the bot token>",
    },
)
print(f"Slack trigger created: {slack_trigger.id}")

# 3. Schedule trigger — fires on a cron expression
schedule_trigger = client.agent_triggers.create(
    name="Daily digest at 9am",
    agent_id=AGENT_ID,
    source="schedule",
    source_config={
        "cron_expression": "0 9 * * *",
        "next_run_at": "2026-05-22T09:00:00Z",
    },
)
print(f"Schedule trigger created: {schedule_trigger.id}")

# List + update + delete still work the same way
triggers = client.agent_triggers.list()
for t in triggers.data:
    print(f"  {t.name} ({t.source}) — active: {t.is_active}")

client.agent_triggers.update(trigger.id, name="GitHub Push Webhook (renamed)")
client.agent_triggers.delete(trigger.id)
client.agent_triggers.delete(slack_trigger.id)
client.agent_triggers.delete(schedule_trigger.id)
print("Triggers cleaned up.")

client.close()
