package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func GetHealthcheck(w http.ResponseWriter, r *http.Request) {
	render.New(w).
		WithStatus(resources.NewStatus(http.StatusOK, "Server runs successfully")).
		Respond()
}
