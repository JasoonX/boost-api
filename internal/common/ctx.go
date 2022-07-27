package common

import "context"

const (
	CtxLocaleKey = "locale"
)

func GetLocaleFromContext(ctx context.Context) string {
	return ctx.Value(CtxLocaleKey).(string)
}
