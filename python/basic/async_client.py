# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
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
