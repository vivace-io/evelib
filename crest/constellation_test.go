package crest

import "testing"

func TestConstellationsGetAll(t *testing.T) {
	result, err := testClient.ConstellationsGetAll()
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1120 {
		t.Errorf("constellation count mismatch - want 1120 but got %v", len(result))
	}
}
