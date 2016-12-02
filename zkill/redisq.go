package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// RedisQReciever is a function that accepts a RedisQ kill.
type RedisQReciever func(Kill)

type redisqResp struct {
	Kill Kill `json:"package"`
}

// RedisQClient is a client to ZKillboard's RedisQ service.
type RedisQClient struct {
	UserAgent string
	locker    *sync.RWMutex
	errc      chan error
	recievers []RedisQReciever
	channels  []chan Kill
	running   bool
}

// NewRedisQClient returns a new client to ZKillboard's RedisQ service.
func NewRedisQClient(userAgent string) (client *RedisQClient) {
	client = new(RedisQClient)
	client.UserAgent = userAgent
	client.locker = new(sync.RWMutex)
	return
}

// AddChannel registers a channel to be sent all future kills from RedisQ. Bear
// in mind that the RedisQClient will start a new goroutine for each channel
// added when pushing the Kill to the channel, so if the channel is not being
// emptied fast enough or is in deadlock, the program implementing client client
// will likely crash and burn. Safe for asynchronous use.
func (client *RedisQClient) AddChannel(output chan Kill) {
	client.locker.Lock()
	defer client.locker.Unlock()
	client.channels = append(client.channels, output)
}

// AddReciever registers a function of type RedisQReciever to be sent all future
// kills from RedisQ. Bear in mind that the clien will send the kill in a
// seperate goroutine for each reciever, so if the funcion takes too long/hangs
// the program will quickly become bloated and crash. Safe for asynchronous use.
func (client *RedisQClient) AddReciever(reciever RedisQReciever) {
	client.locker.Lock()
	defer client.locker.Unlock()
	client.recievers = append(client.recievers, reciever)
}

// Start listening to RedisQ. Once started, does not block.
func (client *RedisQClient) Start() error {
	if client.running {
		return errors.New("already watching redisq")
	}
	client.running = true
	go func() {
		for client.running {
			kill, err := client.fetch()
			if err != nil {
				// Usually if an error is encountered, it's a connection issue. So
				// ease back and wait for a minute when client happens and retry after.
				client.logError(err)
				time.Sleep(1 * time.Minute)
			} else {
				client.send(kill)
			}
		}
	}()
	return nil
}

// SetErrorChannel takes a channel of any size and asynchronously sends it
// all errors RedisQ experiences while running. Errors are sent on a seperate
// goroutine, but take care that the channel never fills for too long.
func (client *RedisQClient) SetErrorChannel(errChan chan error) {
	client.errc = errChan
}

// Stop RedisQ from retrieving any more kills.
func (client *RedisQClient) Stop() {
	client.running = false
}

func (client *RedisQClient) send(k Kill) {
	client.locker.RLock()
	defer client.locker.RUnlock()
	for _, c := range client.channels {
		go func() {
			c <- k
		}()
	}
	for _, r := range client.recievers {
		go r(k)
	}
}

func (client *RedisQClient) fetch() (k Kill, err error) {
	webc := &http.Client{}
	request, err := http.NewRequest("GET", DefaultRedisQURI, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", UserAgent)

	rawresp, err := webc.Do(request)
	if err != nil {
		return
	}
	defer rawresp.Body.Close()

	body, err := ioutil.ReadAll(rawresp.Body)
	if err != nil {
		return
	}
	zresp := redisqResp{}
	err = json.Unmarshal(body, &zresp)
	k = zresp.Kill
	if k.Killmail != nil {
		k.Killmail.KillHash = k.Zkb.Hash
	}
	return
}

func (client *RedisQClient) logError(err error) {
	if client.errc != nil {
		go func() {
			client.errc <- err
		}()
	}
}
