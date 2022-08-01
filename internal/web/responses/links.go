package responses

import (
	"net/http"
	"net/url"

	"github.com/BOOST-2021/boost-app-back/resources"
)

func BuildPageLinks(r http.Request, selfQuery string) *resources.Links {
	urlBuilder := func(rawPath string) *url.URL {
		return &url.URL{
			Scheme:   r.URL.Scheme,
			Host:     r.URL.Host,
			Path:     r.URL.Path,
			RawQuery: rawPath,
		}
	}
	return &resources.Links{
		Self: urlBuilder(selfQuery).String(),
	}
}
