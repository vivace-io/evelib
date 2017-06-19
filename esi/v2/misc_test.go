package esi

import (
	"encoding/json"
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

func TestESITimeUnmarshalJSON(t *testing.T) {
	t.Parallel()
	type Model struct {
		Time ESITime `json:"time"`
	}
	data := []byte(`{"time":"2015-05-01"}`)
	var m *Model
	if err := json.Unmarshal(data, &m); err != nil {
		t.Errorf("unexpected marshal error: %v", err)
		return
	}
	if m.Time.Year() != 2015 {
		t.Errorf("Year mismatch - want 2015 have %v", m.Time.Year)
	}
	if m.Time.Month().String() != "May" {
		t.Errorf("Month mismatch - want 'May' have '%v'", m.Time.Month().String())
	}
	if m.Time.Day() != 1 {
		t.Errorf("Day mismatch - want 1 but have %v", m.Time.Day())
	}
}
