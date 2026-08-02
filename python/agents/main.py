# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Agent management and execution — API v2.

PromptRails v2 has exactly two agent kinds:

* ``agent``    — a prompt (+ optional tools / sub-agents), built with
  :class:`PromptAgentConfig`.
* ``workflow`` — a deterministic DAG of nodes, built with
  :class:`WorkflowAgentConfig`.

Model + sampling (``ModelConfig``), the execution-tree budget
(``RunBudget``), the approval policy, tools and sub-agents are *version*-scoped
— you pass them to ``create_version`` alongside ``config``, not inside it.
Prompts are content-only (system/user text); they carry no model config.
"""

import os

from promptrails import (
    ModelConfig,
    PromptAgentConfig,
    PromptRails,
    RunBudget,
    WorkflowAgentConfig,
    WorkflowNode,
)

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# An agent references a prompt by id, so create the (content-only) prompt
# first. Model/sampling do NOT live here anymore — they live on the agent
# version below.
prompt = client.prompts.create(
    name="Support reply",
    description="Answers a customer question politely",
)
client.prompts.create_version(
    prompt.id,
    version="1.0.0",
    system_prompt="You are a helpful customer support assistant.",
    user_prompt="Customer asks: {{ message }}",
    set_current=True,
    message="Initial version",
)

# Create an ``agent`` and attach a version. PromptAgentConfig serializes to
# {"type": "agent", "prompt_id": ...}. Model + sampling + budget ride on the
# version as siblings of ``config``.
agent = client.agents.create(
    name="Customer Support Bot",
    type="agent",
    description="Handles customer inquiries",
)
version = client.agents.create_version(
    agent.id,
    version="1.0.0",
    config=PromptAgentConfig(prompt_id=prompt.id),
    model_config=ModelConfig(model_id="gpt-4o", temperature=0.7),
    run_budget=RunBudget(max_cost=0.50, max_total_tokens=20_000),
    set_current=True,
    message="Initial version",
)
print(f"Agent {agent.id} @ {version.version}")

# Execute the agent
result = client.agents.execute(
    agent.id,
    input={"message": "What is your refund policy?"},
)
print(f"Status: {result.status}")
print(f"Output: {result.output}")
print(f"Cost: ${result.cost:.4f}  Duration: {result.duration_ms}ms")
print(f"Trace ID: {result.trace_id}")

# A ``workflow`` agent runs a deterministic DAG. Each node pins a prompt and
# declares its dependencies; the runner topologically orders them.
extract = client.prompts.create(name="Extract", description="Extract entities")
client.prompts.create_version(
    extract.id,
    version="1.0.0",
    system_prompt="Extract named entities as a JSON list.",
    user_prompt="{{ text }}",
    set_current=True,
    message="v1",
)

workflow_agent = client.agents.create(
    name="Extract-then-Answer",
    type="workflow",
    description="Two-step workflow",
)
client.agents.create_version(
    workflow_agent.id,
    version="1.0.0",
    config=WorkflowAgentConfig(
        nodes=[
            WorkflowNode(id="extract", prompt_id=extract.id),
            WorkflowNode(id="answer", prompt_id=prompt.id, depends_on=["extract"]),
        ],
    ),
    model_config=ModelConfig(model_id="gpt-4o-mini"),
    set_current=True,
    message="Initial workflow",
)

# Housekeeping — list, update, promote, delete
agents = client.agents.list(type="agent", page=1, limit=10)
for a in agents.data:
    print(f"  - {a.name} ({a.type}) — {a.status}")

client.agents.update(agent.id, description="Updated description")

versions = client.agents.list_versions(agent.id)
for v in versions:
    print(f"  Version {v.version} — {v.id}")

client.agents.promote_version(agent.id, version.id)

client.agents.delete(workflow_agent.id)
client.agents.delete(agent.id)
client.prompts.delete(extract.id)
client.prompts.delete(prompt.id)
print("Cleanup complete.")

client.close()
