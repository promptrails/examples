import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// List all available media models
const models = await client.mediaModels.list({ limit: 50 });
console.log(`Total media models: ${models.meta.total}`);

for (const m of models.data) {
  console.log(`  ${m.model_id} (${m.display_name})`);
  console.log(
    `    Provider: ${m.provider}, Type: ${m.media_type}, Category: ${m.category}`,
  );
  console.log(`    Price: $${m.unit_price.toFixed(4)} ${m.unit_type}`);
}

// Filter by provider — see which models a specific provider offers
const stabilityModels = await client.mediaModels.list({
  provider: "stability",
});
console.log(`\nStability models: ${stabilityModels.meta.total}`);
for (const m of stabilityModels.data) {
  console.log(`  ${m.model_id} — ${m.display_name} (${m.media_type})`);
}

// Filter by media type — find all text-to-speech models
const ttsModels = await client.mediaModels.list({ media_type: "tts" });
console.log(`\nText-to-Speech models: ${ttsModels.meta.total}`);
for (const m of ttsModels.data) {
  console.log(`  ${m.model_id} — ${m.provider}`);
}

// Filter by category — find all image-related models
const imageModels = await client.mediaModels.list({ category: "image" });
console.log(`\nImage models: ${imageModels.meta.total}`);
for (const m of imageModels.data) {
  console.log(`  ${m.model_id} — ${m.provider} (${m.media_type})`);
}

// Combine filters — ElevenLabs TTS models only
const elevenlabsTts = await client.mediaModels.list({
  provider: "elevenlabs",
  media_type: "tts",
});
console.log(`\nElevenLabs TTS models: ${elevenlabsTts.meta.total}`);
for (const m of elevenlabsTts.data) {
  console.log(`  ${m.model_id} — ${m.display_name}`);
  console.log(`    Price: $${m.unit_price.toFixed(4)} ${m.unit_type}`);
}
