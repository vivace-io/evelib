package zkill

import (
	"testing"
	"time"
)

func TestKillmailGet(t *testing.T) {
	t.Parallel()
	kill, err := testClient.KillmailGet(59972778)
	if err != nil {
		t.Errorf("expected nil error but got: %v", err)
		return
	}
	if kill.KillID != 59972778 {
		t.Errorf("kill ID mismatch -- want 59972778 but have %v", kill.KillID)
	}
}

func TestKillmailGetError(t *testing.T) {
	t.Parallel()
	_, err := testClient.KillmailGet(0)
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestHistorical(t *testing.T) {
	t.Parallel()
	result, err := testClient.Historical(time.Now())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(result) == 0 {
		t.Error("result was unexpectedly empty")
	}
}
