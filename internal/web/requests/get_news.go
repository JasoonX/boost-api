package requests

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals"
)

type GetNewsRequest struct {
	urlvals.PageParams
}

func NewGetNewsRequest(r *http.Request) (*GetNewsRequest, error) {
	req := GetNewsRequest{}
	if err := urlvals.Decode(r.URL.Query(), &req); err != nil {
		return nil, errors.Wrap(err, "failed to decode params")
	}
	return &req, nil
}
