import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

const AGENT_ID = process.env.AGENT_ID || "your-agent-id";

// Create a chat session
const session = await client.chat.createSession({
  agent_id: AGENT_ID,
  title: "Support conversation",
  metadata: { source: "example" },
});
console.log(`Session created: ${session.id}`);

// Send messages
const msg1 = await client.chat.sendMessage(session.id, {
  content: "What is PromptRails?",
});
console.log(`Assistant: ${msg1.content}`);

const msg2 = await client.chat.sendMessage(session.id, {
  content: "How do I create an agent?",
});
console.log(`Assistant: ${msg2.content}`);

// List message history
const messages = await client.chat.listMessages(session.id);
for (const msg of messages.data) {
  console.log(`  [${msg.role}] ${msg.content.slice(0, 80)}...`);
}

// List sessions
const sessions = await client.chat.listSessions();
for (const s of sessions.data) {
  console.log(`  Session: ${s.title} (${s.id})`);
}

// Clean up
await client.chat.deleteSession(session.id);
