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

	// List all available media models
	models, err := client.MediaModels.List(ctx, &promptrails.ListMediaModelsParams{
		ListParams: promptrails.ListParams{Limit: 50},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total media models: %d\n", models.Meta.Total)

	for _, m := range models.Data {
		fmt.Printf("  %s (%s)\n", m.ModelID, m.DisplayName)
		fmt.Printf("    Provider: %s, Type: %s, Category: %s\n", m.Provider, m.MediaType, m.Category)
		fmt.Printf("    Price: $%.4f %s\n", m.UnitPrice, m.UnitType)
	}

	// Filter by provider — see which models a specific provider offers
	stabilityModels, err := client.MediaModels.List(ctx, &promptrails.ListMediaModelsParams{
		Provider: "stability",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStability models: %d\n", stabilityModels.Meta.Total)
	for _, m := range stabilityModels.Data {
		fmt.Printf("  %s — %s (%s)\n", m.ModelID, m.DisplayName, m.MediaType)
	}

	// Filter by media type — find all text-to-speech models
	ttsModels, err := client.MediaModels.List(ctx, &promptrails.ListMediaModelsParams{
		MediaType: "tts",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nText-to-Speech models: %d\n", ttsModels.Meta.Total)
	for _, m := range ttsModels.Data {
		fmt.Printf("  %s — %s\n", m.ModelID, m.Provider)
	}

	// Filter by category — find all image-related models
	imageModels, err := client.MediaModels.List(ctx, &promptrails.ListMediaModelsParams{
		Category: "image",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nImage models: %d\n", imageModels.Meta.Total)
	for _, m := range imageModels.Data {
		fmt.Printf("  %s — %s (%s)\n", m.ModelID, m.Provider, m.MediaType)
	}

	// Combine filters — ElevenLabs TTS models only
	elevenlabsTTS, err := client.MediaModels.List(ctx, &promptrails.ListMediaModelsParams{
		Provider:  "elevenlabs",
		MediaType: "tts",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nElevenLabs TTS models: %d\n", elevenlabsTTS.Meta.Total)
	for _, m := range elevenlabsTTS.Data {
		fmt.Printf("  %s — %s\n", m.ModelID, m.DisplayName)
		fmt.Printf("    Price: $%.4f %s\n", m.UnitPrice, m.UnitType)
	}
}
