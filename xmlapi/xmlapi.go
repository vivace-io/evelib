// Package xmlapi implements a client and accompanying functions to query Eve
// Online's XML API. The package does not enforce chaching or rate limiting,
// instead leaving that the the developers that choose to use this package.
package xmlapi

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

const (
	TranquilityURI = "https://api.eveonline.com"
	SingularityURI = "https://"
)

// Client is a client to the XML API
type Client struct {
	Server    string
	UserAgent string
}

// NewClient returns a new client configured with the passed server address
// and user agent.
func NewClient(serverAddr, userAgent string) (client *Client) {
	client = &Client{
		Server:    serverAddr,
		UserAgent: userAgent,
	}
	return
}

// Key is the a key/value pair used to authenticate requests with the XML API
type Key struct {
	ID    string
	VCode string
}

// NewKey returns a new Key (duh...)
func NewKey(id, vCode string) Key {
	return Key{id, vCode}
}

// Result is embedded in to all requests and contains information pertaining to
// the executed request
type Result struct {
	Version     int       `xml:"version,attr"`
	CurrentTime eTime     `xml:"currentTime"`
	Error       *APIError `xml:"error,omitempty"`
	CachedUntil eTime     `xml:"cachedUntil"`
}

func (this *Client) fetch(path string, args url.Values, key Key, model interface{}) error {
	uri := this.Server + path
	if args == nil {
		args = url.Values{}
	}
	args.Set("keyID", key.ID)
	args.Set("vCode", key.VCode)
	resp, err := http.PostForm(uri, args)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//TODO: LimitReader if it explodes?
	err = xml.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return err
	}
	return nil
}
