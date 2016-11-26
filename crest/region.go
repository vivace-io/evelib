package crest

import "strconv"

// RegionsGetAll returns all regions in Eve.
func (c *Client) RegionsGetAll() (result []Region, err error) {
	collection := regionCollection{}
	err = c.get("regions/", &collection)
	if err != nil {
		return
	}
	result = collection.Items
	return
}

// RegionsGet returns a region that matches the `id` parameter, if it exists.
func (c *Client) RegionsGet(id int) (result Region, err error) {
	err = c.get("regions/"+strconv.Itoa(id)+"/", &result)
	return
}

// Region represents a region in Eve
type Region struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	ID             int             `json:"id"`
	Href           string          `json:"href"`
	Constellations []Constellation `json:"constellations"`
}

// regionCollection is a collection of regions
type regionCollection struct {
	Items []Region `json:"items"`
}
