package esi

import "fmt"

// ItemGroup from EVE Online.
type ItemGroup struct {
	GroupID    int32   `json:"group_id"`
	CategoryID int32   `json:"category_id"`
	Name       string  `json:"name"`
	Published  bool    `json:"published"`
	Types      []int32 `json:"types"`
}

// ItemGroupIDs returns a list of Item Group IDs from ESI.
func (client *Client) ItemGroupIDs() (results []int, err error) {
	for p := 1; ; p++ {
		var ids []int
		path := client.buildPath(fmt.Sprintf("/universe/groups/?page=%v", p))
		if err = client.get(path, &ids); err != nil {
			err = fmt.Errorf("failed to retrieve page %v of item groups: %v", p, err)
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

// ItemGroupGet returns the ESI Item Group representation for the given ID.
func (client *Client) ItemGroupGet(id int) (result *ItemGroup, err error) {
	path := client.buildPath(fmt.Sprintf("/universe/groups/%v/", id))
	err = client.get(path, &result)
	return
}

// Item is an item type from EVE Online.
// TODO - Dogma Attributes/Effects
type Item struct {
	TypeID      int     `json:"type_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Published   bool    `json:"published"`
	GroupID     int32   `json:"group_id"`
	Radius      float32 `json:"radius"`
	Volume      float32 `json:"volume"`
	Capacity    float32 `json:"capacity"`
	PortionSize int32   `json:"portion_size"`
	Mass        float32 `json:"mass"`
	IconID      int32   `json:"icon_id"`
}

// ItemIDs returns a list of all Item Type IDs from ESI.
func (client *Client) ItemIDs() (results []int32, err error) {
	for p := 1; ; p++ {
		var ids []int32
		path := client.buildPath(fmt.Sprintf("/universe/types/?page=%v", p))
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

// ItemGet returns the ESI Item Type representation for the given ID.
func (client *Client) ItemGet(id int) (result *Item, err error) {
	path := client.buildPath(fmt.Sprintf("/universe/types/%v/", id))
	err = client.get(path, &result)
	return
}
