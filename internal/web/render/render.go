package render

import (
	"encoding/json"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/resources"
	"github.com/pkg/errors"
	"net/http"
)

// Respond valid json respond rendering
func Respond[D, M json.Marshaler](w http.ResponseWriter, status *resources.Status, data D, meta M) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(status.GetCode()))

	if status.GetCode() == http.StatusNoContent {
		return
	}

	res := responses.Response[D, M]{
		Status: status,
		Data:   data,
		Meta:   meta,
	}
	res.Status = status

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}
