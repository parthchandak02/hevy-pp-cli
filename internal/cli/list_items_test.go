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
