package ctx

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/data/store"
)

const (
	ctxLog          = "ctxLog"
	ctxDataProvider = "ctxDataProvider"
	ctxAuthProvider = "ctxAuthProvider"
)

func CtxLog(log *logrus.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxLog, log)
	}
}

func Log(r *http.Request) *logrus.Entry {
	return r.Context().Value(ctxLog).(*logrus.Entry)
}

func CtxDataProvider(provider store.DataProvider) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxDataProvider, provider)
	}
}

func DataProvider(r *http.Request) store.DataProvider {
	return r.Context().Value(ctxDataProvider).(store.DataProvider)
}

func CtxAuthProvider(provider auth.Provider) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxAuthProvider, provider)
	}
}

func AuthProvider(r *http.Request) auth.Provider {
	return r.Context().Value(ctxAuthProvider).(auth.Provider)
}
