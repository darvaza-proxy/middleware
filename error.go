package middleware

import (
	"context"
	"net/http"

	"github.com/darvaza-proxy/core"
)

// ErrorHandlerFunc is the signature of a function used as ErrorHandler
type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error)

// WithErrorHandlerMiddleware returns a middleware that attaches a
// given ErrorHandler to their contexts
func WithErrorHandlerMiddleware(eh ErrorHandlerFunc) func(http.Handler) http.Handler {
	if eh == nil {
		core.Panic("no error handler provided")
	}

	fn := func(next http.Handler) http.Handler {
		h := func(rw http.ResponseWriter, req *http.Request) {
			ctx := WithErrorHandler(req.Context(), eh)
			req = req.WithContext(ctx)

			next.ServeHTTP(rw, req)
		}

		return http.HandlerFunc(h)
	}

	return fn
}

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
