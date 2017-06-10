package zkill

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Killmail is a killmail as it is represented by the ZKillboard API.
type Killmail struct {
	KillID        int        `json:"killID"`
	SolarSystemID int        `json:"solarSystemID"`
	KillTime      KillTime   `json:"killTime"`
	Victim        Victim     `json:"victim"`
	Attackers     []Attacker `json:"attackers"`
	Items         []Item     `json:"items"`
	Position      Position   `json:"position"`
	Zkb           Zkb        `json:"zkb"`
}

type KillTime struct {
	time.Time
}

func (kt *KillTime) UnmarshalJSON(data []byte) (err error) {
	// Trim leading/trailing quotation marks.
	str := strings.Trim(string(data), "\"")
	kt.Time, err = time.Parse("2006-01-02 15:04:05", str)
	return
}

// InvolvedEntity holds values shared between attackers and victims and is
// embedded in to both models.
type InvolvedEntity struct {
	CharacterID     int    `json:"characterID"`
	CharacterName   string `json:"characterName"`
	CorporationID   int    `json:"corporationID"`
	CorporationName string `json:"corporationName"`
	AllianceID      int    `json:"allianceID"`
	AllianceName    string `json:"allianceName"`
	FactionID       int    `json:"factionID"`
	FactionName     string `json:"factionName"`
	ShipTypeID      int    `json:"shipTypeID"`
}

// Victim in a killmail.
type Victim struct {
	InvolvedEntity
	DamageTaken int `json:"damageTaken"`
}

// Attacker in a killmail.
type Attacker struct {
	InvolvedEntity
	SecurityStatus float32 `json:"securityStatus"`
	DamageDone     int     `json:"damageDone"`
	FinalBlow      uint8   `json:"finalBlow"`
	WeaponTypeID   int     `json:"weaponTypeID"`
}

// Item that was fitted to or held within a ship.
type Item struct {
	TypeID       int   `json:"typeID"`
	Flag         int   `json:"flag"`
	QtyDropped   int   `json:"qtyDropped"`
	QtyDestroyed int   `json:"qtyDestroyed"`
	Singleton    uint8 `json:"singleton"`
}

// Position of a killmail.
type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// Zkb is the extra data returned from zKillboard's API
type Zkb struct {
	LocationID int     `json:"locationID"`
	Hash       string  `json:"hash"`
	TotalValue float32 `json:"totalValue"`
	Points     int     `json:"points"`
}

// KillmailGet takes a killmail ID and returns its ZKillboard representation.
func (client *Client) KillmailGet(killID int) (result *Killmail, err error) {
	var resp []*Killmail
	if err = client.fetch(fmt.Sprintf("/killID/%v/", killID), &resp); err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		err = errors.New("empty response from ZKillboard")
		return
	}
	result = resp[0]
	return
}

// Historical returns all kill ID's and their accompanying hashes submittted to
// zKillboard on the provided date.
func (client *Client) Historical(date time.Time) (killmails map[int]string, err error) {
	killmails = make(map[int]string)
	results := make(map[string]string)
	client.fetch(fmt.Sprintf("/history/%v/", date.Format("20060102")), &results)
	var id int
	for key, val := range results {
		if id, err = strconv.Atoi(key); err != nil {
			err = fmt.Errorf("unable to convert string value '%v' to type int: %v", key, err)
			return
		} else {
			killmails[id] = val
		}
	}
	return
}
