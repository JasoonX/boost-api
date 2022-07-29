package render

import (
	"encoding/json"
	"net/http"

	"github.com/BOOST-2021/boost-app-back/resources"
)

func InternalServerError(w http.ResponseWriter, meta json.Marshaler) {
	Error(w, resources.NewError(http.StatusInternalServerError, "Something went wrong"), nil, meta)
}
