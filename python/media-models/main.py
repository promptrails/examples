# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Browse available media models."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# List all available media models
models = client.media_models.list(page=1, limit=50)
print(f"Total media models: {models.meta.total}")

for m in models.data:
    print(f"  {m.model_id} ({m.display_name})")
    print(f"    Provider: {m.provider}, Type: {m.media_type}, Category: {m.category}")
    print(f"    Price: ${m.unit_price:.4f} {m.unit_type}")
    print()

# Filter by provider — see which models a specific provider offers
stability_models = client.media_models.list(provider="stability")
print(f"Stability models: {stability_models.meta.total}")
for m in stability_models.data:
    print(f"  {m.model_id} — {m.display_name} ({m.media_type})")

# Filter by media type — find all text-to-speech models
tts_models = client.media_models.list(media_type="tts")
print(f"\nText-to-Speech models: {tts_models.meta.total}")
for m in tts_models.data:
    print(f"  {m.model_id} — {m.provider}")

# Filter by category — find all image-related models
image_models = client.media_models.list(category="image")
print(f"\nImage models: {image_models.meta.total}")
for m in image_models.data:
    print(f"  {m.model_id} — {m.provider} ({m.media_type})")

# Combine filters — ElevenLabs TTS models only
elevenlabs_tts = client.media_models.list(provider="elevenlabs", media_type="tts")
print(f"\nElevenLabs TTS models: {elevenlabs_tts.meta.total}")
for m in elevenlabs_tts.data:
    print(f"  {m.model_id} — {m.display_name}")
    print(f"    Price: ${m.unit_price:.4f} {m.unit_type}")

client.close()
