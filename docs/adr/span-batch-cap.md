# Span batches are capped at 1,000 spans

Status: proposed (Chiral gate test, do not merge)

## Decision

The collector writes spans to storage in batches of at most 1,000 spans.

## Why

Larger batches made Elasticsearch bulk requests exceed the default 100 MB request limit during traffic spikes.

## Rejected

- Raising the Elasticsearch request limit: it has to be changed on every customer cluster.
