package crest

import "fmt"

// Alliance in Eve
type Alliance struct {
	ID                  int           `json:"id"`
	Name                string        `json:"name"`
	ShortName           string        `json:"shortName"`
	Href                string        `json:"href"`
	StartDateStr        string        `json:"startDate"`
	Description         string        `json:"description"`
	CorporationsCount   int           `json:"corporationsCount"`
	ExecutorCorporation Corporation   `json:"executorCorporation"`
	Deleted             bool          `json:"deleted"`
	CreatorCorporation  Corporation   `json:"creatorCorporation"`
	URL                 string        `json:"url"`
	CreatorCharacter    Character     `json:"creatorCharacter"`
	Corporations        []Corporation `json:"corporations"`
}

func (c *Client) Alliance(id int) (result Alliance, err error) {
	err = c.get(fmt.Sprintf("alliances/%v/", id), &result)
	return
}

func (c *Client) AllAlliances() (result []*Alliance, err error) {
	// TODO
	return
}
