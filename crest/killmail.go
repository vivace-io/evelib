package crest

import (
	"fmt"
	"strings"
	"time"
)

func (c *Client) KillmailGet(id int, hash string) (result *Killmail, err error) {
	err = c.get(fmt.Sprintf("killmails/%v/%v/", id, hash), &result)
	if err == nil {
		// Only assign hash on success.
		result.KillHash = hash
	}
	return
}

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

type KillTime struct {
	time.Time
}

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
	Position    Position    `json:"position"`
}

// Item dropped/destroyed in a killmail
type Item struct {
	Singleton  int  `json:"singleton"`
	Type       Type `json:"itemType"`
	QtyDropped int  `json:"quantityDropped"`
	Flag       int  `json:"flag"`
}
