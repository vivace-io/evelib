package crest

import "testing"

/* TestRegionsGetAll
 * *********************************************************************************
 * EXECUTE: RegionsGetAll() to call /regions/ for a list of regions.               *
 * *********************************************************************************
 * PASS - If no error is returned, result data/length is as expected.              *
 * FAIL - If either an error is returned or any number of items than other 100     *
 *        is returned.                                                             *
 * *********************************************************************************/
func TestRegionsGetAll(t *testing.T) {
	t.Parallel()
	regions, err := testClient.RegionsGetAll()
	if err != nil {
		t.Errorf("Regions(false) failed - error returned: %v", err)
	}
	if len(regions) != 100 {
		t.Errorf("Regions(false) failed - expected 100 items returned but got %v", len(regions))
	}
	for _, r := range regions {
		if r.ID < 10000001 || r.ID > 11000033 {
			t.Errorf("region ID of value %v is outside expected domain [10000001,11000033]", r.ID)
		}
		if r.Name == "" {
			t.Errorf("region of ID %v has empty name", r.ID)
		}
		if r.Href == "" {
			t.Errorf("region of ID %v has empty href", r.Href)
		}
	}
}
