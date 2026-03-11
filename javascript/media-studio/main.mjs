import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// --- Image Generation ---
// Generate an image using the Stability provider
const imageResult = await client.media.generate({
  provider: "stability",
  media_type: "image_gen",
  model: "stable-diffusion-xl-1024-v1-0",
  prompt: "A serene mountain landscape at sunset with a calm lake reflection",
  config: {
    width: 1024,
    height: 1024,
    steps: 30,
  },
});
console.log(`Image generation status: ${imageResult.status}`);

// For synchronous providers the asset URL is returned immediately
if (imageResult.status === "completed") {
  console.log(`Asset URL: ${imageResult.asset_url}`);
  console.log(`Content type: ${imageResult.content_type}`);
}

// --- Text-to-Speech ---
// Generate speech audio using ElevenLabs
const ttsResult = await client.media.generate({
  provider: "elevenlabs",
  media_type: "tts",
  model: "eleven_multilingual_v2",
  prompt:
    "Welcome to PromptRails! This is an example of text-to-speech generation.",
  config: {
    voice_id: "21m00Tcm4TlvDq8ikWAM", // Rachel voice
    stability: 0.5,
    similarity_boost: 0.75,
  },
});
console.log(`\nTTS generation status: ${ttsResult.status}`);
if (ttsResult.status === "completed") {
  console.log(`Audio URL: ${ttsResult.asset_url}`);
}

// --- Async Job Polling ---
// Some providers (e.g., video generation) return a pending job that needs polling.
// Here we demonstrate the pattern with a video generation request.
const videoResult = await client.media.generate({
  provider: "runway",
  media_type: "video_gen",
  model: "gen3a_turbo",
  prompt: "A timelapse of clouds moving over a mountain range",
  config: {
    duration: 4,
  },
});
console.log(`\nVideo generation status: ${videoResult.status}`);

// If the job is pending, poll until it completes
if (videoResult.status === "pending" && videoResult.job_id) {
  console.log(`Job ID: ${videoResult.job_id}`);
  console.log("Waiting for video generation to complete...");

  // Poll every 5 seconds (video generation can take 30-120 seconds)
  for (let attempt = 1; attempt <= 30; attempt++) {
    await new Promise((r) => setTimeout(r, 5000));
    const status = await client.media.getJobStatus(videoResult.job_id);
    console.log(`  Attempt ${attempt}: ${status.status}`);

    if (status.status === "completed") {
      console.log(`Video URL: ${status.asset_url}`);
      break;
    } else if (status.status === "failed") {
      console.log(`Generation failed:`, status.metadata);
      break;
    }
  }
}
