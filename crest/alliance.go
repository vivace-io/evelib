package crest

import "fmt"

func GetAlliance(id int) (result *Alliance, err error) {
	err = fetch(fmt.Sprintf("alliances/%v/", id), &result)
	return
}

// Alliance in Eve
type Alliance struct {
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