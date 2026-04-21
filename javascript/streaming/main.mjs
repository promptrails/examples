// Server-Sent Events streaming — live chat turns and executions.
//
// The SDK exposes two entry points, both async generators yielding the
// same discriminated StreamEvent union:
//   client.chat.sendMessageStream(sessionId, { content })
//   client.executions.stream(executionId)
//
// Unknown event types are dropped, so the client is forward-compatible
// with new server event kinds.

import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

const AGENT_ID = process.env.AGENT_ID || "your-agent-id";

// ---- Stream a chat turn ---------------------------------------------

const session = await client.chat.createSession({
  agent_id: AGENT_ID,
  title: "Streaming demo",
});
console.log(`Session: ${session.id}\n`);

let executionId;
const controller = new AbortController();

for await (const event of client.chat.sendMessageStream(
  session.id,
  { content: "What is PromptRails?" },
  { signal: controller.signal },
)) {
  switch (event.type) {
    case "execution":
      executionId = event.executionId;
      console.log(`[execution] ${executionId}`);
      break;
    case "thinking":
      console.log(`[thinking] ${event.content}`);
      break;
    case "tool_start":
      console.log(`[tool_start] ${event.name}`);
      break;
    case "tool_end":
      console.log(`[tool_end] ${event.name} — ${event.summary ?? ""}`);
      break;
    case "content":
      process.stdout.write(event.content);
      break;
    case "done":
      console.log(`\n[done] ${event.tokenUsage?.total_tokens ?? 0} tokens`);
      break;
    case "error":
      throw new Error(event.message);
  }
}

// ---- Re-attach to the execution stream ------------------------------
//
// client.executions.stream lets a second process follow an execution
// that started elsewhere (e.g. agents.execute). The backend replays
// in-flight events and then streams new ones.

if (executionId) {
  console.log(`\nReplaying ${executionId}`);
  for await (const event of client.executions.stream(executionId)) {
    if (event.type === "content") process.stdout.write(event.content);
    if (event.type === "done") break;
  }
}

await client.chat.deleteSession(session.id);
