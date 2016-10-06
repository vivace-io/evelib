package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type RedisQReciever func(Kill)

type redisqResp struct {
	Kill Kill `json:"package"`
}

type RedisQClient struct {
	UserAgent string
	errc      chan error
	recievers []RedisQReciever
	channels  []chan Kill
	running   bool
}

func NewRedisQClient(userAgent string) (client *RedisQClient) {
	client = new(RedisQClient)
	client.UserAgent = userAgent
	return
}

func (this *RedisQClient) Start() error {
	if this.running {
		return errors.New("already watching redisq")
	}
	this.running = true
	go func() {
		for this.running {
			kill, err := this.fetch()
			if err != nil {
				this.logError(err)
			} else {
				this.send(kill)
			}
		}
	}()
	return nil
}

func (this *RedisQClient) SetErrorChannel(errChan chan error) {
	this.errc = errChan
}

func (this *RedisQClient) Stop() {
	this.running = false
}

func (this *RedisQClient) send(k Kill) {
	for _, c := range this.channels {
		go func() {
			c <- k
		}()
	}
	for _, r := range this.recievers {
		go r(k)
	}
}

func (this *RedisQClient) AddChannel(output chan Kill) {
	this.channels = append(this.channels, output)
}

func (this *RedisQClient) AddReciever(reciever RedisQReciever) {
	this.recievers = append(this.recievers, reciever)
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
	} else {
		log.Printf("[ERROR][REDISQ] - %v", err)
	}
}
