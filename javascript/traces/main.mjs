import { PromptRails } from "@promptrails/sdk";

const client = new PromptRails({
  apiKey: process.env.PROMPTRAILS_API_KEY,
});

// List recent traces
const traces = await client.traces.list({ limit: 5 });
console.log(`Total traces: ${traces.meta.total}`);

for (const trace of traces.data) {
  console.log(`  Trace: ${trace.trace_id}`);
  console.log(`    Span: ${trace.name} (${trace.kind})`);
  console.log(`    Status: ${trace.status}, Duration: ${trace.duration_ms}ms`);
  if (trace.cost) console.log(`    Cost: $${trace.cost.toFixed(4)}`);
  console.log();
}

// Get all spans for a specific trace
if (traces.data.length > 0) {
  const traceId = traces.data[0].trace_id;
  const spans = await client.traces.getByTraceId(traceId);
  console.log(`Trace ${traceId} has ${spans.length} spans:`);
  for (const span of spans) {
    const indent = span.parent_span_id ? "  " : "";
    console.log(
      `  ${indent}${span.name} (${span.kind}) — ${span.duration_ms}ms`,
    );
  }
}
