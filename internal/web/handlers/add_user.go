package handlers

import (
	"net/http"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/data"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/web/ctx"
	"github.com/BOOST-2021/boost-app-back/internal/web/render"
	"github.com/BOOST-2021/boost-app-back/internal/web/requests"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/internal/web/utils"
	webconvert "github.com/BOOST-2021/boost-app-back/internal/web/utils/convert"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r).WithField("handler", "AddUser")
	reqCtx := r.Context()

	req, err := requests.NewAddUserRequest(r)
	if err != nil {
		if valErrors, err := utils.UnwrapValidationErrors(err); err == nil {
			render.BadRequest(w, valErrors)
			return
		}
		log.WithError(err).Error("failed to get add user request")
		render.InternalServerError(w)
		return
	}

	provider := ctx.DataProvider(r)

	if _, err := provider.EmailsProvider().GetEmail(reqCtx, req.Body.Data.Attributes.Email); err != nil {
		if !errors.Is(err, data.ErrNotFound) {
			// TODO: fixme
			duplicateError := resources.NewErrorWithDefaults()
			duplicateError.SetDetail(map[string]interface{}{
				"reason": "email already exists",
			})
			render.BadRequest(w, responses.JSONServerErrors{
				duplicateError,
			})
			return
		}
	}

	pass, err := auth.HashPassword(req.Body.Data.Attributes.Password)
	if err != nil {
		log.WithError(err).Error("failed to hash password")
		render.InternalServerError(w)
	}

	newUser, err := provider.UsersProvider().AddUser(reqCtx, model.User{
		Username:     req.Body.Data.Attributes.Username,
		FirstName:    req.Body.Data.Attributes.FirstName,
		LastName:     req.Body.Data.Attributes.LastName,
		PasswordHash: pass,
		Status:       model.UserStatusUnverified,
		// TODO: refactor that
		Role: pq.StringArray{string(model.UserRoleViewer)},
	})
	if err != nil {
		log.WithError(err).Error("failed to add new user")
		render.InternalServerError(w)
		return
	}

	newEmail, err := provider.EmailsProvider().AddEmail(reqCtx, model.Email{
		Email:      req.Body.Data.Attributes.Email,
		IsVerified: false,
		IsPrimary:  true,
		UserID:     newUser.ID,
	})
	if err != nil {
		log.WithError(err).Error("failed to add email")
		render.InternalServerError(w)
		return
	}

	if req.Body.Data.Attributes.HasPhone() {
		newPhone, err := provider.PhonesProvider().AddPhone(reqCtx, model.Phone{
			SubscriberNumber: req.Body.Data.Attributes.Phone.SubscriberNumber,
			CountryCode:      &req.Body.Data.Attributes.Phone.CountryCode,
			Extension:        req.Body.Data.Attributes.Phone.Extension,
			IsVerified:       false,
			IsPrimary:        true,
			UserID:           newUser.ID,
		})
		if err != nil {
			log.WithError(err).Error("failed to add phone")
			render.InternalServerError(w)
			return
		}
		newUser.Phones = []model.Phone{
			convert.FromPtr(newPhone),
		}
	}

	newUser.Emails = []model.Email{
		convert.FromPtr(newEmail),
	}

	authProvider := ctx.AuthProvider(r)
	tokenPair, err := authProvider.GenerateTokenPair(auth.UserInfo{
		Email: req.Body.Data.Attributes.Email,
		Roles: newUser.RolesStrings(),
	})
	if err != nil {
		log.WithError(err).Error("failed to generate token pair")
		render.InternalServerError(w)
		return
	}
	render.Success(w, resources.UsersPost201ResponseData{
		User:      convert.FromPtr(webconvert.ToResponseUser(newUser)),
		TokenPair: webconvert.ToResponseTokenPair(tokenPair),
	})
}
