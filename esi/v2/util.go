package esi

import (
	"fmt"
	"strings"
)

// buildPath is a quick and dirty fix for the fact that ESI does not have all
// v1 paths for v2 (looking at you, /killmails/). Essentially, this checks the
// route given, attempts to find an absolute path URL to the best match.
// EXAMPLES:
// `/killmails/{someID}/{someHash}`
//      returns
//          https://esi.tech.ccp.is/v1/killmails/{someID}/{someHash}
func (client *Client) buildPath(route string) (path string) {
	if strings.Contains(route, "/killmails/") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if strings.Contains(route, "/markets/groups/") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if route == "/markets/prices/" {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if route == "/status/" {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if strings.Contains(route, "/universe/regions/") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if strings.Contains(route, "/universe/groups/") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if strings.Contains(route, "/universe/types/?page=") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if strings.Contains(route, "/universe/types/") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v2", route)
	}
	if strings.Contains(route, "/universe/constellations/") {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if route == "/universe/systems/" {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v1", route)
	}
	if strings.Contains("/universe/systems/", route) {
		return fmt.Sprintf("%v/%v%v", client.apiRoot, "v2", route)
	}
	return fmt.Sprintf("%v/%v%v", client.apiRoot, "v2", route)
}
