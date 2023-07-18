# func(http.Handler) http.Handler

[![Go Reference](https://pkg.go.dev/badge/darvaza.org/middleware.svg)](https://pkg.go.dev/darvaza.org/middleware)
[![Codebeat Score](https://codebeat.co/badges/d6e66082-010e-4b2b-8ef7-98de1345713b)](https://codebeat.co/projects/github-com-darvaza-proxy-middleware-main)

## General Middleware

* `NOOP` does nothing

## Content Negotiation

* `AcceptMiddleware()` alters the `Accept` header so only a given set is passed
  to the next handler in the chain

## Error Handling Middleware

* `WithErrorHandlerMiddleware()`

## Special Handlers

* HTTPSRedirectHandler a handler that will redirect to https if
  non-https, or return 404 already https

## Error Handlers

we define an ErrorHandler as a function like:

```go
func (http.ResponseWriter, *http.Request, error)
```

`WithErrorHandler()` and `ErrorHandler()` add and read the context
for such handler
