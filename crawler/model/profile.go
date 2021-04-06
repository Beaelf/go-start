package model

import (
	"encoding/json"
)

type Profile struct {
	Name  string
	Email string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return Profile{}, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
