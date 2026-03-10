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

	// List recent traces
	traces, err := client.Traces.List(ctx, &promptrails.ListTracesParams{
		ListParams: promptrails.ListParams{Limit: 5},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total traces: %d\n", traces.Meta.Total)

	for _, trace := range traces.Data {
		fmt.Printf("  Trace: %s\n", trace.TraceID)
		fmt.Printf("    Span: %s (%s)\n", trace.Name, trace.Kind)
		fmt.Printf("    Status: %s, Duration: %dms\n", trace.Status, trace.DurationMs)
		if trace.Cost > 0 {
			fmt.Printf("    Cost: $%.4f\n", trace.Cost)
		}
		fmt.Println()
	}

	// Get all spans for a specific trace
	if len(traces.Data) > 0 {
		traceID := traces.Data[0].TraceID
		spans, err := client.Traces.GetByTraceID(ctx, traceID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Trace %s has %d spans:\n", traceID, len(spans))
		for _, span := range spans {
			indent := ""
			if span.ParentSpanID != "" {
				indent = "  "
			}
			fmt.Printf("  %s%s (%s) — %dms\n", indent, span.Name, span.Kind, span.DurationMs)
		}
	}
}
