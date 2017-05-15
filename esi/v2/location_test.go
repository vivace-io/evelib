package esi

import (
	"math/rand"
	"testing"
)

func TestConstellationIDs(t *testing.T) {
	t.Parallel()
	results, err := testClient.ConstellationIDs()
	if err != nil {
		t.Errorf("failed to retrieve constellation IDs: %v", err)
		t.FailNow()
	}
	if len(results) != 1120 {
		t.Errorf("bad result - want results length 1120, got %v", len(results))
	}
}

func TestConstellationGet(t *testing.T) {
	t.Parallel()
	ids, err := testClient.ConstellationIDs()
	if err != nil {
		t.Errorf("failed to retrieve constellation IDs: %v", err)
		t.FailNow()
	}

	// Since pulling every Constellation could timeout on Codeship or the
	// developer's computer, we're just going to take a random pool of 50
	// constellations and retrieve them.
	// TODO
	//		Make duplicates impossible.
	//		Use result, don't ignore it.
	for x := 0; x < 50; x++ {
		i := ids[rand.Intn(len(ids))]
		if _, err = testClient.ConstellationGet(i); err != nil {
			t.Errorf("failed to retrieve constellation %v: %v", i, err)
			t.FailNow()
		}
	}
}

func TestSystemIDs(t *testing.T) {
	t.Parallel()
	results, err := testClient.SystemIDs()
	if err != nil {
		t.Errorf("failed to retrieve system IDs: %v", err)
		t.FailNow()
	}
	if len(results) != 8035 {
		t.Errorf("bad result - want results length 8035, got %v", len(results))
	}
}

func TestSystemGet(t *testing.T) {
	t.Parallel()
	// Get Jita's System information.
	result, err := testClient.SystemGet(30000142)
	if err != nil {
		t.Errorf("failed to retrieve Jita(30000142) system information: %v", err)
		t.FailNow()
	}
	if result.ID != 30000142 {
		t.Errorf("incorrect system ID - want 30000142 but got %v", result.ID)
	}
	if result.Name != "Jita" {
		t.Errorf("incorrect system name - want Jita but got %v", result.Name)
	}
	// TODO - check the rest
}
