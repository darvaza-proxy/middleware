package middleware

import (
	"context"
	"net/http"
)

// ErrorHandlerFunc is the signature of a function used as ErrorHandler
type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error)

// WithErrorHandler attaches an ErrorHandler function to a context
// for later retrieval
func WithErrorHandler(ctx context.Context, eh ErrorHandlerFunc) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if eh != nil {
		ctx = context.WithValue(ctx, errHandlerKey, eh)
	}

	return ctx
}

// ErrorHandler attempts to pull an ErrorHandler from the context.Context
func ErrorHandler(ctx context.Context) (ErrorHandlerFunc, bool) {
	eh, ok := ctx.Value(errHandlerKey).(ErrorHandlerFunc)
	return eh, ok
}

type contextKey struct {
	name string
}

var (
	errHandlerKey = &contextKey{"ErrorHandler"}
)
