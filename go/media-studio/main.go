package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	promptrails "github.com/promptrails/go-sdk"
)

func main() {
	client := promptrails.NewClient(os.Getenv("PROMPTRAILS_API_KEY"))
	ctx := context.Background()

	// --- Image Generation ---
	// Generate an image using the Stability provider
	imageResult, err := client.Media.Generate(ctx, &promptrails.MediaGenerateParams{
		Provider:  "stability",
		MediaType: "image_gen",
		Model:     "stable-diffusion-xl-1024-v1-0",
		Prompt:    "A serene mountain landscape at sunset with a calm lake reflection",
		Config: map[string]any{
			"width":  1024,
			"height": 1024,
			"steps":  30,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Image generation status: %s\n", imageResult.Status)

	// For synchronous providers the asset URL is returned immediately
	if imageResult.Status == "completed" {
		fmt.Printf("Asset URL: %s\n", imageResult.AssetURL)
		fmt.Printf("Content type: %s\n", imageResult.ContentType)
	}

	// --- Text-to-Speech ---
	// Generate speech audio using ElevenLabs
	ttsResult, err := client.Media.Generate(ctx, &promptrails.MediaGenerateParams{
		Provider:  "elevenlabs",
		MediaType: "tts",
		Model:     "eleven_multilingual_v2",
		Prompt:    "Welcome to PromptRails! This is an example of text-to-speech generation.",
		Config: map[string]any{
			"voice_id":         "21m00Tcm4TlvDq8ikWAM", // Rachel voice
			"stability":        0.5,
			"similarity_boost": 0.75,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nTTS generation status: %s\n", ttsResult.Status)
	if ttsResult.Status == "completed" {
		fmt.Printf("Audio URL: %s\n", ttsResult.AssetURL)
	}

	// --- Async Job Polling ---
	// Some providers (e.g., video generation) return a pending job that needs polling.
	// Here we demonstrate the pattern with a video generation request.
	videoResult, err := client.Media.Generate(ctx, &promptrails.MediaGenerateParams{
		Provider:  "runway",
		MediaType: "video_gen",
		Model:     "gen3a_turbo",
		Prompt:    "A timelapse of clouds moving over a mountain range",
		Config: map[string]any{
			"duration": 4,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nVideo generation status: %s\n", videoResult.Status)

	// If the job is pending, poll until it completes
	if videoResult.Status == "pending" && videoResult.JobID != "" {
		fmt.Printf("Job ID: %s\n", videoResult.JobID)
		fmt.Println("Waiting for video generation to complete...")

		// Poll every 5 seconds (video generation can take 30-120 seconds)
		for attempt := 1; attempt <= 30; attempt++ {
			time.Sleep(5 * time.Second)
			status, err := client.Media.GetJobStatus(ctx, videoResult.JobID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("  Attempt %d: %s\n", attempt, status.Status)

			if status.Status == "completed" {
				fmt.Printf("Video URL: %s\n", status.AssetURL)
				break
			} else if status.Status == "failed" {
				fmt.Printf("Generation failed: %v\n", status.Metadata)
				break
			}
		}
	}
}
