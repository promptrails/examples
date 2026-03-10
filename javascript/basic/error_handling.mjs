import {
  PromptRails,
  NotFoundError,
  ValidationError,
  RateLimitError,
  UnauthorizedError,
  PromptRailsError,
} from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

try {
  await client.agents.get("nonexistent-id");
} catch (err) {
  if (err instanceof NotFoundError) {
    console.log(`Agent not found: ${err.message}`);
  } else if (err instanceof ValidationError) {
    console.log(`Invalid request: ${err.message}`);
  } else if (err instanceof RateLimitError) {
    console.log("Rate limited — retry after a delay");
  } else if (err instanceof UnauthorizedError) {
    console.log("Invalid API key");
  } else if (err instanceof PromptRailsError) {
    console.log(`API error: ${err.message}`);
  }
}
