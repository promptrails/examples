# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
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
