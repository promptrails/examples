import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// Create an agent
const agent = await client.agents.create({
  name: "Customer Support Bot",
  type: "simple",
  description: "Handles customer inquiries",
});
console.log(`Created agent: ${agent.id}`);

// Create a version
const version = await client.agents.createVersion(agent.id, {
  version: "1.0.0",
  config: {
    llm_model_id: "gpt-4o",
    system_prompt: "You are a helpful customer support assistant.",
    temperature: 0.7,
  },
  set_current: true,
  message: "Initial version",
});
console.log(`Created version: ${version.version}`);

// Execute the agent
const result = await client.agents.execute(agent.id, {
  input: { message: "What is your refund policy?" },
});
console.log(`Status: ${result.status}`);
console.log(`Output:`, result.output);
console.log(`Cost: $${result.cost.toFixed(4)}`);
console.log(`Duration: ${result.duration_ms}ms`);

// List agents with filtering
const agents = await client.agents.list({ type: "simple", limit: 10 });
for (const a of agents.data) {
  console.log(`  - ${a.name} (${a.type}) — ${a.status}`);
}

// List versions
const versions = await client.agents.listVersions(agent.id);
for (const v of versions) {
  console.log(`  Version ${v.version} — ${v.id}`);
}

// Promote a version
await client.agents.promoteVersion(agent.id, version.id);

// Delete agent
await client.agents.delete(agent.id);
console.log("Agent deleted.");
