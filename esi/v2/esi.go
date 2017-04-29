package esi

import (
	"errors"
	"net/http"
	// httpmemcache "github.com/gregjones/httpcache/memcache"
	//"github.com/bradfitz/gomemcache/memcache"
	//"github.com/gregjones/httpcache"
	//"crypto/tls"
	"encoding/json"
	"io/ioutil"
)

const (
	APIAddress = "https://esi.tech.ccp.is"
)

type Client struct {
	UserAgent  string
	httpClient *http.Client
}

func NewClient(options *Options) (client *Client) {
	if options != nil {
		client = &Client{
			httpClient: options.HTTPClient,
			UserAgent:  options.UserAgent,
		}
		if client.httpClient == nil {
			client.httpClient = new(http.Client)
		}
	} else {
		client = &Client{
			httpClient: new(http.Client),
		}
	}
	return client
}

func (client *Client) get(path string, dest interface{}) (err error) {
	if client.UserAgent == "" {
		return errors.New("user agent not set for client")
	}
	var request *http.Request
	request, err = http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	request.Header.Add("X-User-Agent", client.UserAgent)
	request.Header.Add("Accept", "application/json")
	var rawresp *http.Response
	rawresp, err = client.httpClient.Do(request)
	if err != nil {
		return
	}
	if rawresp.StatusCode != 200 {
		err = HTTPError{Code: rawresp.StatusCode}
		return
	}
	defer rawresp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(rawresp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &dest)
	return

}
