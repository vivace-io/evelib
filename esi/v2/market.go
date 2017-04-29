package esi

import "time"

type MarketOrder struct {
	OrderID    int       `json:"order_id"`
	BuyOrder   bool      `json:"is_buy_order"`
	Issued     time.Time `json:"issued"`
	LocationID int       `json:"location_id"`
	Duration   int       `json:"duration"`
}
