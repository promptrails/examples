import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// Get workspace cost summary
const summary = await client.costs.getSummary();
console.log(`Workspace costs:`);
console.log(`  Total cost: $${summary.total_cost.toFixed(2)}`);
console.log(`  Total tokens: ${summary.total_tokens}`);
console.log(`  Total executions: ${summary.total_executions}`);

// Filter by date range
const feb = await client.costs.getSummary({
  from: "2026-02-01",
  to: "2026-03-01",
});
console.log(`\nFebruary costs: $${feb.total_cost.toFixed(2)}`);

// Get costs for a specific agent
const AGENT_ID = process.env.AGENT_ID || "your-agent-id";
const agentCosts = await client.costs.getAgentSummary(AGENT_ID);
console.log(`Agent costs: $${agentCosts.total_cost.toFixed(2)}`);
