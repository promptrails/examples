# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Error handling patterns."""

import os
from promptrails import PromptRails
from promptrails import (
    PromptRailsError,
    NotFoundError,
    ValidationError,
    RateLimitError,
    UnauthorizedError,
)

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Handle specific errors
try:
    agent = client.agents.get("nonexistent-id")
except NotFoundError as e:
    print(f"Agent not found: {e.message}")
except ValidationError as e:
    print(f"Invalid request: {e.message}")
except RateLimitError:
    print("Rate limited — retry after a delay")
except UnauthorizedError:
    print("Invalid API key")
except PromptRailsError as e:
    print(f"API error: {e.message}")

client.close()
