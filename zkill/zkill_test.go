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
	var err error
	if testClient, err = NewClient(DefaultAddress, "EVELib Testing Agent <github.com/vivace-io/evelib>"); err != nil {
		fmt.Printf("Failed to create test client: %v", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
