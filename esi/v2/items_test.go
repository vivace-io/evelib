package esi

import "testing"

func TestItemIDs(t *testing.T) {
	t.Parallel()
	ids, err := testClient.ItemIDs()
	if err != nil {
		t.Errorf("failed to retrieve type IDs: %v", err)
	}
	if len(ids) != 32013 {
		t.Errorf("want 32013 IDs but have %v", len(ids))
	}
}
