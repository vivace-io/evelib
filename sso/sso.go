package sso

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// TranquilityOAuth root address.
	TranquilityOAuth = "https://login.eveonline.com"
	// SingularityOAuth root address.
	SingularityOAuth = "https://sisilogin.testeveonline.com"
	// ErrClientID is returned when the client ID is not specified.
	ErrClientID = errors.New("client ID must be set")
	// ErrClientSecret is returned when the client secret is not specified.
	ErrClientSecret = errors.New("client secret must be set")
	// ErrCallbackAddress is returned when the callback address is not set.
	ErrCallbackAddress = errors.New("callback address must be set")
	// ErrBadOAuthAddress is returned when the OAuth address is not equal to
	// TranquilityOAuth or SingularityOAuth. For the sake of the security of users,
	// the client does not support proxies.
	ErrBadOAuthAddress = errors.New("the provided OAuth root address is invalid")
	// ErrTooManyRequests is returned when EVE SSO responds with HTTP status 409,
	// which more generally means the client has made way too many requests to SSO.
	// If this is ever returned, the client should wait for a few minutes and retry.
	ErrTooManyRequests = errors.New("EVE SSO responded with HTTP status 409 (too many requests)")
	// ErrParsingResponse is returned when the authorization code exchage/refresh
	// methods could not parse the JSON response in to the map. Applications
	// should attempt to retry after a few seconds.
	ErrParsingResponse = errors.New("response returned from EVE SSO could not be parsed (do retry)")
)

// Client to EVE Online's Signle Sign-on service.
type Client struct {
	id         string
	secret     string
	oauth      string
	callback   string
	httpClient *http.Client
}

// NewClient configures and returns a new client. For bad options, client is
// returned as nil with an error.
func NewClient(opts *Options) (client *Client, err error) {
	if err = opts.Validate(); err != nil {
		return
	}
	client = &Client{
		id:         opts.ClientID,
		secret:     opts.ClientSecret,
		oauth:      opts.OAuthRoot,
		callback:   opts.CallbackAddress,
		httpClient: new(http.Client),
	}
	return
}

// Login redirects the client to EVE Online SSO. The state parameter is optional,
// however heavily recommened for security purposes. If no scopes are passed,
// then only basic authentication is used.
func (client *Client) Login(w http.ResponseWriter, r *http.Request, state string, scopes ...string) {
	url := fmt.Sprintf("%v/oauth/authorize/?response_type=code&redirect_uri=%v&client_id=%v&state=%v", client.oauth, client.callback, client.id, state)
	if len(scopes) > 0 {
		// Append the scopes, man.
	}
	http.Redirect(w, r, url, http.StatusFound)
}

// Exchange the authorization code for a token.
func (client *Client) Callback(code string) (data map[string]interface{}, err error) {
	url := fmt.Sprintf("%v/oauth/token/?grant_type=authorization_code&code=%v", client.oauth, code)
	var areq *http.Request
	if areq, err = http.NewRequest("POST", url, nil); err != nil {
		return
	}
	// Sweet mother of nested functions...
	areq.Header.Set("Authorization", fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", client.id, client.secret)))))
	var aresp *http.Response
	if aresp, err = client.httpClient.Do(areq); err != nil {
		return
	}
	defer aresp.Body.Close()
	if aresp.StatusCode != http.StatusOK {
		if aresp.StatusCode == http.StatusConflict {
			// We've made too many requests to SSO.
			err = ErrTooManyRequests
			return
		} else {
			err = fmt.Errorf("EVE SSO responded with HTTP status %v", aresp.StatusCode)
			return
		}
	}
	var raw []byte
	if raw, err = ioutil.ReadAll(aresp.Body); err != nil {
		return
	}
	data = make(map[string]interface{})
	if err = json.Unmarshal(raw, &data); err != nil {
		data = nil
		err = ErrParsingResponse
		return
	}
	return
}

// Refresh an old token for a new one.
func (client *Client) Refresh(old map[string]interface{}) (new map[string]interface{}, err error) {
	// TODO - Implement...
	return
}
