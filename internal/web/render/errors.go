package render

import (
	"encoding/json"
	"net/http"

	"github.com/BOOST-2021/boost-app-back/resources"
)

func InternalServerError(w http.ResponseWriter, meta json.Marshaler) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusInternalServerError, "Something went wrong")).
		WithMeta(meta).
		Respond()
}
