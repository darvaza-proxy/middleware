# func(http.Handler) http.Handler

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
