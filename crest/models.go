package crest

import "fmt"

// Sovereignty model
type Sovereignty struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

/***** LOCATION DEFINITIONS *****/

type Stargate struct {
	ID          int          `json:"id"`
	Href        string       `json:"href"`
	Name        string       `json:"name"`
	System      *SolarSystem `json:"system"`
	Position    Position     `json:"position"`
	Destination *SolarSystem `json:"destination"`
	Stargate    *Stargate    `json:"stargate"`
	Type        Type         `json:"type"`
}

// Position represents a 3D Cartesian position in Eve
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Location represents a location in Eve
type Location struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

// UnexpectedHTTPError for all HTTP status code errors not covered.
type UnexpectedHTTPError struct {
	StatusCode int
}

func (e UnexpectedHTTPError) Error() string {
	return fmt.Sprintf("unexpected HTTP status code %v", e.StatusCode)
}

/***** ETC. DEFINITIONS *****/

// Href to resource
type Href struct {
	Href string `json:"href"`
}

// War in Eve
type War struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}
