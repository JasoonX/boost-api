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
		log.WithError(err).Error("failed to get search ownerships request")
		render.BadRequest(w, responses.JSONServerErrors{
			&resources.Error{
				Code:  111,
				Error: string(common.MustMarshal(err)),
			},
		}, nil)
		return
	}

	provider := ctx.DataProvider(r)

	pass, err := auth.HashPassword(req.Body.Data.Password)

	newRole := model.Viewer
	if req.Body.Data.Role != nil {
		newRole = webconvert.FromRequestRole(*req.Body.Data.Role)
	}
	newUser, err := provider.UsersProvider().AddUser(reqCtx, model.User{
		Username:     req.Body.Data.Username,
		FirstName:    req.Body.Data.FirstName,
		LastName:     req.Body.Data.LastName,
		IsAnonymous:  false,
		PasswordHash: pass,
		Status:       model.Unverified,
		Role:         newRole,
	})
	if err != nil {
		log.WithError(err).Error("failed to add new user")
		render.InternalServerError(w, nil)
		return
	}

	render.Success(w, webconvert.ToResponseUser(newUser))
}
