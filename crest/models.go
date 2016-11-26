package crest

// Sovereignty model
type Sovereignty struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

/***** LOCATION DEFINITIONS *****/

type Stargate struct {
	ID          int         `json:"id"`
	Href        string      `json:"href"`
	Name        string      `json:"name"`
	System      SolarSystem `json:"system"`
	Position    Position    `json:"position"`
	Destination SolarSystem `json:"destination"`
	Type        Type        `json:"type"`
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
