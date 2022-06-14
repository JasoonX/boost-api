package handlers

import (
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/internal/web/utils"
	"github.com/BOOST-2021/boost-app-back/resources"
	"net/http"
)

func GetHealthcheck(w http.ResponseWriter, r *http.Request) {
	render.Respond[*responses.EmptyDataResponse](
		w,
		resources.NewStatus(http.StatusOK, "Server runs successfully"),
		nil,
		utils.ServiceRuns(r),
	)
}
