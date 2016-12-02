package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type RedisQReciever func(Kill)

type redisqResp struct {
	Kill Kill `json:"package"`
}

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
// emptied fast enough or is in deadlock, the program implementing this client
// will likely crash and burn. Safe for asynchronous use.
func (this *RedisQClient) AddChannel(output chan Kill) {
	this.locker.Lock()
	defer this.locker.Unlock()
	this.channels = append(this.channels, output)
}

// AddReciever registers a function of type RedisQReciever to be sent all future
// kills from RedisQ. Bear in mind that the clien will send the kill in a
// seperate goroutine for each reciever, so if the funcion takes too long/hangs
// the program will quickly become bloated and crash. Safe for asynchronous use.
func (this *RedisQClient) AddReciever(reciever RedisQReciever) {
	this.locker.Lock()
	defer this.locker.Unlock()
	this.recievers = append(this.recievers, reciever)
}

// Start listening to RedisQ. Once started, does not block.
func (this *RedisQClient) Start() error {
	if this.running {
		return errors.New("already watching redisq")
	}
	this.running = true
	go func() {
		for this.running {
			kill, err := this.fetch()
			if err != nil {
				// Usually if an error is encountered, it's a connection issue. So
				// ease back and wait for a minute when this happens and retry after.
				this.logError(err)
				time.Sleep(1 * time.Minute)
			} else {
				this.send(kill)
			}
		}
	}()
	return nil
}

// SetErrorChannel takes a channel of any size and asynchronously sends it
// all errors RedisQ experiences while running. Errors are sent on a seperate
// goroutine, but take care that the channel never fills for too long.
func (this *RedisQClient) SetErrorChannel(errChan chan error) {
	this.errc = errChan
}

// Stop RedisQ from retrieving any more kills.
func (this *RedisQClient) Stop() {
	this.running = false
}

func (this *RedisQClient) send(k Kill) {
	this.locker.RLock()
	defer this.locker.RUnlock()
	for _, c := range this.channels {
		go func() {
			c <- k
		}()
	}
	for _, r := range this.recievers {
		go r(k)
	}
}

func (this *RedisQClient) fetch() (k Kill, err error) {
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

func (this *RedisQClient) logError(err error) {
	if this.errc != nil {
		go func() {
			this.errc <- err
		}()
	}
}
