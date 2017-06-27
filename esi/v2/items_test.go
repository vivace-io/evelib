package esi

import "testing"

func TestItemGroupIDs(t *testing.T) {
	t.Parallel()
	ids, err := testClient.ItemGroupIDs()
	if err != nil {
		t.Errorf("failed to retrieve item group IDs: %v", err)
	}
	if !(len(ids) >= 1293) {
		t.Errorf("want 1293 or more item group IDs, but got %v", len(ids))
	}
}

func TestItemGroupGet(t *testing.T) {
	t.Parallel()
	// Retrieve group 462 (veldspar)
	group, err := testClient.ItemGroupGet(462)
	if err != nil {
		t.Errorf("failed to retrieve group 462: %v", err)
		return
	}
	if group == nil {
		t.Error("group returned unexpectedly nil")
		return
	}
	if group.GroupID != 462 {
		t.Errorf("ItemGroup.GroupID mismatch -- want 462 but got %v", group.GroupID)
	}
	if group.Name != "Veldspar" {
		t.Errorf("ItemGroup.Name mismatch -- want 'Veldspar' but got '%v'", group.Name)
	}
	if group.CategoryID != 25 {
		t.Errorf("ItemGroup.CategoryID mismatch -- want 25 but got %v", group.CategoryID)
	}
	if group.Published != true {
		t.Error("ItemGroup.Published mismatch -- want true but got false")
	}
	if len(group.Types) != 8 {
		t.Errorf("ItemGroup.Types mismatch - want 8 types but got %v", len(group.Types))
	}
}

func TestItemIDs(t *testing.T) {
	t.Parallel()
	ids, err := testClient.ItemIDs()
	if err != nil {
		t.Errorf("failed to retrieve type IDs: %v", err)
	}
	if !(len(ids) >= 32132) {
		t.Errorf("want 32132 IDs or more but have %v", len(ids))
	}
}

// TestItemGetVeldspar retrieves the Veldspar type from the ESI types endpoint
// and validates it's fields.
func TestItemGet(t *testing.T) {
	t.Parallel()
	var err error
	var veldspar *Item
	if veldspar, err = testClient.ItemGet(1230); err != nil {
		t.Errorf("failed to retrieve Item Type Veldspar: %v", err)
		return
	}
	if veldspar == nil {
		t.Error("veldspar was returned as nil")
		return
	}
	if veldspar.TypeID != 1230 {
		t.Errorf("Field TypeID mismatch - want '1230' but have '%v'", veldspar.TypeID)
	}
	if veldspar.Name != "Veldspar" {
		t.Errorf("Field Name mismatch - want 'Veldspar' but have '%v'", veldspar.Name)
	}
	// Because checking against the entire text description is a little over the top...
	if len(veldspar.Description) == 0 {
		t.Error("Field Description was empty")
	}
	if veldspar.Published != true {
		t.Error("Field Published mismatch - want 'true' but have 'false'")
	}
	if veldspar.GroupID != 462 {
		t.Errorf("Field GroupID mismatch - want '462' but have '%v'", veldspar.GroupID)
	}
	if veldspar.Radius != 1 {
		t.Errorf("Field Radius mismatch - want '1' but have '%v'", veldspar.Radius)
	}
	if veldspar.Volume != 0.1 {
		t.Errorf("Field Volume mismatch - want '0.1' but have '%.2f'", veldspar.Radius)
	}
}
