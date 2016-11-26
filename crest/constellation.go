package crest

import "fmt"

// ConstellationsGetAll returns a list of all constellations.
func (c *Client) ConstellationsGetAll() (result []Constellation, err error) {
	collection := constellationCollection{}
	err = c.get("constellations/", &collection)
	if err != nil {
		return
	}
	result = collection.Items
	return
}

// ConstellationsGet returns a constellation with a matching ID, if it exists.
func (c *Client) ConstellationsGet(id int) (result Constellation, err error) {
	err = c.get(fmt.Sprintf("constellations/%v/", id), &result)
	return
}

// Constellation represents a constellation in Eve
type Constellation struct {
	Name    string         `json:"name"`
	ID      int            `json:"id"`
	Href    string         `json:"href"`
	Postion Position       `json:"position"`
	Region  Region         `json:"region"`
	Systems []*SolarSystem `json:"systems"`
}

// constellationCollection is an intermediate object for walking pages to retrieve all
// constellations from a CREST endpoint.
type constellationCollection struct {
	Items []Constellation `json:"items"`
}
