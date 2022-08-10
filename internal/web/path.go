package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	SSOProviderPathParam = "provider"
)

func SSOProviderContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// this is a workaround since gothic doesn't support chi router natively
		provider := chi.URLParam(r, SSOProviderPathParam)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), SSOProviderPathParam, provider)))
	})
}
