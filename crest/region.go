package crest

import "strconv"

// Regions returns all regions in Eve. If parameter `complete` is set to true, the
// method will visit each individual system endpoint and gather full data.
func (c *Client) Regions(complete bool) (result []Region, err error) {
	collection := regionCollection{}
	err = c.get("regions/", &collection)
	if err != nil {
		return
	}
	result = collection.Items
	if complete {
		// TODO
	}
	return
}

func (c *Client) GetRegion(id int) (result Region, err error) {
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
