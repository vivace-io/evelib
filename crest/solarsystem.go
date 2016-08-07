package crest

import (
	"errors"
	"fmt"
)

// SolarSystems returns all system in Eve with partial data.
func SolarSystems(Walk bool) (systems []*SolarSystem, err error) {
	collection := systemCollection{}
	err = fetch("solarsystems/", &collection)
	if err != nil {
		return
	}
	systems = collection.Items
	if Walk {
		for _, s := range systems {
			err = s.Walk()
			if err != nil {
				return
			}
		}
	}
	return
}

func GetSolarSystem(id int, Walk bool) (system *SolarSystem, err error) {
	err = fetch(fmt.Sprintf("solarsystems/%v/", id), &system)
	if err != nil {
		return
	}
	if Walk {
		err = system.Walk()
	}
	return
}

type systemCollection struct {
	Items []*SolarSystem `json:"items"`
}

// SolarSystem represents a solar system in Eve
type SolarSystem struct {
	ID             int            `json:"id"`
	Href           string         `json:"href"`
	Name           string         `json:"name"`
	SecurityStatus float32        `json:"securityStatus"`
	SecurityClass  string         `json:"securityClass"`
	Position       *Position      `json:"position, omitempty"`
	Constellation  *Constellation `json:"constellation, omitempty"`
	Planets        []*Planet      `json:"planet"`
	Sovereignty    *Sovereignty   `json:"sovereignty, omitempty"`
}

// Complete fills/updates the SolarSystem model, but does not walk its planets.
func (this *SolarSystem) Complete() error {
	if this.Href == "" {
		if this.ID == 0 {
			return errors.New("unable to fill solar system information - href and ID are unspecified.")
		}
		return fetch(fmt.Sprintf("solarsystems/%v/", this.ID), this)
	}
	return fetch(this.Href, this)
}

// Walk visits the SolarSystem's endpoint, then visits the endpoint of each planet.
func (this *SolarSystem) Walk() error {
	if len(this.Planets) == 0 {
		return errors.New("unable to fill system data - no planets in system")
	}
	for _, p := range this.Planets {
		err := p.Walk()
		p.SolarSystem = this
		if err != nil {
			return err
		}
	}
	return nil
}

// Planet in a solar system
type Planet struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	SolarSystem *SolarSystem `json:"solarSystem"`
	Href        string       `json:"href"`
	Moons       []*Moon      `json:"moons"`
}

// Complete fills/updates the Planet model, but does not walk its moons.
func (this *Planet) Complete() error {
	if this.Href == "" {
		if this.ID == 0 {
			return errors.New("unable to fill planet information - href and ID are unspecified")
		}
		return fetch(fmt.Sprintf("planets/%v/", this.ID), this)
	}
	return fetch(this.Href, this)
}

// Walk visits the Planet's endpoint, then visits the endpoint of each moon.
func (this *Planet) Walk() error {
	err := this.Complete()
	if err != nil {
		return err
	}
	for _, m := range this.Moons {
		err = m.Complete()
		m.SolarSystem = this.SolarSystem
		if err != nil {
			return err
		}
		m.SolarSystem = this.SolarSystem
	}
	return nil
}

// Moon of a planet
type Moon struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	SolarSystem *SolarSystem `json:"solarSystem"`
	Type        *Type        `json:"type"`
	Href        string       `json:"href"`
}

// Complete fills/updates the moon model.
func (this *Moon) Complete() (err error) {
	if this.Href == "" {
		if this.ID == 0 {
			return errors.New("unable to fill moon information - href and ID are unspecified.")
		}
		return fetch(fmt.Sprintf("moons/%v/", this.ID), this)
	}
	return fetch(this.Href, this)
}
