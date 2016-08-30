package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/* zkill/util.go
 * Defines utility functions for the library
 */

func checkUserAgent() error {
	if UserAgent == "" {
		return errors.New("UserAgent must be set")
	}
	return nil
}

func okToRun() bool {
	err := checkUserAgent()
	if err != nil {
		return false
	}
	return true
}

func fetch(path string, model interface{}) error {
	if okToRun() {
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

		err = json.Unmarshal(body, &model)
		return err
	}
	return errors.New("UserAgent not set")
}
