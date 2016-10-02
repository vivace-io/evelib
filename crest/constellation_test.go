package crest

import "testing"

func TestConstellations(t *testing.T) {
	result, err := testClient.AllConstellations()
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1120 {
		t.Errorf("constellation count mismatch - want 1120 but got %v", len(result))
	}
}
