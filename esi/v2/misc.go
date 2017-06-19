package esi

import (
	"fmt"
	"strings"
	"time"
)

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

// ESITime implements the json.Unmarshaler interface to parse times/dates that
// returned by ESI and are not RFC 3330 compliaint.
type ESITime struct {
	time.Time
}

func (et *ESITime) UnmarshalJSON(b []byte) (err error) {
	fmt.Println(strings.Replace(string(b), "\"", "", 2))
	et.Time, err = time.Parse("2006-01-02", strings.Replace(string(b), "\"", "", 2))
	return
}
