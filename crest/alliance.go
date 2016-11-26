package crest

import "fmt"

// AlliancesGet returns an alliance with the id matching the parameter, if any.
func (c *Client) AlliancesGet(id int) (result Alliance, err error) {
	err = c.get(fmt.Sprintf("alliances/%v/", id), &result)
	return
}

// AlliancesGetAll returns all known alliances in the game.
func (c *Client) AlliancesGetAll() (result []*Alliance, err error) {
	// TODO
	err = ErrNotImplemented
	return
}

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
