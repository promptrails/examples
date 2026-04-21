# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.3.0"]
# ///
"""Agent management and execution — typed AgentConfig classes.

The SDK exposes one class per agent type (SimpleAgentConfig,
ChainAgentConfig, MultiAgentConfig, WorkflowAgentConfig,
CompositeAgentConfig). Each class injects the ``type`` discriminator
on serialization so you never build the config JSON by hand.
"""

import os

from promptrails import (
    ChainAgentConfig,
    PromptLink,
    PromptRails,
    SimpleAgentConfig,
)

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Agent configs reference one or more prompts by id, so we need a prompt
# first. A real project would probably reuse an existing prompt rather
# than inlining this step.
prompt = client.prompts.create(
    name="Support reply",
    description="Answers a customer question politely",
)
client.prompts.create_version(
    prompt.id,
    system_prompt="You are a helpful customer support assistant.",
    user_prompt="Customer asks: {{ message }}",
    llm_model_id="gpt-4o",
    temperature=0.7,
    set_current=True,
    message="Initial version",
)

# Create an agent and attach a SimpleAgentConfig version. The typed
# config validates shape locally and .to_dict() — called by the SDK —
# writes {"type": "simple", "prompt_id": ...} to the backend.
agent = client.agents.create(
    name="Customer Support Bot",
    type="simple",
    description="Handles customer inquiries",
)
simple_cfg = SimpleAgentConfig(
    prompt_id=prompt.id,
    temperature=0.7,
)
version = client.agents.create_version(
    agent.id,
    version="1.0.0",
    config=simple_cfg,
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

# Chain multiple prompts together — each step feeds the next. The
# ``role`` and ``sort_order`` fields on PromptLink identify steps.
extract = client.prompts.create(name="Extract", description="Extract entities")
client.prompts.create_version(
    extract.id,
    system_prompt="Extract named entities as a JSON list.",
    user_prompt="{{ text }}",
    llm_model_id="gpt-4o-mini",
    set_current=True,
    message="v1",
)

chain_agent = client.agents.create(
    name="Extract-then-Summarize",
    type="chain",
    description="Two-step reasoning chain",
)
chain_cfg = ChainAgentConfig(
    prompt_ids=[
        PromptLink(prompt_id=extract.id, role="extract", sort_order=0),
        PromptLink(prompt_id=prompt.id, role="answer", sort_order=1),
    ],
)
client.agents.create_version(
    chain_agent.id,
    version="1.0.0",
    config=chain_cfg,
    set_current=True,
    message="Initial chain",
)

# Housekeeping — list, update, promote, delete
agents = client.agents.list(type="simple", page=1, limit=10)
for a in agents.data:
    print(f"  - {a.name} ({a.type}) — {a.status}")

client.agents.update(agent.id, description="Updated description")

versions = client.agents.list_versions(agent.id)
for v in versions:
    print(f"  Version {v.version} — {v.id}")

client.agents.promote_version(agent.id, version.id)

client.agents.delete(chain_agent.id)
client.agents.delete(agent.id)
client.prompts.delete(extract.id)
client.prompts.delete(prompt.id)
print("Cleanup complete.")

client.close()
