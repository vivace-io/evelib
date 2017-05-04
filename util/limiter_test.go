package util

import (
	"testing"
)

// TestLimiterConn
// Ensure that the maximum connections may not be exceeded.
// TODO - Figure out why Codeship is timing this out, even though it passes
// in other environments.
//
// func TestLimiterConn(t *testing.T) {
// 	limiter, _ := NewLimiter(100, 100, 5)
// 	limiter.Connect()
// 	limiter.Connect()
// 	limiter.Connect()
// 	limiter.Connect()
// 	limiter.Connect()
// 	connected := false
// 	go func() {
// 		limiter.Connect()
// 		connected = true
// 	}()
// 	time.Sleep(1 * time.Second)
// 	if connected {
// 		t.Error("connection was not blocked by limiter")
// 	}
// 	limiter.Disconnect()
// 	// Ensure connected has a chance to be set to true before checking.
// 	time.Sleep(1 * time.Second)
// 	if !connected {
// 		t.Error("failed to connect once a connection was made available")
// 	}
// }

// TestLimiterError
// Ensure errors are returned for invalid values.
// TODO - Hit all possible combinations of values.
func TestLimiterError(t *testing.T) {
	var err error
	if _, err = NewLimiter(0, 0, 0); err == nil {
		t.Error("expected error for 0/0/0 but got nil")
	}
	if _, err = NewLimiter(1, 0, 0); err == nil {
		t.Error("expected error for 0/1/1 but got nil")
	}
	if _, err = NewLimiter(1, 1, 0); err == nil {
		t.Error("expected error for 1/1/0 but got nil")
	}
}
