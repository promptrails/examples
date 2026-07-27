# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""MCP (Model Context Protocol) tool management."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# List available MCP templates
templates = client.mcp_templates.list()
for t in templates.data:
    print(f"  {t.name} ({t.slug}) — {t.category}")

# Install a tool from template
tool = client.mcp_templates.install(
    "brave-search",
    name="Web Search",
    parameters={"api_key": "your-brave-api-key"},
)
print(f"Installed tool: {tool.id}")

# List installed tools
tools = client.mcp_tools.list()
for t in tools.data:
    print(f"  {t.name} ({t.type}) — {t.status}")

# Discover available tool functions
if tools.data:
    discovered = client.mcp_tools.discover_tools(tools.data[0].id)
    for dt in discovered.tools:
        print(f"  Tool: {dt.name} — {dt.description}")

# Call a tool
result = client.mcp_tools.call_tool(
    tool.id,
    tool_name="search",
    arguments={"query": "PromptRails AI platform"},
)
print(f"Result: {result.content}")

# Clean up
client.mcp_tools.delete(tool.id)

client.close()
