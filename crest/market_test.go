package crest

import "testing"

func TestMarketTypesGet(t *testing.T) {
	t.Parallel()
	types, err := testClient.MarketTypesGet()
	if err != nil {
		t.Errorf("request failed: %v", err)
	}
	if len(types) != 11748 {
		t.Errorf("wanted 11748 types but got %v", len(types))
	}
}
