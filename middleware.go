// Package middleware provices standard http middleware modules
package middleware

import "net/http"

// NOOP is a middleware that doesn't do anything
func NOOP(next http.Handler) http.Handler {
	return next
}
