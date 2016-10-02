package crest

import (
	"errors"
	"fmt"
)

// Alliance in Eve
type Alliance struct {
	Client              *Client        `json:"-"`
	ID                  int            `json:"id"`
	Name                string         `json:"name"`
	ShortName           string         `json:"shortName"`
	Href                string         `json:"href"`
	StartDateStr        string         `json:"startDate"`
	Description         string         `json:"description"`
	CorporationsCount   int            `json:"corporationsCount"`
	ExecutorCorporation *Corporation   `json:"executorCorporation"`
	Deleted             bool           `json:"deleted"`
	CreatorCorporation  *Corporation   `json:"creatorCorporation"`
	URL                 string         `json:"url"`
	CreatorCharacter    *Character     `json:"creatorCharacter"`
	Corporations        []*Corporation `json:"corporations"`
}

func (alli *Alliance) Complete() (err error) {
	if alli.Client != nil {
		// TODO
	} else {
		err = errors.New("the Alliance model cannot be completed as Alliance.Client is nil")
	}
	return
}

func (c *Client) Alliance(id int) (result *Alliance, err error) {
	err = c.get(fmt.Sprintf("alliances/%v/", id), &result)
	if result != nil {
		result.Client = c
	}
	return
}

func (c *Client) AllAlliances() (result []*Alliance, err error) {
	// TODO
	return
}
