package crest

import (
	"fmt"
	"log"
)

// InventoryTypesGet (https://crest-tq.eveonline.com/inventory/types/)
// Returns basic information on all Types from Eve Online. Note that not all
// data is present in each model, as the specific endpoint for each is not
// walked by this method.
func (c *Client) InventoryTypesGet() (result []Type, err error) {
	pageCount := 1000 // picking a random number to be adjusted later...
	for current := 1; current < pageCount+1; current++ {
		page := new(typePage)
		err = c.get(fmt.Sprintf("inventory/types/?page=%v", current), &page)
		if err != nil {
			err = fmt.Errorf("failed to retrieve page with error %v", err)
			return
		}
		result = append(result, page.Items...)
		if pageCount == 1000 {
			pageCount = page.PageCount
			log.Println(pageCount)
		}
	}
	return
}

// Type represents an item type as it is retrieved from CREST.
type Type struct {
	ID          int     `json:"id"`
	Href        string  `json:"href"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IconID      int     `json:"iconID"`
	Icon        Href    `json:"icon"`
	Volume      float32 `json:"volume"`
	Radius      float32 `json:"radius"`
	Published   bool    `json:"published"`
	Mass        float32 `json:"mass"`
	PortionSize float32 `json:"portionSize"`
}

// typePage is an intermediate for walking pages on /invetory/types/
type typePage struct {
	PageCount int    `json:"pageCount"`
	Items     []Type `json:"items"`
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
