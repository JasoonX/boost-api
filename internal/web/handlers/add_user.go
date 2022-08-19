package handlers

import (
	"net/http"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/requests"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	webconvert "github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r).WithField("handler", "AddUser")
	reqCtx := r.Context()

	req, err := requests.NewAddUserRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to get add user request")
		render.BadRequest(w, responses.JSONServerErrors{
			&resources.Error{
				Code:  111,
				Error: string(common.MustMarshal(err)),
			},
		}, nil)
		return
	}

	provider := ctx.DataProvider(r)

	pass, err := auth.HashPassword(req.Body.Data.Attributes.Password)

	newUser, err := provider.UsersProvider().AddUser(reqCtx, model.User{
		Username:     req.Body.Data.Attributes.Username,
		FirstName:    req.Body.Data.Attributes.FirstName,
		LastName:     req.Body.Data.Attributes.LastName,
		PasswordHash: pass,
		Status:       model.UserStatusUnverified,
	})
	if err != nil {
		log.WithError(err).Error("failed to add new user")
		render.InternalServerError(w)
		return
	}

	render.Success(w, webconvert.ToResponseUser(newUser))
}
