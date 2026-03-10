import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

const AGENT_ID = process.env.AGENT_ID || "your-agent-id";

// Get agent card (A2A discovery)
const card = await client.a2a.getAgentCard(AGENT_ID);
console.log(`Agent: ${card.name}`);
console.log(`Description: ${card.description}`);

// Send a blocking message
const task = await client.a2a.sendMessage(
  AGENT_ID,
  "Analyze the sales data for Q1 2026",
  { blocking: true },
);
console.log(`Task: ${task.id}`);
console.log(`Status: ${task.status}`);
console.log(`Result:`, task.result);

// Send a non-blocking message
const asyncTask = await client.a2a.sendMessage(
  AGENT_ID,
  "Generate a monthly report",
  { blocking: false },
);
console.log(`Task created: ${asyncTask.id} (status: ${asyncTask.status})`);

// Poll for completion
const updated = await client.a2a.getTask(asyncTask.id);
console.log(`Task status: ${updated.status}`);

// List tasks
const tasks = await client.a2a.listTasks({ limit: 5 });
for (const t of tasks.data) {
  console.log(`  ${t.id} — ${t.status}`);
}

// Cancel a running task
await client.a2a.cancelTask(asyncTask.id);
