package zkill

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

/* zkill/util.go
 * Defines utility functions for the library
 */

func queryFromOpts(opts *Options) string {
	var result string
	// TODO - time options
	if opts == nil {
		return result
	}
	if opts.BeforeKillID != 0 {
		result += "/beforeKillID/" + strconv.Itoa(opts.BeforeKillID)
	}
	if opts.AfterKillID != 0 {
		result += "/afterKillID/" + strconv.Itoa(opts.AfterKillID)
	}
	if opts.Solo {
		result += "/solo"
	}
	if opts.Kills {
		result += "/kills"
	}
	if opts.Losses {
		result += "/losses"
	}
	if opts.WSpace {
		result += "/w-space"
	}

	// lastly, ensure trailing slash in place
	if result != "" {
		result += "/"
	}
	return result
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
