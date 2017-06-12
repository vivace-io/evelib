package esi

import (
	"math/rand"
	"testing"
)

func TestMargetGroupIDs(t *testing.T) {
	t.Parallel()
	results, err := testClient.MarketGroupIDs()
	if err != nil {
		t.Errorf("failed to retrieve market group IDs: %v", err)
		t.FailNow()
	}
	if len(results) != 2164 {
		t.Errorf("bad result - want results length 2164, got %v", len(results))
	}
}

func TestMarketGroupGet(t *testing.T) {
	t.Parallel()
	ids, err := testClient.MarketGroupIDs()
	if err != nil {
		t.Errorf("failed to retrieve market group IDs: %v", err)
		t.FailNow()
	}

	// Since pulling every market group could timeout on Codeship or the
	// developer's computer, we're just going to take a random pool of 50
	// market groups and retrieve them.
	// TODO
	//		Make duplicates impossible.
	//		Use result, don't ignore it.
	for x := 0; x < 50; x++ {
		i := ids[rand.Intn(len(ids))]
		if _, err = testClient.MarketGroupGet(i); err != nil {
			t.Errorf("failed to retrieve market group %v: %v", i, err)
			t.FailNow()
		}
	}
}

func TestMarketPrices(t *testing.T) {
	t.Parallel()
	prices, err := testClient.MarketPrices()
	if err != nil {
		t.Errorf("failed to retrieve prices: %v", err)
		return
	}
	if len(prices) != 10831 {
		t.Errorf("expected 10831 results but have %v", len(prices))
	}
}
