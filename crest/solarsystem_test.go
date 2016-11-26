package crest

import "testing"

/* TestSolarSystems
 * ************************************************************************************
 * EXECUTE: SolarSystems(true/false) to call /solarsystems/ for system information    *
 * ************************************************************************************
 * PASS - If no error is returned, and 8035 systems                                   *
 * FAIL - If either an error is returned or any number of items other than 8035       *
 *        is returned.                                                                *
 * ************************************************************************************/
func TestSolarSystemsGetAll(t *testing.T) {
	t.Parallel()
	// [1] SolarSystems(false)
	systems, err := testClient.SolarSystemsGetAll()
	if err != nil {
		t.Errorf("error returned: %v", err)
	}
	if len(systems) != 8035 {
		t.Errorf("expected 8035 items returned but got %v", len(systems))
	}
}

func TestSolarSystemsGet(t *testing.T) {
	sysID := 30000142
	sysName := "Jita"
	// []GetSolarSystem(sysID, false)
	system, err := testClient.SolarSystemsGet(sysID)
	if err != nil {
		t.Errorf("error returned: %v", err)
	}
	if system == nil {
		t.Error("system was nil")
		t.FailNow()
	}
	if system.ID != sysID {
		t.Errorf("system ID mismatch - expected %v but got %v", sysID, system.ID)
	}
	if system.Name != sysName {
		t.Errorf("system name mismatch - expected %v but got %v", sysID, system.ID)
	}
}
