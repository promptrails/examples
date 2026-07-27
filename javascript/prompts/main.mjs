import { PromptRails } from "@promptrails/sdk";

// In v2 a prompt version is content only (system/user text + optional
// input_schema). Model, sampling, tools and output schema live on the agent
// version, not the prompt. Prompts are no longer runnable directly — attach
// one to an agent and call agents.execute (see the agents example).

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// Create a prompt
const prompt = await client.prompts.create({
  name: "Support Classifier",
  description: "Classifies customer support tickets",
});
console.log(`Created prompt: ${prompt.id}`);

// Create a content-only version — no llm_model_id / temperature / max_tokens.
const version = await client.prompts.createVersion(prompt.id, {
  version: "1.0.0",
  system_prompt: "You are a support ticket classifier.",
  user_prompt: "Classify this ticket: {{ message }}",
  input_schema: {
    type: "object",
    properties: { message: { type: "string" } },
  },
  set_current: true,
  message: "Initial classifier version",
});
console.log(`Created version: ${version.version}`);

// Preview renders the template with sample input (no LLM call).
const preview = await client.prompts.preview(prompt.id, {
  input: { message: "I want a refund for my order" },
});
console.log(`Rendered preview:`, preview);

// List prompts
const prompts = await client.prompts.list({ limit: 10 });
for (const p of prompts.data) {
  console.log(`  - ${p.name} (${p.id})`);
}

// Clean up
await client.prompts.delete(prompt.id);
