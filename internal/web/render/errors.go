package render

import (
	"encoding/json"
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func InternalServerError(w http.ResponseWriter, meta ...json.Marshaler) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusInternalServerError, "Something went wrong")).
		WithMeta(meta...).
		Respond()
}

func BadRequest(w http.ResponseWriter, errors responses.JSONServerErrors, meta ...json.Marshaler) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusBadRequest, "Wrong payload provided")).
		WithMeta(meta...).
		WithErrors(errors).
		Respond()
}

func Unauthorized(w http.ResponseWriter, meta ...json.Marshaler) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusUnauthorized, "Unauthorized")).
		WithMeta(meta...).
		Respond()
}

func Forbidden(w http.ResponseWriter, meta ...json.Marshaler) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusForbidden, "Forbidden")).
		WithMeta(meta...).
		Respond()
}
