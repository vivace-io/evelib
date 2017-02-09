package zkill

import (
	"testing"
	"time"
)

func TestKillmailGet(t *testing.T) {
	t.Parallel()
	kill, err := testClient.KillmailGet(57373234)
	if err != nil {
		t.Errorf("failed to retrieve kill 57373234 with error: %v", err)
	}
	if kill != nil {
		t.Log(kill)
	}
}

func TestKillmailGetError(t *testing.T) {
	t.Parallel()
	kills, err := testClient.KillmailGet(0, 57373234)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if len(kills) != 1 {
		t.Errorf("expected one killmail but got %v", len(kills))
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
