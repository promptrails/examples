# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails"]
# ///
"""Execution tracing and observability."""

import os
from promptrails import PromptRails

client = PromptRails(api_key=os.environ["PROMPTRAILS_API_KEY"])

# List recent traces
traces = client.traces.list(page=1, limit=5)
print(f"Total traces: {traces.meta.total}")

for trace in traces.data:
    print(f"  Trace: {trace.trace_id}")
    print(f"    Span: {trace.name} ({trace.kind})")
    print(f"    Status: {trace.status}")
    print(f"    Duration: {trace.duration_ms}ms")
    if trace.cost:
        print(f"    Cost: ${trace.cost:.4f}")
    print()

# Get all spans for a specific trace
if traces.data:
    trace_id = traces.data[0].trace_id
    spans = client.traces.get_by_trace_id(trace_id)
    print(f"Trace {trace_id} has {len(spans)} spans:")
    for span in spans:
        indent = "  " if span.parent_span_id else ""
        print(f"  {indent}{span.name} ({span.kind}) — {span.duration_ms}ms")

client.close()
