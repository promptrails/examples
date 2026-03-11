# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Generate media using the Media Studio API."""

import os
import time

from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# --- Image Generation ---
# Generate an image using the Stability provider
image_result = client.media.generate(
    provider="stability",
    media_type="image_gen",
    model="stable-diffusion-xl-1024-v1-0",
    prompt="A serene mountain landscape at sunset with a calm lake reflection",
    config={
        "width": 1024,
        "height": 1024,
        "steps": 30,
    },
)
print(f"Image generation status: {image_result.status}")

# For synchronous providers the asset URL is returned immediately
if image_result.status == "completed":
    print(f"Asset URL: {image_result.asset_url}")
    print(f"Content type: {image_result.content_type}")

# --- Text-to-Speech ---
# Generate speech audio using ElevenLabs
tts_result = client.media.generate(
    provider="elevenlabs",
    media_type="tts",
    model="eleven_multilingual_v2",
    prompt="Welcome to PromptRails! This is an example of text-to-speech generation.",
    config={
        "voice_id": "21m00Tcm4TlvDq8ikWAM",  # Rachel voice
        "stability": 0.5,
        "similarity_boost": 0.75,
    },
)
print(f"\nTTS generation status: {tts_result.status}")
if tts_result.status == "completed":
    print(f"Audio URL: {tts_result.asset_url}")

# --- Async Job Polling ---
# Some providers (e.g., video generation) return a pending job that needs polling.
# Here we demonstrate the pattern with a video generation request.
video_result = client.media.generate(
    provider="runway",
    media_type="video_gen",
    model="gen3a_turbo",
    prompt="A timelapse of clouds moving over a mountain range",
    config={
        "duration": 4,
    },
)
print(f"\nVideo generation status: {video_result.status}")

# If the job is pending, poll until it completes
if video_result.status == "pending" and video_result.job_id:
    print(f"Job ID: {video_result.job_id}")
    print("Waiting for video generation to complete...")

    # Poll every 5 seconds (video generation can take 30-120 seconds)
    for attempt in range(30):
        time.sleep(5)
        status = client.media.get_job_status(video_result.job_id)
        print(f"  Attempt {attempt + 1}: {status.status}")

        if status.status == "completed":
            print(f"Video URL: {status.asset_url}")
            break
        elif status.status == "failed":
            print(f"Generation failed: {status.metadata}")
            break
    else:
        print("Timed out waiting for video generation")

client.close()
