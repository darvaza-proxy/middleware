package middleware

import (
	"context"
	"net/http"

	"darvaza.org/core"
	"darvaza.org/x/web"
)

// ErrorHandlerFunc is the signature of a function used as ErrorHandler
type ErrorHandlerFunc = web.ErrorHandlerFunc

// WithErrorHandlerMiddleware returns a middleware that attaches a
// given ErrorHandler to their contexts
func WithErrorHandlerMiddleware(eh ErrorHandlerFunc) func(http.Handler) http.Handler {
	if eh == nil {
		core.Panic("no error handler provided")
	}

	fn := func(next http.Handler) http.Handler {
		h := func(rw http.ResponseWriter, req *http.Request) {
			ctx := web.WithErrorHandler(req.Context(), eh)
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
	return web.WithErrorHandler(ctx, eh)
}

// ErrorHandler attempts to pull an ErrorHandler from the context.Context
func ErrorHandler(ctx context.Context) (ErrorHandlerFunc, bool) {
	return web.ErrorHandler(ctx)
}
