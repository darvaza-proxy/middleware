package middleware

import (
	"net/http"

	"darvaza.org/darvaza/shared/web/qlist"
)

// AcceptMiddleware filters the Accept header to only include those
func AcceptMiddleware(supported ...string) func(http.Handler) http.Handler {
	valid, fallback := acceptSupportedAsRanges(supported)
	if len(valid) == 0 {
		return NOOP
	}

	return func(next http.Handler) http.Handler {
		h := func(rw http.ResponseWriter, req *http.Request) {
			acceptMiddlewareHandler(rw, req, next, valid, fallback)
		}
		return http.HandlerFunc(h)
	}
}

func acceptSupportedAsRanges(supported []string) (out []qlist.QualityValue, fallback string) {
	out = make([]qlist.QualityValue, 0, len(supported))

	for _, s := range supported {
		r, err := qlist.ParseMediaRange(s)
		if err == nil {
			if len(out) == 0 {
				fallback = s
			}
			out = append(out, r)
		}
	}

	return out, fallback
}

func acceptMiddlewareHandler(rw http.ResponseWriter, req *http.Request,
	next http.Handler, supported []qlist.QualityValue,
	fallback string) {
	//
	var s string
	if len(supported) > 0 {
		accepted, _ := qlist.ParseMediaRangeHeader(req.Header)
		if len(accepted) > 0 {
			s, _, _ = qlist.BestQualityParsed(supported, accepted)
		}
	}

	if s == "" {
		s = fallback
	}

	req.Header[qlist.Accept] = []string{s}
	next.ServeHTTP(rw, req)
}
