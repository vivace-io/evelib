package crest

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := Init("Evelib Testing Agent", TranquilityURI, 150, 400)
	if err != nil {
		log.Printf("Testing setup failed with error: %v", err)
		return
	}
	// Disable TLS verification for Continuous Integration testing.
	TLSEnabled(false)
	code := m.Run()
	fmt.Println()
	log.Printf("Tests Completed with Code %v [%v Queries]", code, queryCount)
	fmt.Println()
	os.Exit(code)
}

func TestCRESTError(t *testing.T) {
	err := fetch("some/endpoint", nil)
	if err == nil {
		t.Error("error returned was unexpectedly nil")
	}
}
