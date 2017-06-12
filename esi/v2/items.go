package esi

import "fmt"

// ItemIDs returns a list of all Item Type IDs from ESI.
func (client *Client) ItemIDs() (results []int, err error) {
	for p := 1; ; p++ {
		var ids []int
		path := buildPath(fmt.Sprintf("/universe/types/?page=%v", p))
		if err = client.get(path, &ids); err != nil {
			err = fmt.Errorf("failed to retrieve page %v of types: %v", p, err)
			break
		}
		if len(ids) != 0 {
			results = append(results, ids...)
		} else {
			break
		}
	}
	return
}

// Item is an item type from EVE Online.
// TODO - Dogma Attributes/Effects
type Item struct {
	TypeID      int     `json:"type_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Published   bool    `json:"published"`
	GroupID     int     `json:"group_id"`
	Radius      float32 `json:"radius"`
	Volume      float32 `json:"volume"`
	Capacity    float32 `json:"capacity"`
	PortionSize int     `json:"portion_size"`
	Mass        float32 `json:"mass"`
	IconID      int     `json:"icon_id"`
}

// ItemGet returns the ESI Item Type representation for the given ID.
func (client *Client) ItemGet(id int) (result *Item, err error) {
	path := buildPath(fmt.Sprintf("/universe/types/%v/", id))
	err = client.get(path, &result)
	return
}
