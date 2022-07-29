package responses

import (
	"encoding/json"

	"github.com/BOOST-2021/boost-app-back/resources"
)

type Response struct {
	Data   json.Marshaler    `json:"data,omitempty"`
	Meta   json.Marshaler    `json:"meta,omitempty"`
	Status *resources.Status `json:"status,omitempty"`
	Errors *JSONServerErrors `json:"error,omitempty"`
}
