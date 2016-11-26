package crest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrCRESTUnmarshal = errors.New("unexpected error unmarshalling CREST response")
	ErrUserAgent      = errors.New("user agent must be specified")
	ErrCRESTURI       = errors.New("CREST root URI must be specified")
	ErrRateDomain     = errors.New("the give rate must be within a domain of [1, 150]")
	ErrBurstDomain    = errors.New("the given burst must be within a domain of [150, 400]")
)

// ErrorHTTPStatus for all HTTP status code errors not covered.
type ErrorHTTPStatus struct {
	StatusCode int
}

func (e ErrorHTTPStatus) Error() string {
	return fmt.Sprintf("unexpected HTTP status code %v", e.StatusCode)
}

// ErrorCREST represents an error returned by CREST
type ErrorCREST struct {
	StatusCode    int    `json:"statusCode"`
	Message       string `json:"message"`
	Key           string `json:"key"`
	ExceptionType string `json:"exceptionType"`
}

func (e ErrorCREST) Error() string {
	return fmt.Sprintf("CREST responded %v : %v : %v", e.Message, e.Key, e.ExceptionType)
}

func IsCRESTUnmarshalErr(err error) bool {
	if strings.Contains(err.Error(), ErrCRESTUnmarshal.Error()) {
		return true
	}
	return false
}

func responseError(resp *http.Response, body []byte) error {
	var err error
	switch resp.StatusCode {
	case 404:
		e := new(ErrorCREST)
		err = json.Unmarshal(body, &e)
		if err != nil {
			if !jsonError(err) {
				return err
			}
		}
		e.StatusCode = resp.StatusCode
		return e
	}
	return nil
}

func jsonError(err error) bool {
	if strings.Contains(err.Error(), "unexpected end of JSON input") {
		return true
	}
	return false
}
