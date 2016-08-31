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

type Client struct {
	Server    string
	UserAgent string
}

func NewClient(serverAddr, userAgent string) (client *Client) {
	client = &Client{
		Server:    serverAddr,
		UserAgent: userAgent,
	}
	return
}

type Key struct {
	ID    string
	VCode string
}

func NewKey(id, vCode string) Key {
	return Key{id, vCode}
}

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
