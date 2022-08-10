package listener

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/store"
	"github.com/BOOST-2021/boost-app-back/internal/listener/middlewares"
	"github.com/BOOST-2021/boost-app-back/internal/web"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/handlers"
)

func (l *service) setupRouter() {
	googleCredentials := l.cfg.Credentials(config.Google)
	goth.UseProviders(
		google.New(googleCredentials.ClientID, googleCredentials.ClientSecret, googleCredentials.RedirectURI),
	)

	l.router = chi.NewRouter()

	l.router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middlewares.Context(
			ctx.CtxLog(l.log),
			ctx.CtxDataProvider(store.New(l.cfg)),
			ctx.CtxAuthProvider(auth.New(l.cfg)),
		),
	)

	l.router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", handlers.GetHealthcheck)

		// Auth
		r.Get("/auth", handlers.GetAuthToken)
		r.With(web.SSOProviderContext).Get(fmt.Sprintf("/auth/{%s}", web.SSOProviderPathParam), handlers.GetAuthProvider)
		r.With(web.SSOProviderContext).Get(fmt.Sprintf("/auth/{%s}/callback", web.SSOProviderPathParam), handlers.GetAuthProviderCallback)
		r.With(web.Paginate).Get("/news", handlers.GetNews)
	})

}
