package redisq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/vivace-io/evelib/crest"
)

// RecieverChan is a channel that may recive a payload containing information
// pertaining to a killmail.
type RecieverChan chan Payload

// RecieverFunc is a function that may called and passed a new Payload model as
// the parameter.
type RecieverFunc func(Payload)

// Payload is the json encoded response containing the killmail that is returned
// as well as some meta information.
type Payload struct {
	KillID int `json:"killID"`
	Zkb    struct {
		LocationID int     `json:"locationID"`
		Hash       string  `json:"hash"`
		TotalValue float32 `json:"totalValue"`
		Points     int     `json:"points"`
		Href       string  `json:"href"`
	}
	Killmail crest.Killmail `json:"killmail"`
}

type response struct {
	Payload *Payload `json:"package"`
}

// Client to listen to ZKillboards RedisQ service for new Killmails recieved
// from EVE Online.
type Client struct {
	locker    *sync.RWMutex
	rchan     []RecieverChan
	rfunc     []RecieverFunc
	addr      string
	userAgent string
	running   bool
	errchan   chan error
}

// NewClient takes an option parameter, which may be nil, and configures and
// returns a new client. Returns error for invalid configuration or any
// connection problems.
func NewClient(opts *Options) (client *Client, err error) {
	if opts == nil {
		opts = DefaultOptions()
	}
	client = &Client{
		locker: new(sync.RWMutex),
		addr:   opts.Addr,
	}
	if client.addr == "" {
		err = errors.New("address for RedisQ endpoint was not set")
		client = nil
		return
	}
	return
}

// AddChan adds a channel to the client which will be sent all future Killmails
// that are recieved.
func (client *Client) AddChan(reciever RecieverChan) {
	client.locker.Lock()
	defer client.locker.Unlock()
	client.rchan = append(client.rchan, reciever)
}

// AddChan adds a function to the client which will be called andpassed all
// future killmails that are recieved.
func (client *Client) AddFunc(reciever RecieverFunc) {
	client.locker.Lock()
	defer client.locker.Unlock()
	client.rfunc = append(client.rfunc, reciever)
}

// SetErrorChannel accepts a channel of type error to which errors encountered
// are sent by the client.
func (client *Client) SetErrorChannel(errchan chan error) {
	client.locker.Lock()
	defer client.locker.Unlock()
	client.errchan = errchan
}

// Listen starts recieving and distributing killmails to added recievers.
func (client *Client) Listen() {
	if client.running {
		return
	}
	client.running = true
	go func() {
		var err error
		var resp response
		for client.running {
			resp, err = client.fetch()
			if err != nil {
				client.sendError(err)
				// Usually if an error is encountered, it's a connection issue. So
				// ease back and wait for a minute when client happens and retry after.
				time.Sleep(1 * time.Minute)
			} else {
				if resp.Payload != nil {
					client.send(*resp.Payload)
				}
			}
		}
	}()
	return
}

// Close the client's connection to RedisQ.
func (client *Client) Close() {
	client.running = false
}

// fetch a killmail from RedisQ.
func (client *Client) fetch() (resp response, err error) {
	webc := &http.Client{}
	request, err := http.NewRequest("GET", client.addr, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", client.userAgent)

	rawresp, err := webc.Do(request)
	if err != nil {
		return
	}
	defer rawresp.Body.Close()

	body, err := ioutil.ReadAll(rawresp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

// send a killmail to all reciever channels/functions on seperate goroutines
// for each.
func (client *Client) send(payload Payload) {
	client.locker.RLock()
	defer client.locker.RUnlock()

	for _, c := range client.rchan {
		go func() {
			c <- payload
		}()
	}
	for _, r := range client.rfunc {
		go r(payload)
	}
}

// sendError to error channel, if it exists.
func (client *Client) sendError(err error) {
	if client.errchan == nil {
		return
	}
	go func() {
		client.errchan <- err
	}()
}
