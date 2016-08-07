package crest

import "testing"

/* TestRegions
 * *********************************************************************************
 * EXECUTE: Regions(true|false) to call /regions/ for regions information          *
 * *********************************************************************************
 * PASS - If no error is returned, result data is as expected                      *
 * FAIL - If either an error is returned or any number of items than other 100     *
 *        is returned.                                                             *
 * *********************************************************************************/
func TestRegions(t *testing.T) {
	t.Parallel()
	// [1] Regions(false)
	regions, err := Regions(false)
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
	// [2] Regions(true)
	regions, err = Regions(true)
	if err != nil {
		t.Errorf("regions retrieval failed - error returned: %v", err)
	}
	if len(regions) != 100 {
		t.Errorf("expected 100 regions returned but got %v", len(regions))
	}
	for _, r := range regions {
		if r.ID < 10000001 || r.ID > 11000033 {
			t.Errorf("region ID of value %v is outside expected domain [10000001,11000033]", r.ID)
		}
		if r.Name == "" {
			t.Errorf("region of ID %v has empty name", r.ID)
		}
		// TODO
		// some regions have an empty description, and this needs to be accounted for
		// in the future. Moving forward, just log to stdout when -test.v flag is set.
		if r.Description == "" {
			//t.Logf("[WARN] Region %v[%v] has an empty description.", r.Name, r.ID)
		}
		if r.Href == "" {
			t.Errorf("region of ID %v has empty href", r.Href)
		}
		if len(r.Constellations) == 0 {
			t.Errorf("region of ID %v has no constellations", r.ID)
		}
	}
}
