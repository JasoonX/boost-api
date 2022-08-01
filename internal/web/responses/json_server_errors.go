package responses

import (
	"encoding/json"

	"github.com/BOOST-2021/boost-app-back/resources"
)

type JSONServerErrors []*resources.Error

func (m JSONServerErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal([]*resources.Error(m))
}

func (m JSONServerErrors) Clean() JSONServerErrors {
	var filtered []*resources.Error
	for _, v := range m {
		if v != nil {
			filtered = append(filtered, v)
		}
	}
	return m
}
