import { PromptRails } from "@promptrails/sdk";

// The dedicated `costs` resource is gone in v2. Cost, token and latency
// aggregates now come from traces.getSummary(filters), which accepts the same
// filters as traces.list plus date_from / date_to, status, level, model_name,
// agent_id, session_id and execution_id.

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// Workspace-wide summary
const summary = await client.traces.getSummary();
console.log("Workspace usage:");
console.log(`  Total cost:    $${summary.total_cost.toFixed(2)}`);
console.log(`  Total tokens:  ${summary.total_tokens}`);
console.log(`  Total traces:  ${summary.total_traces}`);
console.log(`  Errors:        ${summary.error_count}`);
console.log(`  Avg duration:  ${summary.avg_duration_ms.toFixed(0)}ms`);
console.log(`  Unique models: ${summary.unique_models}`);

// Filter by date range
const feb = await client.traces.getSummary({
  date_from: "2026-02-01",
  date_to: "2026-03-01",
});
console.log(
  `\nFebruary cost: $${feb.total_cost.toFixed(2)} over ${feb.total_traces} traces`,
);

// Scope to a single agent
const AGENT_ID = process.env.AGENT_ID || "your-agent-id";
const agentSummary = await client.traces.getSummary({ agent_id: AGENT_ID });
console.log(`Agent cost: $${agentSummary.total_cost.toFixed(2)}`);
