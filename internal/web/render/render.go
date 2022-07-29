package render

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/resources"
)

// Respond valid json respond rendering
func Respond(w http.ResponseWriter, status *resources.Status, data json.Marshaler, meta json.Marshaler) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(status.Code))

	if status.Code == http.StatusNoContent {
		return
	}

	res := responses.Response{
		Status: status,
		Data:   data,
		Meta:   meta,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}

func Error(w http.ResponseWriter, error *resources.Error, serverErrors *responses.JSONServerErrors, meta json.Marshaler) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(error.GetCode()))

	res := responses.Response{
		Errors: serverErrors,
		Meta:   meta,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}
