package web

import (
	"context"
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/params"
)

const (
	ctxPageParams = "ctxPageParams"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := ctx.Log(r).WithField("middleware", "Paginate")

		pageParams := params.PageParams{}
		if err := urlvals.Decode(r.URL.Query(), &pageParams); err != nil {
			log.WithError(err).Error("failed to decode page params")
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxPageParams, pageParams)))
	})
}

func PageParams(r *http.Request) params.PageParams {
	return r.Context().Value(ctxPageParams).(params.PageParams)
}
