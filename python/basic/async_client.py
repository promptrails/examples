# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Async client usage."""

import asyncio
import os
from promptrails import AsyncPromptRails


async def main():
    async with AsyncPromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"]) as client:
        agents = await client.agents.list()
        print(f"Found {agents.meta.total} agents.")

        for agent in agents.data:
            print(f"  - {agent.name} ({agent.type})")


asyncio.run(main())
