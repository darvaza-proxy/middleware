// Package internal provides helpers for middleware implementations
package internal

import (
	"fmt"
	"io"
)

// Fprintf is just like fmt.Fprintf but without returning errors
// we will ignore anyway.
func Fprintf(w io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(w, format, args...)
}
