import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

const AGENT_ID = process.env.AGENT_ID || "your-agent-id";

// Create a webhook trigger
const trigger = await client.webhookTriggers.create({
  name: "GitHub Push Webhook",
  agent_id: AGENT_ID,
  generate_secret: true,
});
console.log(`Trigger created: ${trigger.id}`);
console.log(`Webhook URL: ${trigger.url}`);
console.log(`Secret: ${trigger.secret}`);

// List triggers
const triggers = await client.webhookTriggers.list();
for (const t of triggers.data) {
  console.log(`  ${t.name} — active: ${t.is_active}`);
}

// Update trigger
await client.webhookTriggers.update(trigger.id, { name: "Updated Webhook" });

// Delete trigger
await client.webhookTriggers.delete(trigger.id);
console.log("Trigger deleted.");
