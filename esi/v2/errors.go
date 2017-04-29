package esi

import "fmt"

type HTTPError struct {
	Code int `json:"-"`
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("request returned status code %v", e.Code)
}
