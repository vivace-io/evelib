package crest

import (
	"errors"
	"fmt"
)

// Constellation represents a constellation in Eve
type Constellation struct {
	Client  *Client        `json:"-"`
	Name    string         `json:"name"`
	ID      int            `json:"id"`
	Href    string         `json:"href"`
	Postion *Position      `json:"position"`
	Region  *Region        `json:"region"`
	Systems []*SolarSystem `json:"systems"`
}

func (c *Constellation) Complete() (err error) {
	if c.Client != nil {
		// TODO
	} else {
		err = errors.New("the Constellation model cannot be completed as Constellation.Client is nil")
	}
	return
}

// constellationCollection is an intermediate object for walking pages to retrieve all
// constellations from a CREST endpoint.
type constellationCollection struct {
	Items []*Constellation `json:"items"`
}

// AllConstellations returns a list of all constellations. If parameter `complete`
// is set to true, each constellation's endpoint will be visited once.
func (c *Client) AllConstellations() (result []*Constellation, err error) {
	collection := constellationCollection{}
	err = c.get("constellations/", &collection)
	if err != nil {
		return
	}
	result = collection.Items
	for _, r := range result {
		r.Client = c
	}
	return
}

// Constellation returns the constellation with a matching ID
func (c *Client) Constellation(id int) (result *Constellation, err error) {
	err = c.get(fmt.Sprintf("constellations/%v/", id), result)
	return
}
