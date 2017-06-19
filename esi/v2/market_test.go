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
	if len(prices) != 10859 {
		t.Errorf("expected 10859 results but have %v", len(prices))
	}
}

func TestMarketRegionHistory(t *testing.T) {
	t.Parallel()
	history, err := testClient.MarketRegionHistoryGet(10000002, 1230)
	if err != nil {
		t.Errorf("failed to retrieve market region history: %v", err)
	}
	if !(len(history) >= 365) {
		if len(history) == 0 {
			t.Error("empty response (no records)")
		} else {
			t.Errorf("expecting at least 365 records but got %v (this test may be out of date)", len(history))
		}
		return
	}
	for i, h := range history {
		if h.Date.IsZero() {
			t.Errorf("history[%v].Data was zero", i)
		}
		if h.OrderCount <= 0 {
			t.Errorf("history[%v].OrderCount <= 0", i)
		}
		if h.Volume <= 0 {
			t.Errorf("history[%v].Volume <= 0", i)
		}
		if h.Highest <= 0 {
			t.Errorf("history[%v].Highest <= 0", i)
		}
		if h.Average <= 0 {
			t.Errorf("history[%v].Average <= 0", i)
		}
		if h.Lowest <= 0 {
			t.Errorf("history[%v].Lowest <= 0", i)
		}
	}
}
