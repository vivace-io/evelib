package esi

import "fmt"

// RegionIDs returns a list of constellation IDs from ESI.
func (client *Client) RegionIDs() (results []int, err error) {
	path := buildPath("/universe/regions/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

// RegionGet accepts a region's ID and returns its ESI model.
func (client *Client) RegionGet(id int) (result *Region, err error) {
	path := buildPath(fmt.Sprintf("/universe/regions/%v/", id))
	err = client.get(path, &result)
	if err != nil {
		return nil, err
	}
	return
}

// Region in EVE Online.
type Region struct {
	RegionID       int    `json:"region_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Constellations []int  `json:"constellations"`
}

// ConstellationIDs returns a list of constellation IDs from ESI.
func (client *Client) ConstellationIDs() (results []int, err error) {
	path := buildPath("/universe/constellations/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

// ConstellationGet accepts a constellation's ID and returns its ESI model.
func (client *Client) ConstellationGet(id int) (result *Constellation, err error) {
	path := buildPath(fmt.Sprintf("/universe/constellations/%v/", id))
	err = client.get(path, &result)
	if err != nil {
		return nil, err
	}
	return
}

// Constellation in EVE Online.
type Constellation struct {
	ConstellationID int      `json:"constellation_id"`
	Name            string   `json:"name"`
	Position        Position `json:"position"`
	RegionID        int      `json:"region_id"`
	Systems         []int    `json:"systems"`
}

// SystemIDs returns a list of system IDs from ESI.
func (client *Client) SystemIDs() (results []int, err error) {
	path := buildPath("/universe/systems/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

// SystemGet accepts a solar system's ID and returns its ESI model.
func (client *Client) SystemGet(id int) (result *System, err error) {
	path := buildPath(fmt.Sprintf("/universe/systems/%v/", id))
	err = client.get(path, &result)
	if err != nil {
		return nil, err
	}
	return
}

// System is a Solar System in EVE.
type System struct {
	SystemID       int      `json:"system_id"`
	Name           string   `json:"name"`
	SecurityClass  string   `json:"security_class"`
	SecurityStatus float32  `json:"security_status"`
	Position       Position `json:"position"`
	Stargates      []int    `json:"stargates"`
	// TODO
	// Planets []Planet `json:"planets"`
}

// Position in 3D space in EVE Online.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
