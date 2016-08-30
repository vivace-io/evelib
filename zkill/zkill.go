package zkill

import (
	"net/http"
<<<<<<< HEAD
	"time"

	"gopkg.in/evelib.v0/crest"
)

// ZKillboardAPIURI is the default URL to zKillboard API
const (
	// DefaultAPIAddr for zKillboard
	DefaultAPIAddr = "https://zkillboard.com/api/"
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

func Init(userAgent string) error {
	userAgent = UserAgent
	APIAddr = DefaultAPIAddr
	RedisQAddr = DefaultRedisQURI
	webClient = new(http.Client)
	return checkUserAgent()
}

type Options struct {
	BeforeKillID int
	AfterKillID  int
	Solo         bool
	Kills        bool
	Losses       bool
	WSpace       bool
	Limit        int
=======

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
>>>>>>> 231cf7d91084be67f3f16cd3fd696295b1fc6653
}

// Kill is contains the CREST Killmail and zKillboard's extra data.
type Kill struct {
	KillID   int             `json:"killID"`
<<<<<<< HEAD
	Zkb      *Zkb            `json:"zkb"`
=======
	Zkb      Zkb             `json:"zkb"`
>>>>>>> 231cf7d91084be67f3f16cd3fd696295b1fc6653
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
<<<<<<< HEAD

func manage() {
	clear = make(chan bool, 100)
	go func() {
		for {
			clear <- true
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
=======
>>>>>>> 231cf7d91084be67f3f16cd3fd696295b1fc6653
