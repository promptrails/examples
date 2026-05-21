import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

const AGENT_ID = process.env.AGENT_ID || "your-agent-id";

// 1. Generic webhook trigger (default source)
const trigger = await client.agentTriggers.create({
  name: "GitHub Push Webhook",
  agent_id: AGENT_ID,
  generate_secret: true,
});
console.log(`Generic trigger created: ${trigger.id}`);
console.log(`  Token:  ${trigger.token}`);
console.log(`  Secret: ${trigger.secret}`);

// 2. Slack trigger — point at workspace credentials for signing secret + bot token
const slackTrigger = await client.agentTriggers.create({
  name: "Slack #incidents",
  agent_id: AGENT_ID,
  source: "slack",
  source_config: {
    signing_secret_credential_id: "<credential id with the Slack signing secret>",
    bot_token_credential_id: "<credential id with the bot token>",
  },
});
console.log(`Slack trigger created: ${slackTrigger.id}`);

// 3. Schedule trigger — fires on a cron expression
const scheduleTrigger = await client.agentTriggers.create({
  name: "Daily digest at 9am",
  agent_id: AGENT_ID,
  source: "schedule",
  source_config: {
    cron_expression: "0 9 * * *",
    next_run_at: "2026-05-22T09:00:00Z",
  },
});
console.log(`Schedule trigger created: ${scheduleTrigger.id}`);

// List + update + delete still work the same way
const triggers = await client.agentTriggers.list();
for (const t of triggers.data) {
  console.log(`  ${t.name} (${t.source}) — active: ${t.is_active}`);
}

await client.agentTriggers.update(trigger.id, { name: "GitHub Push Webhook (renamed)" });
await client.agentTriggers.delete(trigger.id);
await client.agentTriggers.delete(slackTrigger.id);
await client.agentTriggers.delete(scheduleTrigger.id);
console.log("Triggers cleaned up.");
