package esi

import (
	"fmt"
	"time"
)

type MarketGroup struct {
	MarketGroupID int    `json:"market_group_id"`
	ParentGroupID int    `json:"parent_group_id"`
	Description   string `json:"description"`
	Types         []int  `json:"types"`
}

func (client *Client) MarketGroupIDs() (results []int, err error) {
	path := client.buildPath("/markets/groups/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) MarketGroupGet(id int) (result *MarketGroup, err error) {
	path := client.buildPath(fmt.Sprintf("/markets/groups/%v/", id))
	err = client.get(path, &result)
	if err != nil {
		return nil, err
	}
	return
}

// MarketPrice of an item type, as retrieved from the `/markets/prices/` endpoint.
type MarketPrice struct {
	TypeID        int     `json:"type_id"`
	AveragePrice  float32 `json:"average_price"`
	AdjustedPrice float32 `json:"adjusted_price"`
}

// MarketPrices returns the market prices for all published items in EVE Online
// from the `/markets/prices/` endpoint.
func (client *Client) MarketPrices() (results []*MarketPrice, err error) {
	path := client.buildPath("/markets/prices/")
	err = client.get(path, &results)
	return
}

// MarketRegionHistory general market statistics for a historical day in EVE.
type MarketRegionHistory struct {
	Average    float32 `json:"average"`
	Date       ESITime `json:"date"`
	Highest    float32 `json:"highest"`
	Lowest     float32 `json:"lowest"`
	OrderCount int     `json:"order_count"`
	Volume     int64   `json:"volume"`
}

// MarketRegionHistoryGet returns the historical market statistics for an item
// in a region.
func (client *Client) MarketRegionHistoryGet(regionID, typeID int) (results []*MarketRegionHistory, err error) {
	path := client.buildPath(fmt.Sprintf("/markets/%v/history/?type_id=%v", regionID, typeID))
	err = client.get(path, &results)
	return
}

type MarketOrder struct {
	OrderID    int       `json:"order_id"`
	BuyOrder   bool      `json:"is_buy_order"`
	Issued     time.Time `json:"issued"`
	LocationID int       `json:"location_id"`
	Duration   int       `json:"duration"`
}
