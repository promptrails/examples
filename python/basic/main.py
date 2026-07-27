# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Basic client setup and configuration."""

import os
from promptrails import PromptRails, AsyncPromptRails

# Synchronous client
client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# List agents to verify connection
agents = client.agents.list()
print(f"Connected! Found {agents.meta.total} agents.")

client.close()


# Using context manager (recommended)
with PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"]) as client:
    agents = client.agents.list()
    for agent in agents.data:
        print(f"  - {agent.name} ({agent.type})")


# Custom configuration
client = PromptRails(
    api_key=os.environ["PROMPTRAILS_API_KEY"],
    base_url="https://api.promptrails.ai",  # custom endpoint
    timeout=60.0,  # 60 second timeout
    max_retries=5,  # retry up to 5 times on failures
)
client.close()
