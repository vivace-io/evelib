package redisq

import (
	"strings"
	"time"
)

// Killmail from a kill
type Killmail struct {
	SolarSystem   SolarSystem `json:"solarSystem"`
	LocationID    int         `json:"locationID"`
	KillID        int         `json:"killID"`
	KillHash      string      `json:"killHash"`
	Timestamp     KillTime    `json:"killTime"`
	Attackers     []Attacker  `json:"attackers"`
	AttackerCount int         `json:"attackerCount"`
	Victim        Victim      `json:"victim"`
	Value         float32     `json:"value"`
}

// KillTime embeds time.Time and implements the UnmarshalJSON interface to
// handle CREST's non RFC 3339 timestamp.
type KillTime struct {
	time.Time
}

// UnmarshalJSON parses the timestamp from CREST in to Go's time.Time type.
func (t *KillTime) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse("2006.01.02 15:04:05", strings.Replace(string(b), "\"", "", 2))
	return err
}

// Attacker in a killmail
type Attacker struct {
	Character      Character   `json:"character"`
	SecurityStatus float32     `json:"securityStatus"`
	Corporation    Corporation `json:"corporation"`
	Alliance       Alliance    `json:"alliance"`
	ShipType       Type        `json:"shipType"`
	WeaponType     Type        `json:"weaponType"`
	DamageDone     int         `json:"damageDone"`
	FinalBlow      bool        `json:"finalBlow"`
}

// Victim in a killmail
type Victim struct {
	Character   Character   `json:"character"`
	Corporation Corporation `json:"corporation"`
	Alliance    Alliance    `json:"alliance"`
	ShipType    Type        `json:"shipType"`
	DamageTaken int         `json:"damageTaken"`
	Items       []Item      `json:"items"`
	Position    struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"position"`
}

// Item dropped/destroyed in a killmail
type Item struct {
	Singleton  int  `json:"singleton"`
	Type       Type `json:"itemType"`
	QtyDropped int  `json:"quantityDropped"`
	Flag       int  `json:"flag"`
}

// Corporation model for a Corporation entity in EVE Online.
type Corporation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

// Alliance in EVE.
type Alliance struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

// Character in EVE.
type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

// Type of an item in a killmail
type Type struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
	Icon struct {
		Href string `json:"href"`
	}
}

// SolarSystem represents a solar system in EVE.
type SolarSystem struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}
