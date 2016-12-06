package zkill

import (
	"fmt"
	"strconv"
	"time"
)

// Historical returns all kill ID's and their accompanying hashes submittted to
// zKillboard on the provided date.
func (client *Client) Historical(date time.Time) (killmails map[int]string, err error) {
	// TODO
	results := make(map[string]string)
	client.fetch(fmt.Sprintf("/history/%v/", date.Format("20060102")), &results)
	var id int
	for key, val := range results {
		if id, err = strconv.Atoi(key); err != nil {
			err = fmt.Errorf("unable to convert string value '%v' to type int: %v", key, err)
			return
		} else {
			killmails[id] = val
		}
	}
	return
}
