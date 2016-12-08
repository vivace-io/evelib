package zkill

import "testing"

func TestKillmailGet(t *testing.T) {
	client, _ := NewClient(nil)
	kill, err := client.KillmailGet(57373234)
	if err != nil {
		t.Errorf("failed to retrieve kill 57373234 with error: %v", err)
	}
	if kill != nil {
		t.Log(kill)
	}
}
