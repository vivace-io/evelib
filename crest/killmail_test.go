package crest

import "testing"

/* TestGetttKillMail
 * ***********************************************************************************
 * EXECUTE: GetKillmail(id, hash) to fetch one or more killmails                     *
 * ***********************************************************************************
 * PASS - If no error is returned, result data is as expected                        *
 * FAIL - If either an unexpected error is retured or result data is not as expected *
 * ***********************************************************************************/
func TestGetKillmail(t *testing.T) {
	t.Parallel()
	// Zkillboard - https://zkillboard.com/kill/55214047/
	// CREST - https://crest-tq.eveonline.com/killmails/55214047/4d78d27888c789fef959f59c6e417ed91ad4c502/
	id := 55214047
	hash := "4d78d27888c789fef959f59c6e417ed91ad4c502"
	kill, err := testClient.KillmailGet(id, hash)
	if err != nil {
		t.Error("error retrieving killmail: %v", err)
	}
	if kill.KillID != id {
		t.Error("kill id mismatch - want %v but got %v", id, kill.KillID)
	}
	if kill.Victim.Character.ID != 325487984 {
		t.Error("victim character id mismatch - want 325487984 but got %v", kill.Victim.Character.ID)
	}
	if len(kill.Attackers) != 4 {
		t.Errorf("unexpected attacker count - wanted 4 but got %v", len(kill.Attackers))
	}
}
