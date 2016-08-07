package crest

import (
	"errors"
	"fmt"
)

// Constellations returns a list of all constellations. If parameter `complete`
// is set to true, each constellation's endpoint will be visited once.
func Constellations(complete bool) (result []*Constellation, err error) {
	collection := constellationCollection{}
	err = fetch("constellations/", &collection)
	if err != nil {
		return
	}
	result = collection.Items
	if complete {
		for _, r := range result {
			err = r.Complete()
			if err != nil {
				return
			}
		}
	}
	return
}

// GetConstellation returns the constellation with a matching ID
func GetConstellation(id int) (result *Constellation, err error) {
	err = fetch(fmt.Sprintf("constellations/%v/", id), result)
	return
}

// Constellation represents a constellation in Eve
type Constellation struct {
	Name    string         `json:"name"`
	ID      int            `json:"id"`
	Href    string         `json:"href"`
	Postion *Position      `json:"position"`
	Region  *Region        `json:"region"`
	Systems []*SolarSystem `json:"systems"`
}

// Complete fills/updates the region model, but does not walk its SolarSystems.
func (this *Constellation) Complete() error {
	if this.Href == "" {
		if this.ID == 0 {
			return errors.New("unable to fill constellation data - href and ID are unspecified.")
		}
		return fetch(fmt.Sprintf("constellations/%v/", this.ID), this)
	}
	return fetch(this.Href, this)
}

// Walk visits the endpoint of each constellation in the constellation and fills/updates it.
func (this *Constellation) Walk() error {
	if len(this.Systems) == 0 {
		return errors.New("unable to walk constellation model systems - no systems in constellation")
	}
	for _, s := range this.Systems {
		err := s.Complete()
		if err != nil {
			return fmt.Errorf("unable to completely walk constellation's systems - system %v failed to walk with error %v", s.ID, err)
		}
		s.Constellation = this
	}
	return nil
}

// constellationCollection is an intermediate object for walking pages to retrieve all
// constellations from a CREST endpoint.
type constellationCollection struct {
	Items []*Constellation `json:"items"`
}
