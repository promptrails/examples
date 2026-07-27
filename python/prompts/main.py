# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Prompt template management — API v2.

In v2 a prompt version is *content only*: system/user text (+ optional
input_schema). Model, sampling, tools, output schema and cache TTL live on the
agent version, not the prompt (see the ``agents`` example). Prompts are no
longer runnable directly — attach one to an agent and call ``agents.execute``.
"""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Create a prompt
prompt = client.prompts.create(
    name="Support Classifier",
    description="Classifies customer support tickets",
)
print(f"Created prompt: {prompt.id}")

# Create a content-only version. No llm_model_id / temperature / max_tokens
# here — those belong to the agent version that references this prompt.
version = client.prompts.create_version(
    prompt.id,
    version="1.0.0",
    system_prompt="You are a support ticket classifier.",
    user_prompt="Classify this ticket: {{ message }}",
    input_schema={"type": "object", "properties": {"message": {"type": "string"}}},
    set_current=True,
    message="Initial classifier version",
)
print(f"Created version: {version.version}")

# Preview renders the template with sample input (no LLM call).
preview = client.prompts.preview(
    prompt.id,
    input={"message": "I want a refund for my order"},
)
print(f"Rendered preview: {preview}")

# List all prompts
prompts = client.prompts.list(page=1, limit=10)
for p in prompts.data:
    print(f"  - {p.name} ({p.id})")

# Clean up
client.prompts.delete(prompt.id)

client.close()
