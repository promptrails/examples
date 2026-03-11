import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// List all assets with pagination
const assets = await client.assets.list({ limit: 10 });
console.log(`Total assets: ${assets.meta.total}`);

for (const asset of assets.data) {
  console.log(`  ${asset.id} — ${asset.type}`);
  console.log(`    Provider: ${asset.provider}, Model: ${asset.model}`);
  console.log(`    File: ${asset.file_name} (${asset.content_type})`);
  console.log(`    Size: ${asset.size} bytes`);
}

// Filter assets by type (e.g., only images)
const images = await client.assets.list({ type: "image", limit: 5 });
console.log(`\nImage assets: ${images.meta.total}`);
for (const img of images.data) {
  console.log(`  ${img.id} — ${img.file_name}`);
}

// Filter assets by provider
const stabilityAssets = await client.assets.list({
  provider: "stability",
  limit: 5,
});
console.log(`\nStability assets: ${stabilityAssets.meta.total}`);

// Get a specific asset by ID
if (assets.data.length > 0) {
  const assetId = assets.data[0].id;
  const asset = await client.assets.get(assetId);
  console.log(`\nAsset detail:`);
  console.log(`  ID: ${asset.id}`);
  console.log(`  Type: ${asset.type}`);
  console.log(`  Provider: ${asset.provider}`);
  console.log(`  Model: ${asset.model}`);
  console.log(`  File: ${asset.file_name}`);
  console.log(`  Content-Type: ${asset.content_type}`);
  console.log(`  Size: ${asset.size} bytes`);
  if (asset.prompt) {
    console.log(`  Prompt: ${asset.prompt}`);
  }
  if (asset.cost) {
    console.log(`  Cost: $${asset.cost.toFixed(4)}`);
  }

  // Get a signed URL for downloading the asset
  // Signed URLs are temporary (typically 1 hour) and provide direct access to the file
  const signed = await client.assets.getSignedUrl(assetId);
  console.log(`\n  Signed URL: ${signed.url}`);
  console.log(`  Expires at: ${signed.expires_at}`);

  // Delete an asset (removes from storage and soft-deletes the record)
  // await client.assets.delete(assetId);
  // console.log(`\n  Asset ${assetId} deleted.`);
}
