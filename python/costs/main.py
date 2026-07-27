# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Usage & cost reporting — API v2.

The dedicated ``costs`` resource is gone. Cost, token and latency aggregates
now come from ``traces.get_summary(**filters)``, which accepts the same
filters as ``traces.list`` plus ``date_from`` / ``date_to``, ``status``,
``level``, ``model_name``, ``agent_id``, ``session_id`` and ``execution_id``.
"""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Workspace-wide summary
summary = client.traces.get_summary()
print("Workspace usage:")
print(f"  Total cost:     ${summary.total_cost:.2f}")
print(f"  Total tokens:   {summary.total_tokens}")
print(f"  Total traces:   {summary.total_traces}")
print(f"  Errors:         {summary.error_count}")
print(f"  Avg duration:   {summary.avg_duration_ms:.0f}ms")
print(f"  Unique models:  {summary.unique_models}")

# Filter by date range
feb = client.traces.get_summary(date_from="2026-02-01", date_to="2026-03-01")
print(f"\nFebruary cost: ${feb.total_cost:.2f} over {feb.total_traces} traces")

# Scope to a single agent
AGENT_ID = os.environ.get("AGENT_ID", "your-agent-id")
agent_summary = client.traces.get_summary(agent_id=AGENT_ID)
print(f"Agent cost: ${agent_summary.total_cost:.2f}")

client.close()
