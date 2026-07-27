package main

import (
	"context"
	"fmt"
	"log"
	"os"

	promptrails "github.com/promptrails/go-sdk"
)

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// List all assets with pagination
	assets, err := client.Assets.List(ctx, &promptrails.ListAssetsParams{
		Limit: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total assets: %d\n", assets.Meta.Total)

	for _, asset := range assets.Data {
		fmt.Printf("  %s — %s\n", asset.ID, asset.MediaType)
		fmt.Printf("    Provider: %s\n", asset.Provider)
		fmt.Printf("    File: %s (%s)\n", asset.FileName, asset.ContentType)
		fmt.Printf("    Size: %d bytes\n", asset.Size)
	}

	// Filter assets by media type (e.g., only images)
	images, err := client.Assets.List(ctx, &promptrails.ListAssetsParams{
		MediaType: "image",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nImage assets: %d\n", images.Meta.Total)
	for _, img := range images.Data {
		fmt.Printf("  %s — %s\n", img.ID, img.FileName)
	}

	// Filter assets by provider
	stabilityAssets, err := client.Assets.List(ctx, &promptrails.ListAssetsParams{
		Provider: "stability",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStability assets: %d\n", stabilityAssets.Meta.Total)

	// Get a specific asset by ID
	if len(assets.Data) > 0 {
		assetID := assets.Data[0].ID
		asset, err := client.Assets.Get(ctx, assetID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nAsset detail:\n")
		fmt.Printf("  ID: %s\n", asset.ID)
		fmt.Printf("  Media type: %s\n", asset.MediaType)
		fmt.Printf("  Provider: %s\n", asset.Provider)
		fmt.Printf("  File: %s\n", asset.FileName)
		fmt.Printf("  Content-Type: %s\n", asset.ContentType)
		fmt.Printf("  Size: %d bytes\n", asset.Size)
		if len(asset.Metadata) > 0 {
			fmt.Printf("  Metadata: %v\n", asset.Metadata)
		}

		// Get a signed URL for downloading the asset
		// Signed URLs are temporary (typically 1 hour) and provide direct access to the file
		signed, err := client.Assets.GetSignedURL(ctx, assetID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n  Signed URL: %s\n", signed.URL)
		fmt.Printf("  Expires at: %s\n", signed.ExpiresAt)

		// Delete an asset (removes from storage and soft-deletes the record)
		// if err := client.Assets.Delete(ctx, assetID); err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("\n  Asset %s deleted.\n", assetID)
	}
}
