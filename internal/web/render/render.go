package render

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/web"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set(web.HeaderContentType, "application/json; charset=utf-8")
}

type Responder[T json.Marshaler] interface {
	Respond()

	WithStatus(status *resources.Status) Responder[T]
	WithData(data T) Responder[T]
	WithMeta(meta ...T) Responder[T]
	WithPage(page *resources.Page) Responder[T]
	WithLinks(links *resources.Links) Responder[T]
	WithErrors(errors responses.JSONServerErrors) Responder[T]
}

type responder[T json.Marshaler] struct {
	w http.ResponseWriter

	status *resources.Status
	data   T
	meta   []T
	page   *resources.Page
	links  *resources.Links
	errors responses.JSONServerErrors
}

func New(w http.ResponseWriter) Responder[json.Marshaler] {
	return &responder[json.Marshaler]{w: w}
}

func (r *responder[T]) Respond() {
	setDefaultHeaders(r.w)

	r.w.WriteHeader(int(r.status.Code))

	if r.status.Code == http.StatusNoContent {
		return
	}

	res := responses.Response[T]{
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

func (r *responder[T]) WithStatus(status *resources.Status) Responder[T] {
	r.status = status
	return r
}

func (r *responder[T]) WithData(data T) Responder[T] {
	r.data = data
	return r
}

func (r *responder[T]) WithMeta(meta ...T) Responder[T] {
	r.meta = meta
	return r
}

func (r *responder[T]) WithPage(page *resources.Page) Responder[T] {
	r.page = page
	return r
}

func (r *responder[T]) WithLinks(links *resources.Links) Responder[T] {
	r.links = links
	return r
}

func (r *responder[T]) WithErrors(errors responses.JSONServerErrors) Responder[T] {
	r.errors = errors
	return r
}
