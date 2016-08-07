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
func TestSolarSystems(t *testing.T) {
	t.Parallel()
	start := queryCount
	// [1] SolarSystems(false)
	systems, err := SolarSystems(false)
	if err != nil {
		t.Errorf("error returned: %v", err)
	}
	if len(systems) != 8035 {
		t.Errorf("expected 8035 items returned but got %v", len(systems))
	}

	// [2] SolarSystems(true)
	/* PROCESS TOO LONG!
	systems, err = SolarSystems(true)
	if err != nil {
		t.Errorf("error returned: %v", err)
	}
	if len(systems) != 8035 {
		t.Errorf("expected 8035 items returned but got %v", len(systems))
	}
	for _, s := range systems {
		if s.Position == nil {
			t.Errorf("position of system %v[%v] is unexpectedly nil", s.Name, s.ID)
		}
		if len(s.Planets) == 0 {
			t.Errorf("no planets returned in system %v[%v]", s.Name, s.ID)
		}
	}
	*/
	t.Logf("[NOTIFY] TestSolarSystems - Executed %v Queries.", queryCount-start)

}

func TestGetSolarSystem(t *testing.T) {
	sysID := 30000142
	sysName := "Jita"
	// []GetSolarSystem(sysID, false)
	system, err := GetSolarSystem(sysID, false)
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
