package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
<<<<<<< HEAD
=======
	"time"
>>>>>>> 231cf7d91084be67f3f16cd3fd696295b1fc6653
)

/* zkill/util.go
 * Defines utility functions for the library
 */

<<<<<<< HEAD
func checkUserAgent() error {
	if UserAgent == "" {
=======
func (c *Client) checkUserAgent() error {
	if c.UserAgent == "" {
>>>>>>> 231cf7d91084be67f3f16cd3fd696295b1fc6653
		return errors.New("UserAgent must be set")
	}
	return nil
}

<<<<<<< HEAD
func fetch(path string, model interface{}) error {
	err := checkUserAgent()
	if err == nil {
		select {
		case <-clear:
			request, err := http.NewRequest("GET", APIAddr+path, nil)
			if err != nil {
				return err
			}
			request.Header.Add("User-Agent", UserAgent)

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

			return json.Unmarshal(body, &model)
		}
	}
	return err
}

func formatOptions(opts Options) (result string) {
	switch {
	case opts.Solo:
		result += "/solo"
	case opts.Kills:
		result += "kills"
	case opts.Losses:
	}
	if opts.Solo {
		result += "/solo/"
	}
	if opts.Kills
=======
func (c *Client) okToRun() bool {
	err := c.checkUserAgent()
	if err != nil {
		return false
	}
	return true
}

func (c *Client) fetch(path string, model interface{}) error {
	if c.okToRun() {
		request, err := http.NewRequest("GET", c.APIAddress+path, nil)
		if err != nil {
			return err
		}
		request.Header.Add("User-Agent", c.UserAgent)

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
		throttle()
		return err
	}
	return errors.New("UserAgent not set")
}

// ServerStatus calls zKillboard to measure latency in MS and determine if there
// are connectivity issues.
func (c *Client) ServerStatus() (latency time.Duration, err error) {
	return 100 * time.Millisecond, nil
}

func throttle() {
	time.Sleep(10 * time.Millisecond)
>>>>>>> 231cf7d91084be67f3f16cd3fd696295b1fc6653
}
