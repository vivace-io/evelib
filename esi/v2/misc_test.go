package esi

import (
	"testing"
	"time"
)

func TestStatusGet(t *testing.T) {
	t.Parallel()
	status, err := testClient.StatusGet()
	if err != nil {
		t.Errorf("failed to retrieve status: %v", err)
		return
	}
	if status.Players == 0 {
		t.Error("Status.Players was returned as zero")
	}
	if status.ServerVersion == "" {
		t.Error("Status.ServerVersion was returned empty")
	}
	if status.StartTime.IsZero() {
		t.Error("Status.StartTime was returned as zero")
	} else if time.Since(status.StartTime) > 24*time.Hour {
		// Because the EVE cluster is rebooted daily, StartTime can be assumed to
		// never be older than 24 hours. There are times where its possible, but for
		// or purposes we can safely assume it is not.
		t.Errorf("Status.StartTime was older than 24 hours, returned as: %v", status.StartTime)
	}
}
