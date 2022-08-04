package render

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

type Responder interface {
	Respond()

	WithStatus(status *resources.Status) Responder
	WithData(data json.Marshaler) Responder
	WithMeta(meta json.Marshaler) Responder
	WithPage(page *resources.Page) Responder
	WithLinks(links *resources.Links) Responder
	WithErrors(errors responses.JSONServerErrors) Responder
}

type responder struct {
	w http.ResponseWriter

	status *resources.Status
	data   json.Marshaler
	meta   json.Marshaler
	page   *resources.Page
	links  *resources.Links
	errors responses.JSONServerErrors
}

func New(w http.ResponseWriter) Responder {
	return &responder{w: w}
}

func (r *responder) Respond() {
	setDefaultHeaders(r.w)

	r.w.WriteHeader(int(r.status.Code))

	if r.status.Code == http.StatusNoContent {
		return
	}

	res := responses.Response{
		Status: r.status,
		Data:   r.data,
		Meta:   r.meta,
		Page:   r.page,
		Links:  r.links,
		Errors: r.errors,
	}

	err := json.NewEncoder(r.w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}

func (r *responder) WithStatus(status *resources.Status) Responder {
	r.status = status
	return r
}

func (r *responder) WithData(data json.Marshaler) Responder {
	r.data = data
	return r
}

func (r *responder) WithMeta(meta json.Marshaler) Responder {
	r.meta = meta
	return r
}

func (r *responder) WithPage(page *resources.Page) Responder {
	r.page = page
	return r
}

func (r *responder) WithLinks(links *resources.Links) Responder {
	r.links = links
	return r
}

func (r *responder) WithErrors(errors responses.JSONServerErrors) Responder {
	r.errors = errors
	return r
}
