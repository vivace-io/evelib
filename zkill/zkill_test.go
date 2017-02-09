package zkill

import (
	"os"
	"testing"
)

var testClient *Client

/* TestMain
 * ****************************************************************************
 * Main entry point for package testing, sets up fixture(s) for testing.      *
 * ****************************************************************************/
func TestMain(m *testing.M) {
	opts := DefaultOptions()
	// Disable TLS verification for Continuous Integration testing.
	opts.DisableTLS = true
	opts.UserAgent = "EveLib Testing Agent"
	testClient, _ = NewClient(opts)
	os.Exit(m.Run())
}
