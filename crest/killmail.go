package crest

import "fmt"

func (c *Client) GetKillmail(id int, hash string) (result *Killmail, err error) {
	err = c.get(fmt.Sprintf("killmails/%v/%v/", id, hash), &result)
	return
}

// Killmail from a kill
type Killmail struct {
	SolarSystem   SolarSystem `json:"solarSystem"`
	KillID        int         `json:"killID"`
	TimestampStr  string      `json:"killTime"`
	Attackers     []Attacker  `json:"attackers"`
	AttackerCount int         `json:"attackerCount"`
	Victim        Victim      `json:"victim"`
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
