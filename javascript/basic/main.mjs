import { PromptRails } from "@promptrails/sdk";

// Basic client setup
const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// List agents to verify connection
const agents = await client.agents.list();
console.log(`Connected! Found ${agents.meta.total} agents.`);

for (const agent of agents.data) {
  console.log(`  - ${agent.name} (${agent.type})`);
}

// Custom configuration
const customClient = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
  baseUrl: "https://api.promptrails.ai",
  timeout: 60000,
  maxRetries: 5,
});
