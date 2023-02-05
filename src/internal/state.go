package internal

import "encoding/json"

type State struct {
	Villians map[string]Villian `json:"villians"`
}

func (s *State) JSON() ([]byte) {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
