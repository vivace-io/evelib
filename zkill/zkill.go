package zkill

import (
	"net/http"
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
}

// Kill is contains the CREST Killmail and zKillboard's extra data.
type Kill struct {
	KillID   int             `json:"killID"`
	Zkb      *Zkb            `json:"zkb"`
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

func manage() {
	clear = make(chan bool, 100)
	go func() {
		for {
			clear <- true
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
