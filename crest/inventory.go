package crest

import "fmt"

// TypesGetAll (https://crest-tq.eveonline.com/inventory/types/)
// Returns all Types from Eve Online. If `complete` is true, it will fill all
// missing information for each type.
func (c *Client) TypesGetAll() (result []Type, err error) {
	collection := typeCollection{}
	err = c.get("inventory/types/", &collection)
	if err != nil {
		return
	}
	result = append(result, collection.Items...)
	for collection.Next.Href != "" {
		err = c.get("types/", &collection)
		if err != nil {
			err = fmt.Errorf("unable to pull all item types with error %v", err)
			return
		}
	}
	return
}

// Type represents an item type as it is retrieved from CREST.
type Type struct {
	Name        string  `json:"name"`
	ID          int     `json:"id"`
	IconID      int     `json:"iconID"`
	Description string  `json:"description"`
	Volume      float32 `json:"volume"`
	Radius      float32 `json:"radius"`
	Published   bool    `json:"published"`
	Mass        float32 `json:"mass"`
	PortionSize float32 `json:"portionSize"`
	Href        string  `json:"href"`
}

// typeCollection is an intermediate object for walking pages to retrieve all
// types from a CREST endpoint
type typeCollection struct {
	PageCount int    `json:"pageCount"`
	Items     []Type `json:"items"`
	Next      struct {
		Href string `json:"href"`
	} `json:"next, omitempty"`
}

// Group represents an inventory group.
type Group struct {
	Name        string   `json:"name"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
	Types       []Type   `json:"types"`
	Published   bool     `json:"true"`
}

// Category represents an inventory category.
type Category struct {
	Name      string  `json:"name"`
	Groups    []Group `json:"groups"`
	Href      string  `json:"href"`
	Published bool    `json:"published"`
}
