package crest

import "testing"

/* TestConstellationsGetAll
 * ************************************************************************************
 * EXECUTE: ConstellationsGetAll to call /constellations/ for a list of consellations.*
 * ************************************************************************************
 * PASS - If no error is returned, and result is of length 1120.                      *
 * FAIL - If either an error is returned or result is not of length 1120.             *
 * ************************************************************************************/
func TestConstellationsGetAll(t *testing.T) {
	result, err := testClient.ConstellationsGetAll()
	if err != nil {
		t.Errorf("bad request: %v", err)
	}
	if len(result) != 1120 {
		t.Errorf("constellation count mismatch - want 1120 but got %v", len(result))
	}
}
