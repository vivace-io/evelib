package esi

import "testing"

func TestBuildPathKillmails(t *testing.T) {
	t.Parallel()
	result := buildPath("/killmails/61369468/342ad3ed800d1552df4b1958bcbfdcc832d16aab/")
	if result != "https://esi.tech.ccp.is/v1/killmails/61369468/342ad3ed800d1552df4b1958bcbfdcc832d16aab/" {
		t.Errorf("build path returned unexpected result: %v", result)
	}
	// TODO - expand coverage
}
