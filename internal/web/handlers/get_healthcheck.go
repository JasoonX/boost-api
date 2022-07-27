package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/utils"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func GetHealthcheck(w http.ResponseWriter, r *http.Request) {
	render.Respond(
		w,
		resources.NewStatus(http.StatusOK, "Server runs successfully"),
		nil,
		utils.ServiceRuns(r),
	)
}
