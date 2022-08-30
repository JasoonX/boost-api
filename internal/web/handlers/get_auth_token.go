package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/requests"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	webconvert "github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
)

func GetAuthToken(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r).WithField("handler", "GetAuthToken")
	reqCtx := r.Context()

	req, err := requests.NewGetAuthTokenRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to get search ownerships request")
		render.InternalServerError(w)
		return
	}

	dataProvider := ctx.DataProvider(r)
	user, err := dataProvider.UsersProvider().GetUserByEmail(reqCtx, req.Body.Data.Attributes.Email)
	if err != nil {
		log.WithError(err).Error("failed to get user by email")
		render.InternalServerError(w)
		return
	}

	if !auth.CheckPasswordHash(req.Body.Data.Attributes.Password, user.PasswordHash) {
		render.Forbidden(w, responses.Reason{Reason: "invalid password"})
		return
	}

	authProvider := ctx.AuthProvider(r)
	tokenPair, err := authProvider.GenerateTokenPair(auth.UserInfo{
		Email: req.Body.Data.Attributes.Email,
		Roles: user.RolesStrings(),
	})
	if err != nil {
		log.WithError(err).Error("failed to generate token pair")
		render.InternalServerError(w)
		return
	}
	render.Success(w, webconvert.ToResponseTokenPair(tokenPair))
}
