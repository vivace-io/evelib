package crest

import (
	"os"
	"testing"
)

var testClient *Client

func TestMain(m *testing.M) {
	testClient = NewClient("EveLib Testing Agent", TranquilityURI, 150, 400)

	// Disable TLS verification for Continuous Integration testing.
	testClient.TLS(false)
	os.Exit(m.Run())
}

func TestCRESTError(t *testing.T) {
	err := testClient.get("some/endpoint", nil)
	if err == nil {
		t.Error("error returned was unexpectedly nil")
	}
}
