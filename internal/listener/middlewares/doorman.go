package middlewares

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/web"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
)

func Doorman(cfg config.Runtime, roles ...model.UserRole) func(http.Handler) http.Handler {
	// always has access to everything by default
	roles = append(roles, model.UserRoleSuperAdmin)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := ctx.Log(r).WithField("middleware", "Validate Headers")

			// disabling auth for test purposes
			if cfg.AuthDisabled() {
				log.Warning("auth is disabled")
				next.ServeHTTP(w, r)
				return
			}

			rawToken := r.Header.Get(web.HeaderAuthorization)
			claims, err := ctx.AuthProvider(r).Verify(rawToken)
			if err != nil {
				if errors.Is(err, auth.ErrExpiredToken) || errors.Is(err, auth.ErrInvalidToken) {
					render.Forbidden(w, responses.Reason{Reason: errors.Cause(err).Error()})
					return
				}
			}
			user, err := ctx.DataProvider(r).UsersProvider().GetUserByEmail(r.Context(), claims.Email)
			if err != nil {
				if errors.Is(err, data.ErrNotFound) {
					render.Forbidden(w, responses.Reason{Reason: auth.ErrUserNotFound.Error()})
					return
				}
				render.InternalServerError(w)
				return
			}
			for _, role := range roles {
				if user.IsRole(role) {
					next.ServeHTTP(w, r)
					return
				}
			}
			render.Forbidden(w, responses.Reason{Reason: auth.ErrInsufficientRole.Error()})
		})
	}
}
