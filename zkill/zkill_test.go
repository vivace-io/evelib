package zkill

import (
	"fmt"
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
	opts.MaxConn = 5
	opts.RateLimit = 10
	opts.RateBurst = 10
	var err error
	testClient, err = NewClient(opts)
	if err != nil {
		fmt.Printf("Failed to create test client: %v", err)
	}
	os.Exit(m.Run())
}
