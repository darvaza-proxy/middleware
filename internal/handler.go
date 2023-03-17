package internal

import "net/http"

// BadRequestHandler renders and 400 response with an error
func BadRequestHandler(rw http.ResponseWriter, _ *http.Request, err error) {
	http.Error(rw, err.Error(), http.StatusBadRequest)
}
