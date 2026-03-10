# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Multi-turn chat conversations with agents."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")

# Create a chat session
session = client.chat.create_session(
    agent_id=AGENT_ID,
    title="Support conversation",
    metadata={"source": "example"},
)
print(f"Session created: {session.id}")

# Send messages
msg1 = client.chat.send_message(session.id, content="What is PromptRails?")
print(f"Assistant: {msg1.content}")
print(f"  Tokens: {msg1.token_count}, Cost: ${msg1.cost:.4f}")

msg2 = client.chat.send_message(session.id, content="How do I create an agent?")
print(f"Assistant: {msg2.content}")

# List message history
messages = client.chat.list_messages(session.id)
for msg in messages.data:
    print(f"  [{msg.role}] {msg.content[:80]}...")

# List all sessions
sessions = client.chat.list_sessions()
for s in sessions.data:
    print(f"  Session: {s.title} ({s.id})")

# Clean up
client.chat.delete_session(session.id)

client.close()
