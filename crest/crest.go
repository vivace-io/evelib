package crest

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	TranquilityURI   = "https://crest-tq.eveonline.com/"
	SingularityURI   = "https://api-sisi.testeveonline.com/"
	MaxRateSustained = 150
	MaxRateBurst     = 400
	CRESTVersion     = "application/vnd.ccp.eve.Api-v2+json"
)

type Client struct {
	UserAgent  string
	URI        string
	tlsEnabled bool
	clear      chan bool
}

func NewClient(userAgent string, crestURI string, rate, burst int) (client *Client) {
	client = &Client{
		UserAgent: userAgent,
		URI:       crestURI,
	}
	client.manage(rate, burst)
	return
}

// TLS will enable/disable TLS verification when executing domain certificate checks.
// It is STRONGLY RECOMMENDED to not disable this. This is only used by the developer
// to allow for execution of library testing on Continuous Integration environments
// that do not support or include TLS certificates in the provisioned environment.
func (c *Client) TLS(enabled bool) {
	c.tlsEnabled = enabled
}

func (c *Client) get(path string, model interface{}) (err error) {
	select {
	case <-c.clear:
		if c.UserAgent == "" {
			return errors.New("user agent must be set")
		}
		var request *http.Request
		request, err = http.NewRequest("GET", c.URI+path, nil)
		if err != nil {
			return
		}
		request.Header.Add("User-Agent", c.UserAgent)
		request.Header.Add("Accept", CRESTVersion)

		webClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.tlsEnabled},
			},
		}
		var rawresp *http.Response
		rawresp, err = webClient.Do(request)
		defer rawresp.Body.Close()
		if err != nil {
			return
		}
		var body []byte
		body, err = ioutil.ReadAll(rawresp.Body)
		if err != nil {
			return fmt.Errorf("%v: %v", ErrCRESTUnmarshal, err)
		}
		// Check if there's an error in the response.
		err = responseError(rawresp, body)
		if err != nil {
			return
		}
		err = json.Unmarshal(body, &model)
		return err
	}
}

func (c *Client) manage(rate int, burst int) {
	c.clear = make(chan bool, burst)
	go func() {
		for {
			c.clear <- true
			time.Sleep(time.Duration(rate/1000) * time.Millisecond)
		}
	}()
}
