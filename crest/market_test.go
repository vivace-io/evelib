package crest

import "testing"

func TestMarketTypesGet(t *testing.T) {
	t.Parallel()
	types, err := testClient.MarketTypesGet()
	if err != nil {
		t.Errorf("bad request: %v", err)
	}
	if len(types) != 11782 {
		t.Errorf("wanted 11782 types but got %v", len(types))
	}
}
