package middlewares

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/common/enums"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/web"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
)

func ValidateHeaders(cfg config.Runtime) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !cfg.AuthDisabled() && r.Header.Get(web.HeaderAuthorization) == "" {
				render.Unauthorized(w, responses.Reason{Reason: "missing authorization header"})
				return
			}
			// we might also want to query request from browser window (for debug purposes)
			if cfg.Environment() == config.EnvironmentLocal {
				r.Header.Set(web.HeaderAuthorization, enums.EN)
				r.Header.Set(web.HeaderDevice, enums.Web)
				r.Header.Set(web.HeaderVersion, cfg.Version())
			}
			// basic headers validation
			if r.Header.Get(web.HeaderLocalization) == "" {
				render.BadRequest(w, nil, responses.Reason{Reason: "missing header: Accept-Language"})
				return
			}
			if r.Header.Get(web.HeaderDevice) == "" {
				render.BadRequest(w, nil, responses.Reason{Reason: "missing header: User-Agent"})
				return
			}
			if r.Header.Get(web.HeaderVersion) == "" {
				render.BadRequest(w, nil, responses.Reason{Reason: "missing header: Accept-Version"})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
