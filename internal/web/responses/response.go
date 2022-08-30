package responses

import (
	"encoding/json"

	"github.com/BOOST-2021/boost-app-back/resources"
)

type Response[T json.Marshaler] struct {
	Data     T                 `json:"data,omitempty"`
	Included T                 `json:"included,omitempty"`
	Meta     []T               `json:"meta,omitempty"`
	Page     *resources.Page   `json:"page,omitempty"`
	Links    *resources.Links  `json:"links,omitempty"`
	Status   *resources.Status `json:"status,omitempty"`
	Errors   JSONServerErrors  `json:"errors,omitempty"`
}
