package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vivace-io/evelib/crest"
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
	}
	client.manage()
	return
}

func (client *Client) fetch(path string, model interface{}) error {
	select {
	case <-client.clear:
		request, err := http.NewRequest("GET", client.Server+path, nil)
		if err != nil {
			return err
		}
		if client.UserAgent == "" {
			return errors.New("UserAgent not set")
		}
		request.Header.Add("User-Agent", client.UserAgent)

		client := &http.Client{}
		rawresp, err := client.Do(request)
		defer rawresp.Body.Close()
		if err != nil {
			return err
		}
		var body []byte
		body, err = ioutil.ReadAll(rawresp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &model)
		return err
	}
}

func (client *Client) manage() {
	if client.clear == nil || len(client.clear) >= 150 {
		client.clear = make(chan bool, 100)
	}
	go func() {
		for {
			client.clear <- true
			time.Sleep(10 * time.Millisecond)
		}
	}()
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
