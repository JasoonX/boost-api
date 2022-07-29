package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/requests"
	"github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)

	_, err := requests.NewGetNewsRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to get search ownerships request")
		render.InternalServerError(w, nil)
		return
	}

	provider := ctx.Provider(r)

	news, err := provider.NewsProvider().ListNews(r.Context())
	if err != nil {
		log.WithError(err).Error("failed to get news from provider")
		render.InternalServerError(w, nil)
		return
	}

	render.Respond(w, resources.NewStatus(http.StatusOK, "Request result ok"), &resources.NewsGet200ResponseData{
		News: convert.ToResponseNews(news),
	}, nil)
}