package listener

import (
	"github.com/BOOST-2021/boost-app-back/internal/data/store"
	"github.com/BOOST-2021/boost-app-back/internal/listener/middlewares"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (l *service) setupRouter() {
	l.router = chi.NewRouter()

	l.router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middlewares.Context(
			ctx.CtxLog(l.log),
			ctx.CtxProvider(store.New(l.cfg)),
		),
	)

	l.router.Get("/healthcheck", handlers.GetHealthcheck)
}
