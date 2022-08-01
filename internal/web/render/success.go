package render

import (
	"encoding/json"
	"net/http"

	"github.com/BOOST-2021/boost-app-back/resources"
)

// these are just aliases for convenience
// for more flexibility use directly Responder

func Success(w http.ResponseWriter, data json.Marshaler) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusOK, "Request result ok")).
		WithData(data).
		Respond()
}

func SuccessPaged(w http.ResponseWriter, data json.Marshaler, page *resources.Page) {
	New(w).
		WithStatus(resources.NewStatus(http.StatusOK, "Request result ok")).
		WithPage(page).
		WithData(data).
		Respond()
}
