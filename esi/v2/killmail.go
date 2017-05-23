package esi

import (
	"fmt"
	"time"
)

func (client *Client) KillmailGet(id int, hash string) (kill *Killmail, err error) {
	path := buildPath(fmt.Sprintf("/killmails/%v/%v/", id, hash))
	err = client.get(path, &kill)
	if err != nil {
		return nil, err
	}
	kill.KillmailHash = hash
	return
}

type KillmailTime struct {
	time.Time
}

type Killmail struct {
	KillmailID    int          `json:"killmail_id"`
	KillmailHash  string       `json:"killmail_hash"`
	KillmailTime  KillmailTime `json:"killmail_time"`
	SolarSystemID int          `json:"solar_system_id"`
	Victim        struct {
		DamageTaken   int `json:"damage_taken"`
		ShipTypeID    int `json:"ship_type_id"`
		CharacterID   int `json:"character_id"`
		CorporationID int `json:"corporation_id"`
		AllianceID    int `json:"alliance_id"`
		Position      struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"position"`
		Items []struct {
			ItemTypeID      int `json:"item_type_id"`
			Singleton       int `json:"singleton"`
			Flag            int `json:"flag"`
			QuantityDropped int `json:"quantity_dropped"`
		} `json:"items"`
	} `json:"victim"`
	Attackers []struct {
		CharacterID    int     `json:"character_id"`
		CorporationID  int     `json:"corporation_id"`
		AllianceID     int     `json:"alliance_id"`
		ShipTypeID     int     `json:"ship_type_id"`
		WeaponTypeID   int     `json:"weapon_type_id"`
		FinalBlow      bool    `json:"final_blow"`
		DamageDone     int     `json:"damage_done"`
		SecurityStatus float32 `json:"security_status"`
	} `json:"attackers"`
}
