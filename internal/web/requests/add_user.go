package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/resources"
)

type AddUserRequest struct {
	Body resources.UsersPostRequest
}

// TODO: full validation
func (r AddUserRequest) Validate() error {
	return validation.Errors{
		"body":     validation.Validate(&r.Body, validation.Required),
		"data":     validation.Validate(&r.Body.Data, validation.Required),
		"email":    validation.Validate(&r.Body.Data.Attributes.Email, validation.Required, is.Email),
		"password": validation.Validate(&r.Body.Data.Attributes.Password, validation.Required, validation.Length(8, 32)),
	}.Filter()
}

func NewAddUserRequest(r *http.Request) (*AddUserRequest, error) {
	req := AddUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &req, req.Validate()
}
