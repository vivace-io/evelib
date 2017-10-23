package redisq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	// DefaultAddr is the default address the RedisQ client queries.
	DefaultAddr = "https://redisq.zkillboard.com/listen.php"
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
	Killmail Killmail `json:"killmail"`
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
	queueID   string
	userAgent string
	running   bool
	errchan   chan error
	webClient *http.Client
}

// NewClient takes an option parameter, which may be nil, and configures and
// returns a new client.
func NewClient(opts *Options) (client *Client) {
	if opts == nil {
		opts = DefaultOptions()
	}
	client = &Client{
		locker:    new(sync.RWMutex),
		addr:      opts.Addr,
		queueID:   opts.QueueID,
		webClient: &http.Client{},
	}
	if client.addr == "" {
		client.addr = DefaultAddr
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

// AddFunc adds a function to the client which will be called andpassed all
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
	var request *http.Request
	if request, err = http.NewRequest("GET", client.addr, nil); err != nil {
		if err != nil {
			// Panic instead of returning, as any error returned is a bug on our end
			// that breaks package functionality entirely.
			panic(err)
		}
	}
	if client.queueID != "" {
		request.URL.Query().Set("queueID", client.queueID)
	}

	var rawresp *http.Response
	if rawresp, err = client.webClient.Do(request); err != nil {
		if err != nil {
			return
		}
	}
	defer rawresp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(rawresp.Body); err != nil {
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
	payload.Killmail.Value = payload.Zkb.TotalValue
	payload.Killmail.LocationID = payload.Zkb.LocationID
	payload.Killmail.KillHash = payload.Zkb.Hash
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
