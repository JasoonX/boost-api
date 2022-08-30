package listener

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
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
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
			AllowedHeaders:   []string{web.HeaderAccept, web.HeaderAuthorization, web.HeaderContentType, web.HeaderLocalization, web.HeaderDevice, web.HeaderVersion},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)

	l.router.Get("/healthcheck", handlers.GetHealthcheck)

	l.router.With(middlewares.ValidateHeaders(l.cfg)).Route("/v1", func(r chi.Router) {

		// Auth
		r.Post("/auth", handlers.GetAuthToken)
		r.With(web.SSOProviderContext).Get(fmt.Sprintf("/auth/{%s}", web.SSOProviderPathParam), handlers.GetAuthProvider)
		r.With(web.SSOProviderContext).Get(fmt.Sprintf("/auth/{%s}/callback", web.SSOProviderPathParam), handlers.GetAuthProviderCallback)

		// Users
		r.Post("/users", handlers.AddUser)

		// News
		r.With(web.Paginate, middlewares.Doorman(l.cfg, model.UserRoleViewer, model.UserRoleReactViewer)).Get("/news", handlers.GetNews)
	})

}
