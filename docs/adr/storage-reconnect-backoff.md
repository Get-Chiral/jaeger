# Storage reconnects use exponential backoff with jitter

Status: proposed (Chiral gate test, do not merge)

## Decision

When a storage backend (Elasticsearch, OpenSearch, Cassandra) becomes unreachable, the collector
reconnects with exponential backoff starting at 500 ms, doubling up to a cap of 30 s, with full jitter.

## Why

Fixed one-second retries from every collector replica hit a recovering cluster at the same moment
and kept it overloaded. Jitter spreads the reconnects out.

## Rejected

- Fixed one-second retries: synchronised load on a recovering cluster.
- Giving up after five attempts: collectors would need a restart after every outage.
