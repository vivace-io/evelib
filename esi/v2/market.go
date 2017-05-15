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

type MarketOrder struct {
	OrderID    int       `json:"order_id"`
	BuyOrder   bool      `json:"is_buy_order"`
	Issued     time.Time `json:"issued"`
	LocationID int       `json:"location_id"`
	Duration   int       `json:"duration"`
}
