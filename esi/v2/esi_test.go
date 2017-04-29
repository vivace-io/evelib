package esi

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
	testClient = NewClient(nil)
	testClient.UserAgent = "EveLib Testing Agent (https://github.com/vivace-io/evelib)"
	os.Exit(m.Run())
}
