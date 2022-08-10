package requests

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/resources"
)

type GetAuthTokenRequest struct {
	Body resources.AuthPostRequest
}

func NewGetAuthTokenRequest(r *http.Request) (*GetAuthTokenRequest, error) {
	req := GetAuthTokenRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(err, "failed to decode creds")
	}
	return &req, nil
}
