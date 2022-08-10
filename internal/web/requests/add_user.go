package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/resources"
)

type AddUserRequest struct {
	Body resources.UsersAddPostRequest
}

// TODO: full validation
func (r AddUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Body, validation.Required),
		validation.Field(&r.Body.Data, validation.Required),
		validation.Field(&r.Body.Data.Email, validation.Required, is.Email),
		validation.Field(&r.Body.Data.Password, validation.Required, validation.Length(8, 32)),
		validation.Field(&r.Body.Data.Role, validation.In(model.UserRoles)),
	)
}

func NewAddUserRequest(r *http.Request) (*AddUserRequest, error) {
	req := AddUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &req, req.Validate()
}
