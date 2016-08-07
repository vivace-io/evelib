package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// ZKillboardRedisQ default endpoint
const ZKillboardRedisQ = "https://redisq.zkillboard.com/listen.php"

var alreadyRunning = false

// RedisQClient is a client to zKillboard's RedisQ service
type RedisQClient struct {
	RedisQURI string
	webClient http.Client
	UserAgent string
}

// NewRedisQ returns a client with default endpoints
func NewRedisQ() *RedisQClient {
	c := &RedisQClient{
		RedisQURI: ZKillboardRedisQ,
		webClient: http.Client{},
	}
	return c
}

type redisqResp struct {
	Kill Kill `json:"package"`
}

// FetchKillmails starts retrieving Killmails from ZKillboard RedisQ, sending
// them (and any errors encountered) through the channels passed
func (c *RedisQClient) FetchKillmails(output chan Kill, errChan chan error) {
	if c.UserAgent == "" {
		errChan <- errors.New("user-agent must be set first")
		return
	}
	go func() {
		for {
			kill, err := c.fetchRedisQ()
			if err != nil {
				errChan <- err
			} else {
				output <- kill
			}
		}
	}()
}

func (c *RedisQClient) fetchRedisQ() (k Kill, err error) {
	request, err := http.NewRequest("GET", c.RedisQURI, nil)
	if err != nil {
		return
	}
	request.Header.Add("User-Agent", c.UserAgent)

	rawresp, err := c.webClient.Do(request)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(rawresp.Body)
	if err != nil {
		return
	}
	zresp := redisqResp{}
	err = json.Unmarshal(body, &zresp)
	k = zresp.Kill
	return
}
