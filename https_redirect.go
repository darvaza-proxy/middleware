package middleware

import (
	"fmt"
	"net"
	"net/http"

	"github.com/darvaza-proxy/middleware/internal"
)

// HTTPSRedirectHandler provides an automatic redirect to HTTPS
type HTTPSRedirectHandler struct {
	Port int
}

func (h *HTTPSRedirectHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Scheme != "https" {
		url := *req.URL
		url.Scheme = "https"

		host := url.Hostname()
		if h.Port != 0 && h.Port != 443 {
			port := fmt.Sprintf("%v", h.Port)
			url.Host = net.JoinHostPort(host, port)
		} else {
			url.Host = host
		}

		loc := url.String()

		rw.Header().Add("Location", loc)
		rw.WriteHeader(http.StatusPermanentRedirect)
		internal.Fprintf(rw, "Redirected to %s", loc)
	} else {
		http.NotFound(rw, req)
	}
}

// NewHTTPSRedirectHandler creates a new automatic redirect to HTTPS handler
func NewHTTPSRedirectHandler(port int) http.Handler {
	return &HTTPSRedirectHandler{
		Port: port,
	}
}
