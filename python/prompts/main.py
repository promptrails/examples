# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Prompt template management and execution."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# Create a prompt
prompt = client.prompts.create(
    name="Support Classifier",
    description="Classifies customer support tickets",
)
print(f"Created prompt: {prompt.id}")

# Create a version with template
version = client.prompts.create_version(
    prompt.id,
    version="1.0.0",
    system_prompt="You are a support ticket classifier.",
    user_prompt="Classify this ticket: {{ message }}",
    llm_model_id="gpt-4o",
    temperature=0.3,
    max_tokens=100,
    input_schema={"type": "object", "properties": {"message": {"type": "string"}}},
    set_current=True,
    message="Initial classifier version",
)
print(f"Created version: {version.version}")

# Run the prompt directly
response = client.prompts.run_prompt(
    prompt.id,
    data={
        "system_prompt": "You are a support ticket classifier.",
        "user_prompt": "Classify: I want a refund for my order",
        "llm_model_id": "gpt-4o",
        "temperature": 0.3,
    },
)
print(f"Response: {response.content}")
print(f"Tokens: {response.token_usage}")
print(f"Cost: ${response.cost:.4f}")

# List all prompts
prompts = client.prompts.list(page=1, limit=10)
for p in prompts.data:
    print(f"  - {p.name} ({p.id})")

# Clean up
client.prompts.delete(prompt.id)

client.close()
