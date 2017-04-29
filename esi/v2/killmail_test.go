package esi

import "testing"

func TestKillmailGet(t *testing.T) {
	t.Parallel()
	result, err := testClient.KillmailGet(61369468, "342ad3ed800d1552df4b1958bcbfdcc832d16aab")
	if err != nil {
		t.Errorf("failed to retrieve killmail: %v", err)
		t.FailNow()
	}
	if result.KillmailID != 61369468 {
		t.Errorf("bad result - want 61369468 but got %v", result.KillmailID)
	}
}
