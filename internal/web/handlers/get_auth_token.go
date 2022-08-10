package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/requests"
	webconvert "github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func GetAuthToken(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r).WithField("handler", "GetAuthToken")
	reqCtx := r.Context()

	req, err := requests.NewGetAuthTokenRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to get search ownerships request")
		render.InternalServerError(w, nil)
		return
	}

	dataProvider := ctx.DataProvider(r)
	user, err := dataProvider.UsersProvider().GetUserByEmail(reqCtx, req.Body.Data.Email)
	if err != nil {
		log.WithError(err).Error("failed to get user by email")
		render.InternalServerError(w, nil)
		return
	}

	if !auth.CheckPasswordHash(req.Body.Data.Password, user.PasswordHash) {
		render.Unauthorized(w, nil)
	}

	authProvider := ctx.AuthProvider(r)
	tokenPair, err := authProvider.GenerateTokenPair(auth.UserInfo{
		Email: req.Body.Data.Email,
		Role:  string(user.Role),
	})
	if err != nil {
		log.WithError(err).Error("failed to generate token pair")
		render.InternalServerError(w, nil)
		return
	}
	render.Success(w, &resources.AuthPost200Response{
		Data: webconvert.ToResponseTokenPair(tokenPair),
	})
}
