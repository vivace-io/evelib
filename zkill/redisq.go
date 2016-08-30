package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	running         = false
	redisQRecievers []RedisQReciever
	redisQChannels  []chan Kill
	redisqErrors    chan error
)

type RedisQReciever func(Kill)

// ZKillboardRedisQ default endpoint
const ZKillboardRedisQ = "https://redisq.zkillboard.com/listen.php"

var alreadyRunning = false

type redisqResp struct {
	Kill Kill `json:"package"`
}

func RedisQStart() error {
	if running {
		return errors.New("already watching redisq")
	}
	running = true
	go func() {
		for running {
			kill, err := fetchRedisQ()
			if err != nil {
				logRedisQError(err)
			} else {
				redisqSend(kill)
			}
		}
	}()
	return nil
}

func RedisQSetErrorChannel(errChan chan error) {
	redisqErrors = errChan
}

func RedisQStop() {
	running = false
}

func redisqSend(k Kill) {
	for _, c := range redisQChannels {
		go func() {
			c <- k
		}()
	}
	for _, r := range redisQRecievers {
		go r(k)
	}
}

func RedisQAddChannel(output chan Kill) {
	redisQChannels = append(redisQChannels, output)
}

func RedisQAddReciever(reciever RedisQReciever) {
	redisQRecievers = append(redisQRecievers, reciever)
}

func fetchRedisQ() (k Kill, err error) {
	request, err := http.NewRequest("GET", RedisQAddr, nil)
	if err != nil {
		return
	}
	request.Header.Add("User-Agent", UserAgent)

	rawresp, err := webClient.Do(request)
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

func logRedisQError(err error) {
	if redisqErrors != nil {
		go func() {
			redisqErrors <- err
		}()
	} else {
		log.Printf("[ERROR][REDISQ] - %v", err)
	}
}
