# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Monitor executions + drive the approval inbox — API v2.

Approvals are now execution-scoped. A run that hits an approval-gated tool /
sub-agent parks at status ``waiting_approval``; resume it with
``executions.approve`` / ``executions.deny``. ``executions.tree`` returns the
full parent→child execution tree and ``executions.cancel`` requests
cooperative cancellation (status ``cancel_requested``).
"""

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

# Filter by status
failed = client.executions.list(status="failed", limit=5)
print(f"Failed executions: {failed.meta.total}")
for ex in failed.data:
    print(f"  {ex.id}: {ex.error}")

# Get execution details, then walk its full child tree
if executions.data:
    detail = client.executions.get(executions.data[0].id)
    print("Execution detail:")
    print(f"  Input: {detail.input}")
    print(f"  Output: {detail.output}")
    print(f"  Token usage: {detail.token_usage}")

    tree = client.executions.tree(detail.id)
    print(f"  Child executions: {len(tree.children)}")

# Approval inbox — runs parked at waiting_approval
inbox = client.executions.approval_inbox(page=1, limit=10)
print(f"Awaiting approval: {inbox.meta.total}")
for ex in inbox.data:
    print(f"  {ex.id} — {ex.status}")
    # Approve to resume, or deny to resume with a denial:
    # client.executions.approve(ex.id, reason="Looks safe")
    # client.executions.deny(ex.id, reason="Not allowed")

# Cancel a still-running execution
# client.executions.cancel("exec-id")

client.close()
