// Copyright 2026 Parth Chandak and contributors. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"testing"
)

func TestListItemsUnwrapsPaginatedEnvelope(t *testing.T) {
	items, ok := listItems(json.RawMessage(`{"page":1,"page_count":2,"workouts":[{"id":"a"},{"id":"b"}]}`))
	if !ok || len(items) != 2 || items[1]["id"] != "b" {
		t.Fatalf("envelope: got %v, %v", items, ok)
	}
}

func TestListItemsBareArray(t *testing.T) {
	items, ok := listItems(json.RawMessage(`[{"id":1}]`))
	if !ok || len(items) != 1 {
		t.Fatalf("array: got %v, %v", items, ok)
	}
}

func TestListItemsRejectsSingleObjects(t *testing.T) {
	for _, body := range []string{
		`{"workout_count":3}`,
		`{"data":{"id":"u1"}}`,
		`{"a":[{"id":1}],"b":[{"id":2}]}`,
	} {
		if _, ok := listItems(json.RawMessage(body)); ok {
			t.Fatalf("%s: expected not a list", body)
		}
	}
}

func TestPrioritizeHeadersIsStableAcrossRuns(t *testing.T) {
	item := map[string]any{"zeta": 1, "alpha": 2, "mid": 3, "beta": 4, "id": "x", "title": "t"}
	want := prioritizeHeaders(item)
	if want[0] != "id" || want[1] != "title" {
		t.Fatalf("identity fields first: got %v", want)
	}
	for i := 0; i < 50; i++ {
		got := prioritizeHeaders(item)
		for j := range want {
			if got[j] != want[j] {
				t.Fatalf("run %d: order changed: %v vs %v", i, got, want)
			}
		}
	}
}

func TestWriteThroughUnwrapsHevyPageCountEnvelope(t *testing.T) {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal([]byte(`{"page":1,"page_count":3,"workouts":[{"id":"a"},{"id":"b"}]}`), &envelope); err != nil {
		t.Fatal(err)
	}
	items, ok := extractWriteThroughListItems(envelope)
	if !ok || len(items) != 2 {
		t.Fatalf("expected 2 workout rows, got %d (ok=%v)", len(items), ok)
	}
}
