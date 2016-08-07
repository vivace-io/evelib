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
	TranquilityURI   CRESTRoot = "https://crest-tq.eveonline.com/"
	SingularityURI   CRESTRoot = "https://api-sisi.testeveonline.com/"
	MaxRateSustained           = 150
	MaxRateBurst               = 400
)

var (
	UserAgent      string
	Root           CRESTRoot
	ErrUserAgent   = errors.New("user agent must be specified")
	ErrCRESTRoot   = errors.New("CREST root URI must be specified")
	ErrRateDomain  = errors.New("the give rate must be within a domain of [1, 150]")
	ErrBurstDomain = errors.New("the given burst must be within a domain of [150, 400]")
	clear          chan bool
	webClient      *http.Client
	queryCount     int
)

type CRESTRoot string

// Init the CREST package for basic functionality.
// `userAgent` Your contact information, passed in the header of all CREST requests.
// `root` The root server of CREST. Use package TranquilityURI, SingularityURI or a custom CRESTRoot.
// `rate` The maximum requests per second to be made. Input domain of [1, 150].
// `burst` The maximum request burst allowed. Input domain of [150, 400].
// Returns error if input incorrect.
func Init(userAgent string, root CRESTRoot, rate, burst int) error {
	if userAgent == "" {
		return ErrUserAgent
	}
	UserAgent = userAgent
	if root == "" {
		return ErrCRESTRoot
	}
	Root = root
	webClient = new(http.Client)
	if rate < 1 || rate > 150 {
		return ErrRateDomain
	}
	if burst < 150 || burst > 400 {
		return ErrBurstDomain
	}
	manage(rate, burst)
	return nil
}

// TLSEnabled will enable/disable TLS verification when executing checks.
// It is STRONGLY RECOMMENDED to not disable this. This is only used by the Developer
// to allow for execution of library testing on Continuous Integration environments
// that do not support or include TLS certificates in the provisioned environment.
func TLSEnabled(enabled bool) {
	webClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !enabled},
	}
}

// QueryCount returns the number of queries executed since the package was
// initialized or since the count was last reset.
func QueryCount() int {
	return queryCount
}

// ResetQueryCount resets the query counter.
func ResetQueryCount() {
	queryCount = 0
}

func fetch(path string, model interface{}) (err error) {
	select {
	case <-clear:
		if UserAgent == "" {
			return errors.New("user agent must be set")
		}
		var request *http.Request
		request, err = http.NewRequest("GET", string(Root)+path, nil)
		if err != nil {
			return
		}
		request.Header.Add("User-Agent", UserAgent)

		var rawresp *http.Response
		rawresp, err = webClient.Do(request)
		queryCount++
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

func manage(rate int, burst int) {
	clear = make(chan bool, burst)
	go func() {
		for {
			clear <- true
			time.Sleep(time.Duration(rate/1000) * time.Millisecond)
		}
	}()
}
