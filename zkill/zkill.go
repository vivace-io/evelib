package zkill

import (
	"net/http"
	"time"

	"github.com/vivace-io/evelib/crest"
)

// ZKillboardAPIURI is the default URL to zKillboard API
const (
	// DefaultAPIAddr for zKillboard
	DefaultAPIAddr = "https://zkillboard.com/api"
	// DefaultRedisQURI for zKillboard's RedisQ service
	DefaultRedisQURI = "https://redisq.zkillboard.com/listen.php"
)

var (
	UserAgent  string
	APIAddr    string
	RedisQAddr string
	clear      chan bool
	webClient  *http.Client
)

// Client is a client to access zKillboard's API.
type Client struct {
	UserAgent string
	Server    string
	clear     chan bool
}

func NewClient(address, userAgent string) (client *Client) {
	client = &Client{
		UserAgent: userAgent,
		Server:    address,
		clear:     make(chan bool, 100),
	}
	client.manage()
	return
}

func (this *Client) manage() {
	go func() {
		for {
			this.clear <- true
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

// Options is passed to query functions (i.e. CharacterKills, CorporationKills)
// and modifies the scope of the request and the kills returned.
// Options is passed to the client by design, as it is not required.
type Options struct {
	BeforeKillID int  // Returns kills before the kill ID, if set.
	AfterKillID  int  // Returns kills after the kill ID, if set.
	Solo         bool // Only returns solo kills if true.
	Kills        bool // Only returns kills if true.
	Losses       bool // Only returns losses if true.
	WSpace       bool // Only returns w-space kills if true.
	Limit        int  // Maximum kills returned (default 200 if not set)
}

// Kill is contains the CREST Killmail and zKillboard's extra data.
type Kill struct {
	KillID   int            `json:"killID"`
	Zkb      Zkb            `json:"zkb"`
	Killmail crest.Killmail `json:"killmail"`
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
