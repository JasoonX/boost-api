package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/resources"
)

type AddUserRequest struct {
	Body resources.UsersPostRequest
}

func (r AddUserRequest) Validate() error {
	return validation.ValidateStruct(&r.Body,
		validation.Field(&r.Body.Data),
	)
}

func NewAddUserRequest(r *http.Request) (*AddUserRequest, error) {
	req := AddUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &req, req.Validate()
}
