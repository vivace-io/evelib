package crest

import "fmt"

// Constellations returns a list of all constellations. If parameter `complete`
// is set to true, each constellation's endpoint will be visited once.
func (c *Client) Constellations(complete bool) (result []*Constellation, err error) {
	collection := constellationCollection{}
	err = c.get("constellations/", &collection)
	if err != nil {
		return
	}
	result = collection.Items
	if complete {
		// TODO
	}
	return
}

// GetConstellation returns the constellation with a matching ID
func (c *Client) GetConstellation(id int) (result *Constellation, err error) {
	err = c.get(fmt.Sprintf("constellations/%v/", id), result)
	return
}

// Constellation represents a constellation in Eve
type Constellation struct {
	Name    string         `json:"name"`
	ID      int            `json:"id"`
	Href    string         `json:"href"`
	Postion *Position      `json:"position"`
	Region  *Region        `json:"region"`
	Systems []*SolarSystem `json:"systems"`
}

// constellationCollection is an intermediate object for walking pages to retrieve all
// constellations from a CREST endpoint.
type constellationCollection struct {
	Items []*Constellation `json:"items"`
}
