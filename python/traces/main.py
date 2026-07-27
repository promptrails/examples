# /// script
# requires-python = ">=3.9"
# dependencies = ["promptrails>=0.9.0"]
# ///
#
# NOTE (API v2): promptrails 0.9.0 is not yet on PyPI. Until it publishes, run
# against the local sibling SDK from inside this folder, e.g.:
#   uv run --with-editable ../../../python-sdk main.py
# The pinned spec above reconciles automatically once 0.9.0 ships.
"""Execution tracing and observability — API v2."""

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

# Get all spans for a specific trace
if traces.data:
    trace_id = traces.data[0].trace_id
    spans = client.traces.get_by_trace_id(trace_id)
    print(f"Trace {trace_id} has {len(spans)} spans:")
    for span in spans:
        indent = "  " if span.parent_span_id else ""
        print(f"  {indent}{span.name} ({span.kind}) — {span.duration_ms}ms")

# Aggregate stats over a filtered set of traces (cost, tokens, latency, errors)
summary = client.traces.get_summary(date_from="2026-01-01")
print("\nSummary since 2026-01-01:")
print(f"  Traces: {summary.total_traces}, Cost: ${summary.total_cost:.2f}")
print(f"  Tokens: {summary.total_tokens}, Errors: {summary.error_count}")

# PII-masking report over the same filter set
pii = client.traces.pii_report(date_from="2026-01-01")
print(f"\nPII report: {pii}")

client.close()
