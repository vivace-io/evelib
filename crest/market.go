package crest

import (
	"fmt"
)

type MarketGroup struct {
}

type MarketType struct {
	ID          int         `json:"id"`
	Type        Type        `json:"type"`
	Name        string      `json:"name"`
	MarketGroup MarketGroup `json:"marketGroup"`
}

// MarketTypesGet (https://crest-tq.eveonline.com/market/types/)
// Returns basic information on all Markett Types from Eve Online. Note that not
// all data is present in each model, as the specific endpoint for each is not
// walked by this method.
func (c *Client) MarketTypesGet() (result []MarketType, err error) {
	pageCount := 1000 // picking a random number to be adjusted later...
	for current := 1; current < pageCount+1; current++ {
		page := new(marketTypePage)
		err = c.get(fmt.Sprintf("market/types/?page=%v", current), &page)
		if err != nil {
			err = fmt.Errorf("failed to retrieve page with error %v", err)
			return
		}
		result = append(result, page.Items...)
		if pageCount == 1000 {
			pageCount = page.PageCount
		}
	}
	return
}

type marketTypePage struct {
	PageCount int          `json:"pageCount"`
	Items     []MarketType `json:"items"`
}
