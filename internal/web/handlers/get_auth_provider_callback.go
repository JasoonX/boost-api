package handlers

import (
	"net/http"

	"github.com/markbates/goth/gothic"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	webconvert "github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func GetAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r).WithField("handler", "GetAuthProviderCallback")
	reqCtx := r.Context()

	googleUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.WithError(err).Error("failed to complete user auth")
		render.InternalServerError(w)
		return
	}

	dataProvider := ctx.DataProvider(r)
	user, err := dataProvider.UsersProvider().GetUserByEmail(reqCtx, googleUser.Email)
	if err != nil {
		log.WithError(err).Error("failed to get user by email")
		render.InternalServerError(w)
		return
	}

	authProvider := ctx.AuthProvider(r)
	tokenPair, err := authProvider.GenerateTokenPair(auth.UserInfo{
		Email: googleUser.Email,
		Roles: user.RolesStrings(),
	})
	if err != nil {
		log.WithError(err).Error("failed to generate token pair")
		render.InternalServerError(w)
		return
	}
	render.Success(w, &resources.AuthPost200Response{
		Data: webconvert.ToResponseTokenPair(tokenPair),
	})
}
