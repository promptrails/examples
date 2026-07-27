import { PromptRails } from "@promptrails/sdk";

// Approvals are execution-scoped in v2. A run that hits an approval-gated tool
// or sub-agent parks at status "waiting_approval"; resume it with
// executions.approve / executions.deny. executions.tree returns the full
// parent -> child tree, executions.cancel requests cooperative cancellation.

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

// Get execution details, then walk its full child tree
if (executions.data.length > 0) {
  const detail = await client.executions.get(executions.data[0].id);
  console.log(`\nExecution detail:`);
  console.log(`  Input:`, detail.input);
  console.log(`  Output:`, detail.output);

  const tree = await client.executions.tree(detail.id);
  console.log(`  Child executions: ${tree.children?.length ?? 0}`);
}

// Approval inbox — runs parked at waiting_approval
const inbox = await client.executions.approvalInbox({ limit: 10 });
console.log(`\nAwaiting approval: ${inbox.meta.total}`);
for (const ex of inbox.data) {
  console.log(`  ${ex.id} — ${ex.status}`);
  // Approve to resume, or deny to resume with a denial:
  // await client.executions.approve(ex.id, { reason: "Looks safe" });
  // await client.executions.deny(ex.id, { reason: "Not allowed" });
}

// Cancel a still-running execution
// await client.executions.cancel("exec-id");
