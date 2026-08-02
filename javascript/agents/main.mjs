import { PromptRails } from "@promptrails/sdk";

// PromptRails v2 has two agent kinds: "agent" (a prompt + optional tools /
// sub-agents) and "workflow" (a deterministic DAG). Model + sampling
// (model_config), the run budget, approval policy, tools and sub-agents are
// version-scoped — they sit alongside `config`, not inside it. Prompts are
// content-only and carry no model config.

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// An agent references a prompt by id — create the (content-only) prompt first.
const prompt = await client.prompts.create({
  name: "Support reply",
  description: "Answers a customer question politely",
});
await client.prompts.createVersion(prompt.id, {
  version: "1.0.0",
  system_prompt: "You are a helpful customer support assistant.",
  user_prompt: "Customer asks: {{ message }}",
  set_current: true,
  message: "Initial version",
});

// Create an "agent" and attach a version. Model + sampling + budget ride on
// the version as siblings of `config`.
const agent = await client.agents.create({
  name: "Customer Support Bot",
  type: "agent",
  description: "Handles customer inquiries",
});
const version = await client.agents.createVersion(agent.id, {
  version: "1.0.0",
  config: { type: "agent", prompt_id: prompt.id },
  model_config: { model_id: "gpt-4o", temperature: 0.7 },
  run_budget: { max_cost: 0.5, max_total_tokens: 20000 },
  set_current: true,
  message: "Initial version",
});
console.log(`Agent ${agent.id} @ ${version.version}`);

// Execute the agent
const result = await client.agents.execute(agent.id, {
  input: { message: "What is your refund policy?" },
});
console.log(`Status: ${result.status}`);
console.log(`Output:`, result.output);
console.log(`Cost: $${result.cost.toFixed(4)}  Duration: ${result.duration_ms}ms`);

// A "workflow" agent runs a deterministic DAG — each node pins a prompt and
// declares its dependencies.
const extract = await client.prompts.create({
  name: "Extract",
  description: "Extract entities",
});
await client.prompts.createVersion(extract.id, {
  version: "1.0.0",
  system_prompt: "Extract named entities as a JSON list.",
  user_prompt: "{{ text }}",
  set_current: true,
  message: "v1",
});

const workflowAgent = await client.agents.create({
  name: "Extract-then-Answer",
  type: "workflow",
  description: "Two-step workflow",
});
await client.agents.createVersion(workflowAgent.id, {
  version: "1.0.0",
  config: {
    type: "workflow",
    nodes: [
      { id: "extract", prompt_id: extract.id, depends_on: [] },
      { id: "answer", prompt_id: prompt.id, depends_on: ["extract"] },
    ],
  },
  model_config: { model_id: "gpt-4o-mini" },
  set_current: true,
  message: "Initial workflow",
});

// List agents with filtering
const agents = await client.agents.list({ type: "agent", limit: 10 });
for (const a of agents.data) {
  console.log(`  - ${a.name} (${a.type}) — ${a.status}`);
}

// List versions + promote
const versions = await client.agents.listVersions(agent.id);
for (const v of versions) {
  console.log(`  Version ${v.version} — ${v.id}`);
}
await client.agents.promoteVersion(agent.id, version.id);

// Clean up
await client.agents.delete(workflowAgent.id);
await client.agents.delete(agent.id);
await client.prompts.delete(extract.id);
await client.prompts.delete(prompt.id);
console.log("Cleanup complete.");
