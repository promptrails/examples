import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// List available MCP templates
const templates = await client.mcpTemplates.list();
for (const t of templates.data) {
  console.log(`  ${t.name} (${t.slug}) — ${t.category}`);
}

// Install a tool from template
const tool = await client.mcpTemplates.install("brave-search", {
  name: "Web Search",
  parameters: { api_key: "your-brave-api-key" },
});
console.log(`Installed tool: ${tool.id}`);

// List installed tools
const tools = await client.mcpTools.list();
for (const t of tools.data) {
  console.log(`  ${t.name} (${t.type}) — ${t.status}`);
}

// Discover available tool functions
if (tools.data.length > 0) {
  const discovered = await client.mcpTools.discoverTools(tools.data[0].id);
  for (const dt of discovered.tools) {
    console.log(`  Tool: ${dt.name} — ${dt.description}`);
  }
}

// Call a tool
const result = await client.mcpTools.callTool(tool.id, {
  tool_name: "search",
  arguments: { query: "PromptRails AI platform" },
});
console.log(`Result:`, result.content);

// Clean up
await client.mcpTools.delete(tool.id);
