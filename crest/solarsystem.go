package crest

import "fmt"

// SolarSystemsGetAll returns all solar systems in the game.
func (c *Client) SolarSystemsGetAll() (systems []*SolarSystem, err error) {
	collection := systemCollection{}
	err = c.get("solarsystems/", &collection)
	if err != nil {
		return
	}
	systems = collection.Items
	return
}

// SolarSystemsGet returns a solar system that matches the `id` parameter.
func (c *Client) SolarSystemsGet(id int) (system *SolarSystem, err error) {
	err = c.get(fmt.Sprintf("solarsystems/%v/", id), &system)
	if err != nil {
		return
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

// Planet in a solar system
type Planet struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	SolarSystem *SolarSystem `json:"solarSystem"`
	Href        string       `json:"href"`
	Moons       []*Moon      `json:"moons"`
}

// GetPlanet returns the planet with the corressponding ID.
func (c *Client) GetPlanet(id int) (planet *Planet, err error) {
	err = c.get(fmt.Sprintf("planets/%v/", id), &planet)
	if err != nil {
		return
	}
	return
}

// Moon of a planet
type Moon struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	SolarSystem *SolarSystem `json:"solarSystem"`
	Type        *Type        `json:"type"`
	Href        string       `json:"href"`
}

func (c *Client) GetMoon(id int) (moon *Moon, err error) {
	err = c.get(fmt.Sprintf("moons/%v/", id), &moon)
	if err != nil {
		return
	}
	return
}
