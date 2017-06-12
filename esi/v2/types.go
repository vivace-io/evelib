package esi

import "fmt"

func (client *Client) TypeIDs() (results []int, err error) {
	for p := 1; ; p++ {
		var ids []int
		path := buildPath(fmt.Sprintf("/universe/types/?page=%v", p))
		if err = client.get(path, &ids); err != nil {
			err = fmt.Errorf("failed to retrieve page %v of types: %v", p, err)
			break
		}
		if len(ids) != 0 {
			results = append(results, ids...)
		} else {
			break
		}
	}
	return
}
