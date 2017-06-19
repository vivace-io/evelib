package esi

import "time"

// Status of the EVE Online server.
type Status struct {
	Players       int       `json:"players"`
	ServerVersion string    `json:"server_version"`
	StartTime     time.Time `json:"start_time"`
}

// StatusGet returns the current Status of the EVE Online server.
func (client *Client) StatusGet() (status *Status, err error) {
	err = client.get(client.buildPath("/status/"), &status)
	return
}
