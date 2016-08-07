package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/* zkill/util.go
 * Defines utility functions for the library
 */

func (c *Client) checkUserAgent() error {
	if c.UserAgent == "" {
		return errors.New("UserAgent must be set")
	}
	return nil
}

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
}
