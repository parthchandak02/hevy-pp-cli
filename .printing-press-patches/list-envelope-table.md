# Table output for paginated list envelopes

Hevy wraps every list endpoint in an envelope (`{"page":1,"page_count":N,"workouts":[...]}`).
The generated human-output path only rendered a table for a bare JSON array, so in a terminal
every list command printed pretty JSON and a "0 results" provenance line.

- `internal/cli/helpers.go`: `listItems()` returns the rows of a bare array or of an object whose
  single array-of-objects field holds them; `printOutput()` uses it for the table branch.
- Provenance counts in the endpoint and promoted commands use `listItems()` instead of
  unmarshalling into `[]json.RawMessage`.
- `internal/cli/list_items_test.go` covers both shapes and the single-object cases.

Reprint guard: keep envelope-aware list detection in the human table path.

Follow-up: `prioritizeFields()` broke same-tier ties by map iteration order, so table and card
columns shuffled between runs. Ties now sort by field name (`TestPrioritizeHeadersIsStableAcrossRuns`).

Follow-up: `page_count` / `pageCount` added to `listEnvelopeMetadataKeys` in `data_source.go`, so the
write-through cache unwraps Hevy's envelope into rows instead of warning that one ID-less item was not cached.
