# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Manage generated media assets."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# List all assets with pagination
assets = client.assets.list(page=1, limit=10)
print(f"Total assets: {assets.meta.total}")

for asset in assets.data:
    print(f"  {asset.id} — {asset.type}")
    print(f"    Provider: {asset.provider}, Model: {asset.model}")
    print(f"    File: {asset.file_name} ({asset.mime_type})")
    print(f"    Size: {asset.file_size} bytes")
    print()

# Filter assets by type (e.g., only images)
images = client.assets.list(type="image", limit=5)
print(f"Image assets: {images.meta.total}")
for img in images.data:
    print(f"  {img.id} — {img.file_name}")

# Filter assets by provider
stability_assets = client.assets.list(provider="stability", limit=5)
print(f"\nStability assets: {stability_assets.meta.total}")

# Get a specific asset by ID
if assets.data:
    asset_id = assets.data[0].id
    asset = client.assets.get(asset_id)
    print(f"\nAsset detail:")
    print(f"  ID: {asset.id}")
    print(f"  Type: {asset.type}")
    print(f"  Provider: {asset.provider}")
    print(f"  Model: {asset.model}")
    print(f"  File: {asset.file_name}")
    print(f"  MIME type: {asset.mime_type}")
    print(f"  Size: {asset.file_size} bytes")
    print(f"  Execution: {asset.execution_id}")
    if asset.metadata:
        print(f"  Metadata: {asset.metadata}")

    # Get a signed URL for downloading the asset
    # Signed URLs are temporary (typically 1 hour) and provide direct access to the file
    signed = client.assets.get_signed_url(asset_id)
    print(f"\n  Signed URL: {signed.url}")
    print(f"  Expires at: {signed.expires_at}")

    # Delete an asset (removes from storage and soft-deletes the record)
    # client.assets.delete(asset_id)
    # print(f"\n  Asset {asset_id} deleted.")

client.close()
