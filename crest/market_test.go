package crest

import "testing"

func TestMarketTypesGet(t *testing.T) {
	t.Parallel()
	types, err := testClient.MarketTypesGet()
	if err != nil {
		t.Errorf("request failed: %v", err)
	}
	if len(types) != 11752 {
		t.Errorf("wanted 11752 types but got %v", len(types))
	}
}
