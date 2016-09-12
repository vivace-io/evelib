package crest

import "testing"

func TestGetAlliance(t *testing.T) {
	t.Parallel()
	// Wrong Hole. [99006213]
	// https://crest-tq.eveonline.com/alliances/99006213/
	result, err := testClient.GetAlliance(99006213)
	if err != nil {
		t.Error(err)
	}
	if result.ID != 99006213 {
		t.Errorf("alliance ID mismatch - expected 99006213 but got", result.ID)
	}
	if result.Name != "Wrong Hole." {
		t.Errorf("alliance name mismatch - expected \"Wrong Hole\" but got \"%v\"", result.Name)
	}
	if len(result.Corporations) != result.CorporationsCount {
		t.Errorf("member corporation mismatch - %v member corporations listed, with %v counted for", len(result.Corporations), result.CorporationsCount)
	}
}
