package esi

import (
	"fmt"
	"time"
)

type MarketGroup struct {
	ID            int    `json:"market_group_id"`
	ParentGroupId int    `json:"parent_group_id"`
	Description   string `json:"description"`
	Types         []int  `json:"types"`
}

func (client *Client) MarketGroupIDs() (results []int, err error) {
	path := buildPath("/markets/groups/")
	err = client.get(path, &results)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) MarketGroupGet(id int) (result *MarketGroup, err error) {
	path := buildPath(fmt.Sprintf("/markets/groups/%v/", id))
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

func (client *Client) MarketPrices() (results []*MarketPrice, err error) {
	path := buildPath("/markets/prices/")
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
