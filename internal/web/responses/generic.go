package responses

import (
	"encoding/json"
	"github.com/BOOST-2021/boost-app-back/resources"
)

type Response[D, M json.Marshaler] struct {
	Data   D                 `json:"data,omitempty"`
	Meta   M                 `json:"meta,omitempty"`
	Status *resources.Status `json:"status,omitempty"`
	Error  *resources.Error  `json:"error,omitempty"`
}
