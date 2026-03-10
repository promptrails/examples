import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// List recent executions
const executions = await client.executions.list({ limit: 10 });
console.log(`Total executions: ${executions.meta.total}`);

for (const ex of executions.data) {
  console.log(`  ${ex.id} — ${ex.status}`);
  console.log(`    Agent: ${ex.agent_id}`);
  console.log(`    Duration: ${ex.duration_ms}ms, Cost: $${ex.cost.toFixed(4)}`);
}

// Filter by status
const failed = await client.executions.list({ status: "failed", limit: 5 });
console.log(`\nFailed executions: ${failed.meta.total}`);
for (const ex of failed.data) {
  console.log(`  ${ex.id}: ${ex.error}`);
}

// Get execution details
if (executions.data.length > 0) {
  const detail = await client.executions.get(executions.data[0].id);
  console.log(`\nExecution detail:`);
  console.log(`  Input:`, detail.input);
  console.log(`  Output:`, detail.output);
  console.log(`  Token usage:`, detail.token_usage);
}
