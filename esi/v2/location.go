package esi

import "fmt"

func (client *Client) SystemIDs() (results []int, err error) {
	path := buildPath("/universe/systems/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

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
	ID             int      `json:"system_id"`
	Name           string   `json:"name"`
	SecurityClass  string   `json:"security_class"`
	SecurityStatus float32  `json:"security_status"`
	Position       Position `json:"position"`
	// TODO
	// Stargates []Stargate `json:"stargates"`
	// Planets []Planet `json:"planets"`
}

func (client *Client) ConstellationIDs() (results []int, err error) {
	path := buildPath("/universe/constellations/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) ConstellationGet(id int) (result *Constellation, err error) {
	err = client.get(fmt.Sprintf("/universe/constellations/%v/"), &result)
	if err != nil {
		return nil, err
	}
	return
}

type Constellation struct {
	ID       int      `json:"constellation_id"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
	RegionID int      `json:"region_id"`
	Systems  []int    `json:"systems"`
}

func (client *Client) RegionIDs() (results []int, err error) {
	path := buildPath("/universe/regions/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) RegionGet(id int) (result *Region, err error) {
	err = client.get(fmt.Sprintf("/universe/regions/%v/"), &result)
	if err != nil {
		return nil, err
	}
	return
}

type Region struct {
	ID             int    `json:"region_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Constellations int    `json:"constellations"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
