package crest

import (
	"errors"
	"fmt"
	"strconv"
)

// Regions returns all regions in Eve. If parameter `complete` is set to true, the
// method will visit each individual system endpoint and gather full data.
func Regions(complete bool) (result []*Region, err error) {
	collection := regionCollection{}
	err = fetch("regions/", &collection)
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

func GetRegion(id int) (result *Region, err error) {
	err = fetch("regions/"+strconv.Itoa(id)+"/", &result)
	return
}

// Region represents a region in Eve
type Region struct {
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	ID             int              `json:"id"`
	Href           string           `json:"href"`
	Constellations []*Constellation `json:"constellations"`
}

// Complete fills/updates the Region model, but does not walk its constellations.
func (this *Region) Complete() error {
	if this.Href == "" {
		if this.ID == 0 {
			return errors.New("unable to fill region data - href and ID are unspecified")
		}
		return fetch(fmt.Sprintf("regions/%v/", this.ID), this)
	}
	return fetch(this.Href, this)
}

// Walk visits the Region's endpoint, then visits the endpoint of each Constellation.
func (this *Region) Walk() error {
	err := this.Complete()
	if err != nil {
		return err
	}
	for _, c := range this.Constellations {
		err := c.Complete()
		if err != nil {
			return fmt.Errorf("unable to completely walk region's constellations - constellation %v failed with error %v", c.ID, err)
		}
		c.Region = this
	}
	return nil
}

// regionCollection is a collection of regions
type regionCollection struct {
	Items []*Region `json:"items"`
}
