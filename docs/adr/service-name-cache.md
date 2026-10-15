# Query service caches service names for 60 seconds

Status: proposed (Chiral gate test, do not merge)

## Decision

The query service caches the list of service names in memory for 60 seconds instead of asking the
storage backend on every UI page load.

## Why

The service-name query is the most frequent request from the UI and scans a large index on
Elasticsearch; a short cache removes most of that load while new services still show up within a minute.

## Rejected

- No cache: repeated full scans for an answer that rarely changes.
- A shared Redis cache: a new dependency for a single list.
