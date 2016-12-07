package crest

// Corporation model for a Corporation entity in EVE Online.
type Corporation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}
