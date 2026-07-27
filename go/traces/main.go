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
		Limit: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total traces: %d\n", traces.Meta.Total)

	for _, trace := range traces.Data {
		fmt.Printf("  Trace: %s\n", trace.TraceID)
		fmt.Printf("    Span: %s (%s)\n", trace.Name, trace.Kind)
		fmt.Printf("    Status: %s, Duration: %dms\n", trace.Status, trace.DurationMs)
		if trace.Cost != nil && *trace.Cost > 0 {
			fmt.Printf("    Cost: $%.4f\n", *trace.Cost)
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
			if span.ParentID != nil {
				indent = "  "
			}
			fmt.Printf("  %s%s (%s) — %dms\n", indent, span.Name, span.Kind, span.DurationMs)
		}
	}

	// Aggregate stats over a filtered set of traces (cost, tokens, latency).
	summary, err := client.Traces.GetSummary(ctx, &promptrails.TraceFilterParams{
		DateFrom: "2026-01-01",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nSummary since 2026-01-01:\n")
	fmt.Printf("  Traces: %d, Cost: $%.2f\n", summary.TotalTraces, summary.TotalCost)
	fmt.Printf("  Tokens: %d, Errors: %d\n", summary.TotalTokens, summary.ErrorCount)

	// PII-masking report over the same filter set
	pii, err := client.Traces.PIIReport(ctx, &promptrails.TraceFilterParams{
		DateFrom: "2026-01-01",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nPII report: %v\n", pii)
}
