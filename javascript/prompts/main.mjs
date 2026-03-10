import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// Create a prompt
const prompt = await client.prompts.create({
  name: "Support Classifier",
  description: "Classifies customer support tickets",
});
console.log(`Created prompt: ${prompt.id}`);

// Create a version with template
const version = await client.prompts.createVersion(prompt.id, {
  version: "1.0.0",
  system_prompt: "You are a support ticket classifier.",
  user_prompt: "Classify this ticket: {{ message }}",
  llm_model_id: "gpt-4o",
  temperature: 0.3,
  max_tokens: 100,
  set_current: true,
  message: "Initial classifier version",
});

// Run the prompt
const response = await client.prompts.runPrompt(prompt.id, {
  system_prompt: "You are a support ticket classifier.",
  user_prompt: "Classify: I want a refund for my order",
  llm_model_id: "gpt-4o",
  temperature: 0.3,
});
console.log(`Response: ${response.content}`);
console.log(`Cost: $${response.cost.toFixed(4)}`);

// List prompts
const prompts = await client.prompts.list({ limit: 10 });
for (const p of prompts.data) {
  console.log(`  - ${p.name} (${p.id})`);
}

// Clean up
await client.prompts.delete(prompt.id);
