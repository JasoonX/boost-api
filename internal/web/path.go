package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	SSOProviderPathParam = "provider"
)

// SSOProviderContext is a workaround since gothic doesn't support chi router natively
func SSOProviderContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, SSOProviderPathParam)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), SSOProviderPathParam, provider)))
	})
}
