package redisq

import (
	"strings"
	"time"
)

// Killmail from a kill
type Killmail struct {
	ID            int        `json:"killmail_id"`
	Hash          string     `json:"killmail_hash"`
	SolarSystemID int        `json:"solar_system_id"`
	Timestamp     KillTime   `json:"killmail_time"`
	Victim        Victim     `json:"victim"`
	Attackers     []Attacker `json:"attackers"`
	Zkb           Zkb        `json:"zkb"`
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
	CharacterID    int     `json:"character_id"`
	CorporationID  int     `json:"corporation_id"`
	AllianceID     int     `json:"alliance_id"`
	ShipTypeID     int     `json:"ship_type_id"`
	WeaponTypeID   int     `json:"weapon_type_id"`
	DamageDone     float32 `json:"damage_done"`
	FinalBlow      bool    `json:"final_blow"`
	SecurityStatus float32 `json:"security_status"`
}

// Victim in a killmail
type Victim struct {
	CharacterID   int    `json:"character_id"`
	CorporationID int    `json:"corporation_id"`
	AllianceID    int    `json:"alliance_id"`
	ShipTypeID    int    `json:"ship_type_id"`
	DamageTaken   int    `json:"damage_taken"`
	Items         []Item `json:"items"`
	Position      struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"position"`
}

// Item dropped/destroyed in a killmail
type Item struct {
	ItemTypeID        int `json:"item_type_id"`
	Flag              int `json:"flag"`
	Singleton         int `json:"singleton"`
	QuantityDropped   int `json:"quantity_dropped"`
	QuantityDestroyed int `json:"quantity_destroyed"`
}

// Zkb is the meta data returned from ZKillboard.
type Zkb struct {
	Hash        string  `json:"hash"`
	FittedValue float32 `json:"fittedValue"`
	TotalValue  float32 `json:"totalValue"`
	Points      int     `json:"points"`
	NPC         bool    `json:"npc"`
	Solo        bool    `json:"solo"`
	AWOX        bool    `json:"aox"`
	Href        string  `json:"href"`
}
