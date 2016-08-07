package zkill

import (
	"net/http"

	"gopkg.in/vivace-io/evelib.v0/crest"
)

// ZKillboardAPIURI is the default URL to zKillboard API
const ZKillboardAPIURI = "https://zkillboard.com/api/"

// Client to zKillboard's API
type Client struct {
	APIAddress string
	webClient  http.Client
	UserAgent  string
}

// New returns a new client for zKillboard
func New(userAgent string) (*Client, error) {
	c := &Client{
		APIAddress: ZKillboardAPIURI,
		webClient:  http.Client{},
		UserAgent:  userAgent,
	}
	crest.UserAgent = userAgent
	return c, c.checkUserAgent()
}

// Kill is contains the CREST Killmail and zKillboard's extra data.
type Kill struct {
	KillID   int             `json:"killID"`
	Zkb      Zkb             `json:"zkb"`
	Killmail *crest.Killmail `json:"killmail"`
}

// Zkb is the extra data returned from zKillboard's API
type Zkb struct {
	LocationID int     `json:"locationID"`
	Hash       string  `json:"hash"`
	TotalValue float32 `json:"totalValue"`
	Points     int     `json:"points"`
	Href       string  `json:"href"`
}

type response struct {
	Kills []Kill
}
