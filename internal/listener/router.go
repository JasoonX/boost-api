package listener

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/BOOST-2021/boost-app-back/internal/data/store"
	"github.com/BOOST-2021/boost-app-back/internal/listener/middlewares"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/handlers"
)

func (l *service) setupRouter() {
	l.router = chi.NewRouter()

	l.router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middlewares.Context(
			ctx.CtxLog(l.log),
			ctx.CtxProvider(store.New(l.cfg)),
		),
	)

	l.router.Get("/v1/healthcheck", handlers.GetHealthcheck)

	l.router.Get("/v1/news", handlers.GetNews)
}
