package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/web"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/requests"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	webconvert "github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r).WithField("handler", "GetNews")
	reqCtx := r.Context()

	_, err := requests.NewGetNewsRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to get search ownerships request")
		render.InternalServerError(w)
		return
	}

	provider := ctx.DataProvider(r)

	count, err := provider.NewsProvider().CountNews(reqCtx)
	if err != nil {
		log.WithError(err).Error("failed to count news")
		render.InternalServerError(w)
		return
	}

	pageParams := web.PageParams(r)
	pageParams.Total = &count

	news, err := provider.NewsProvider().OfPage(reqCtx, webconvert.FromResponsePage(pageParams)).ListNews(reqCtx)
	if err != nil {
		log.WithError(err).Error("failed to get news from provider")
		render.InternalServerError(w)
		return
	}

	render.SuccessPaged(w, &resources.NewsGet200ResponseData{
		News: webconvert.ToResponseNews(news),
	},
		webconvert.ToResponsePage(pageParams),
		responses.BuildLinks(r, pageParams),
	)
}
