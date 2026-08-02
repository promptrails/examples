# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Server-Sent Events streaming — live chat turns and executions.

The SDK exposes two streaming entry points:
- ``client.chat.send_message_stream(session_id, content=...)``
- ``client.executions.stream(execution_id)``

Both yield the same typed events. Dispatch on the concrete class —
unknown event names are silently dropped so the client stays
forward-compatible.
"""

import os
import sys

from promptrails import (
    ContentEvent,
    DoneEvent,
    ErrorEvent,
    ExecutionEvent,
    PromptRails,
    ThinkingEvent,
    ToolEndEvent,
    ToolStartEvent,
)

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])
AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")

# ---- Stream a chat turn ----------------------------------------------

session = client.chat.create_session(agent_id=AGENT_ID, title="Streaming demo")
print(f"Session: {session.id}\n")

execution_id = None
for event in client.chat.send_message_stream(
    session.id, content="What is PromptRails?"
):
    if isinstance(event, ExecutionEvent):
        execution_id = event.execution_id
        print(f"[execution] {execution_id}")
    elif isinstance(event, ThinkingEvent):
        print(f"[thinking] {event.content}")
    elif isinstance(event, ToolStartEvent):
        print(f"[tool_start] {event.name}")
    elif isinstance(event, ToolEndEvent):
        print(f"[tool_end] {event.name} — {event.summary}")
    elif isinstance(event, ContentEvent):
        sys.stdout.write(event.content)
        sys.stdout.flush()
    elif isinstance(event, DoneEvent):
        print(f"\n[done] {event.token_usage.total_tokens} tokens")
    elif isinstance(event, ErrorEvent):
        raise RuntimeError(event.message)

# ---- Re-attach to the same execution via executions.stream -----------
#
# Useful when the execution was kicked off outside of chat (e.g.
# agents.execute) and another process wants to observe progress. The
# backend replays any in-flight events and then streams new ones.

if execution_id:
    print(f"\nReplaying {execution_id}")
    for event in client.executions.stream(execution_id):
        if isinstance(event, ContentEvent):
            sys.stdout.write(event.content)
            sys.stdout.flush()
        elif isinstance(event, DoneEvent):
            break

client.chat.delete_session(session.id)
client.close()
