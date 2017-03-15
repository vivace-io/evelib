package crest

import "testing"

func TestTypesGet(t *testing.T) {
	t.Parallel()
	types, err := testClient.InventoryTypesGet()
	if err != nil {
		t.Errorf("request failed: %v", err)
	}

	// Check for duplicates.
	dupCount := 0
	for indexA, x := range types {
		for indexB, y := range types {
			if x.ID == y.ID && indexA != indexB {
				dupCount++
				t.Errorf("duplicate found for type %v", x.ID)
			}
		}
	}
	if dupCount != 0 {
		t.Errorf("found %v duplicates out of %v types", dupCount, len(types))
	}
	// TODO - Another way to check expected items, so tests don't break with new
	// items...
	if len(types) != 31910 {
		t.Errorf("wanted 31910 types but got %v", len(types))
	}
}
