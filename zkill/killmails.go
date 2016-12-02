package zkill

import "time"

// Historical returns all kill ID's and their accompanying hashes submittted to
// zKillboard on the provided date.
func (client *Client) Historical(date time.Time) (killmails map[int]string, err error) {
	// TODO
	return
}
